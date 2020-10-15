//go:generate mockgen -destination persons_mock/persons_mock.go -source persons.go

package persons

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	pgx "github.com/jackc/pgx/v4"
)

var (
	ErrPersonsNotFount = errors.New("person not found")
)

type Repository interface {
	Get(ctx context.Context, personID int64) (*models.Person, error)
	GetList(ctx context.Context, filter *v1.PersonFilter, limit, offset uint, sort, order string) ([]models.Person, error)
	Save(ctx context.Context, person *models.Person) (*models.Person, error)
	Update(ctx context.Context, person *models.Person) (*models.Person, error)
	Remove(ctx context.Context, personID int64) (int64, error)
}

type queryer interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type repository struct {
	database queryer
}

func NewRepository(database queryer) *repository {
	return &repository{
		database: database,
	}
}

func (r *repository) Get(ctx context.Context, personID int64) (*models.Person, error) {
	// Список запрашиваемых полей
	columns := []string{
		"id",
		"team_id",
		"first_name",
		"middle_name",
		"last_name",
		"birthday",
		"email",
		"phone",
		"slack",
		"role",
		"duty_order",
		"is_active",
		"updated_at",
		"created_at",
	}

	for i, columnName := range columns {
		columns[i] = fmt.Sprintf(`"p".%q`, columnName)
	}

	// Формирование запроса
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf(`select %s from "persons" as "p" where "p"."id" = $1`, strings.Join(columns, ",")))

	// Выполнение запроса
	persons, err := r.query(ctx, query.String(), personID)
	if err != nil {
		return nil, err
	}

	return &persons[0], nil
}

func (r *repository) GetList(
	ctx context.Context,
	filter *v1.PersonFilter,
	limit, offset uint,
	sort, order string,
) ([]models.Person, error) {
	// Список запрашиваемых полей
	columns := []string{
		"id",
		"team_id",
		"first_name",
		"middle_name",
		"last_name",
		"birthday",
		"email",
		"phone",
		"slack",
		"role",
		"duty_order",
		"is_active",
		"updated_at",
		"created_at",
	}

	for i, columnName := range columns {
		columns[i] = fmt.Sprintf(`"p".%q`, columnName)
	}

	// Формирование запроса
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf(`select %s from "persons" as "p"`, strings.Join(columns, ",")))
	where, args := r.where(filter)
	if where != "" {
		query.WriteString(" ")
		query.WriteString(where)
	}
	query.WriteString(fmt.Sprintf(` order by "p".%q %s limit %d offset %d;`, order, sort, limit, offset))

	// Выполнение запроса
	return r.query(ctx, query.String(), args...)
}

func (r *repository) Save(ctx context.Context, p *models.Person) (*models.Person, error) {
	dutyOrder, err := r.getNextDutyOrder(ctx, p.TeamID)
	if err != nil {
		return nil, err
	}

	attributes := map[string]interface{}{
		"team_id":     p.TeamID,
		"first_name":  p.FirstName,
		"middle_name": p.MiddleName,
		"last_name":   p.LastName,
		"birthday":    p.Birthday,
		"email":       p.Email,
		"phone":       p.Phone,
		"role":        p.Role,
		"slack":       p.Slack,
		"duty_order":  dutyOrder,
		"is_active":   p.IsActive,
	}

	return r.put(ctx, attributes)
}

func (r *repository) Update(ctx context.Context, p *models.Person) (*models.Person, error) {
	var attributes = map[string]interface{}{
		"id":          p.ID,
		"first_name":  p.FirstName,
		"middle_name": p.MiddleName,
		"last_name":   p.LastName,
		"birthday":    p.Birthday,
		"email":       p.Email,
		"phone":       p.Phone,
		"slack":       p.Slack,
		"role":        p.Role,
		"is_active":   p.IsActive,
		"updated_at":  time.Now().Local(),
	}

	currentPerson, err := r.Get(ctx, p.TeamID)
	if err != nil {
		return nil, err
	}

	if currentPerson.TeamID != p.TeamID {
		dutyOrder, err := r.getNextDutyOrder(ctx, p.TeamID)
		if err != nil {
			return nil, err
		}
		attributes["team_id"] = p.TeamID
		attributes["duty_order"] = dutyOrder
	}

	return r.put(ctx, attributes)
}

func (r *repository) Remove(ctx context.Context, personID int64) (int64, error) {
	var query bytes.Buffer
	query.WriteString("delete from teams where id = $1;")
	_, err := r.database.Query(ctx, query.String(), personID)

	return personID, err
}

func (r *repository) put(ctx context.Context, attributes map[string]interface{}) (*models.Person, error) {
	var (
		i           uint8
		err         error
		values      []interface{}
		columns     []string
		returns     []string
		placeholder []string
		update      []string
	)

	// Имена возвращаемых полей
	returns = append(returns, "id")
	returns = append(returns, "team_id")
	returns = append(returns, "first_name")
	returns = append(returns, "middle_name")
	returns = append(returns, "last_name")
	returns = append(returns, "birthday")
	returns = append(returns, "email")
	returns = append(returns, "phone")
	returns = append(returns, "slack")
	returns = append(returns, "role")
	returns = append(returns, "duty_order")
	returns = append(returns, "is_active")
	returns = append(returns, "updated_at")
	returns = append(returns, "created_at")

	// Подготовка коллекции значений, имен полей и плейсхолдеров для передаваемых значений
	for k, v := range attributes {
		i += 1
		values = append(values, v)
		columns = append(columns, fmt.Sprintf("%q", k))
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
		update = append(update, fmt.Sprintf("%s=excluded.%s", k, k))
	}

	query := fmt.Sprintf(`insert into "persons" (%s) values (%s) on conflict (id) do update set %s returning %s`,
		strings.Join(columns, ","),
		strings.Join(placeholder, ","),
		strings.Join(update, ","),
		strings.Join(returns, ","))

	// Выполнение запроса
	persons, err := r.query(ctx, query, values...)
	if err != nil {
		return nil, err
	}

	return &persons[0], nil
}

func (r *repository) query(ctx context.Context, query string, args ...interface{}) ([]models.Person, error) {
	logger.Log.WithField("query", query).Info("SQL query")

	rows, err := r.database.Query(ctx, query, args...)
	if isEmpty(err) {
		return nil, ErrPersonsNotFount
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var persons []models.Person

	for rows.Next() {
		var p person
		err := rows.Scan(
			&p.ID,
			&p.TeamID,
			&p.FirstName,
			&p.MiddleName,
			&p.LastName,
			&p.Birthday,
			&p.Email,
			&p.Phone,
			&p.Slack,
			&p.Role,
			&p.DutyOrder,
			&p.IsActive,
			&p.Updated,
			&p.Created,
		)

		if err != nil {
			return nil, err
		}

		person := p.convert()
		persons = append(persons, person)
	}

	return persons, nil
}

func (r *repository) where(filter *v1.PersonFilter) (string, []interface{}) {
	var (
		i           uint8
		where       []string
		values      []interface{}
		placeholder []string
	)

	if filter == nil {
		return "", values
	}

	// Подготовка значений для фильтра по ID и заглушек для их подстановки
	for _, id := range filter.PersonIds {
		i += 1
		values = append(values, id)
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
	}

	// Генерация запроса по ID
	if len(filter.PersonIds) > 0 {
		where = append(where, fmt.Sprintf(`"p"."id" in (%s)`, strings.Join(placeholder, ",")))
		placeholder = nil
	}

	// Подготовка значений для фильтра по TeamID и заглушек для их подстановки
	for _, id := range filter.TeamIds {
		i += 1
		values = append(values, id)
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
	}

	// Генерация запроса по TeamID
	if len(filter.TeamIds) > 0 {
		where = append(where, fmt.Sprintf(`"p"."team_id" in (%s)`, strings.Join(placeholder, ",")))
		placeholder = nil
	}

	// Генерация запроса по DateFrom
	if filter.DateFrom != nil && filter.DateFrom.Seconds > 0 {
		i += 1
		where = append(where, fmt.Sprintf(`"t"."created_at" > $%d`, i))
		values = append(values, time.Unix(filter.DateFrom.Seconds, 0).Format(time.RFC3339Nano))
	}

	// Генерация запроса по DateTo
	if filter.DateTo != nil && filter.DateTo.Seconds > 0 {
		i += 1
		where = append(where, fmt.Sprintf(`"t"."created_at" < $%d`, i))
		values = append(values, time.Unix(filter.DateTo.Seconds, 0).Format(time.RFC3339Nano))
	}

	// Склеивание тела запроса
	if len(where) > 0 {
		return "where " + strings.Join(where, " and "), values
	}

	return "", values
}

func isEmpty(err error) bool {
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return true
		}
	}
	return false
}

type person struct {
	ID         sql.NullInt64
	TeamID     sql.NullInt64
	FirstName  sql.NullString
	MiddleName sql.NullString
	LastName   sql.NullString
	Birthday   sql.NullTime
	Email      sql.NullString
	Phone      sql.NullString
	Slack      sql.NullString
	Role       sql.NullInt64
	DutyOrder  sql.NullInt64
	IsActive   sql.NullBool
	Created    sql.NullTime
	Updated    sql.NullTime
}

func (p *person) convert() models.Person {
	model := models.Person{
		ID:        p.ID.Int64,
		TeamID:    p.TeamID.Int64,
		FirstName: p.FirstName.String,
		LastName:  p.LastName.String,
		Slack:     p.Slack.String,
		Role:      models.Role(p.Role.Int64),
		DutyOrder: p.DutyOrder.Int64,
		IsActive:  p.IsActive.Bool,
		Created:   p.Created.Time,
	}

	if p.MiddleName.Valid {
		model.MiddleName = &p.MiddleName.String
	}

	if p.Email.Valid {
		model.Email = &p.Email.String
	}

	if p.Phone.Valid {
		model.Phone = &p.Phone.String
	}

	if p.TeamID.Valid {
		model.TeamID = p.TeamID.Int64
	}

	if p.Birthday.Valid {
		model.Birthday = &p.Birthday.Time
	}

	if p.Updated.Valid {
		model.Updated = p.Updated.Time
	}

	return model
}

func (r *repository) getNextDutyOrder(ctx context.Context, teamID int64) (int64, error) {
	const query = `select coalesce(max("duty_order"), 0) from "persons" where team_id=$1;`
	var dutyOrder int64

	row := r.database.QueryRow(ctx, query, teamID)
	err := row.Scan(&dutyOrder)
	if isEmpty(err) {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("next duty order for team %v error: %w", teamID, err)
	}

	return dutyOrder + 1, nil
}
