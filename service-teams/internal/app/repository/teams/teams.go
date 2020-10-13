//go:generate mockgen -destination teams_mock/teams_mock.go -source teams.go

package teams

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bgoldovsky/dutyer/service-teams/internal/app/models"
	v1 "github.com/bgoldovsky/dutyer/service-teams/internal/generated/rpc/v1"
	pgx "github.com/jackc/pgx/v4"
)

var (
	ErrTeamNotFount = errors.New("teams not found")
)

type Repository interface {
	Get(ctx context.Context, teamID int64) (*models.Team, error)
	GetList(ctx context.Context, filter *v1.TeamFilter, limit, offset uint, sort, order string) ([]models.Team, error)
	Save(ctx context.Context, team *models.Team) (*models.Team, error)
	Update(ctx context.Context, team *models.Team) (*models.Team, error)
	Remove(ctx context.Context, teamID int64) (int64, error)
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

func (r *repository) Get(ctx context.Context, teamID int64) (*models.Team, error) {
	// Список запрашиваемых полей
	columns := []string{"id", "name", "description", "slack", "created_at", "updated_at"}
	for i, c := range columns {
		columns[i] = fmt.Sprintf(`"t".%q`, c)
	}

	// Формирование запроса
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf(`select %s from "teams" as "t" where "t"."id" = $1`, strings.Join(columns, ",")))

	// Выполнение запроса
	teams, err := r.query(ctx, query.String(), teamID)
	if err != nil {
		return nil, err
	}

	return &teams[0], nil
}

func (r *repository) GetList(
	ctx context.Context,
	filter *v1.TeamFilter,
	limit, offset uint,
	sort, order string,
) ([]models.Team, error) {
	// Список запрашиваемых полей
	columns := []string{"id", "name", "description", "slack", "created_at", "updated_at"}
	for i, c := range columns {
		columns[i] = fmt.Sprintf(`"t".%q`, c)
	}

	// Формирование запроса
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf(`select %s from "teams" as "t"`, strings.Join(columns, ",")))
	where, args := r.where(filter)
	if where != "" {
		query.WriteString(" ")
		query.WriteString(where)
	}
	query.WriteString(fmt.Sprintf(` order by "t".%q %s limit %d offset %d;`, order, sort, limit, offset))

	// Выполнение запроса
	return r.query(ctx, query.String(), args...)
}

func (r *repository) Save(ctx context.Context, team *models.Team) (*models.Team, error) {
	attributes := map[string]interface{}{
		"name":        team.Name,
		"description": team.Description,
		"slack":       team.Slack,
	}

	return r.put(ctx, attributes)
}

func (r *repository) Update(ctx context.Context, t *models.Team) (*models.Team, error) {
	var attributes = map[string]interface{}{
		"id":          t.ID,
		"name":        t.Name,
		"description": t.Description,
		"slack":       t.Slack,
		"updated_at":  time.Now().Local(),
	}

	return r.put(ctx, attributes)
}

func (r *repository) Remove(ctx context.Context, teamID int64) (int64, error) {
	var query bytes.Buffer
	query.WriteString("delete from teams where id = $1;")
	_, err := r.database.Query(ctx, query.String(), teamID)

	return teamID, err
}

func (r *repository) put(ctx context.Context, attributes map[string]interface{}) (*models.Team, error) {
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
	returns = append(returns, "name")
	returns = append(returns, "description")
	returns = append(returns, "slack")
	returns = append(returns, "created_at")
	returns = append(returns, "updated_at")

	// Подготовка коллекции значений, имен полей и плейсхолдеров для передаваемых значений
	for k, v := range attributes {
		i += 1
		values = append(values, v)
		columns = append(columns, fmt.Sprintf("%q", k))
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
		update = append(update, fmt.Sprintf("%s=excluded.%s", k, k))
	}

	// Построение sql запрса
	query := fmt.Sprintf(`insert into "teams" (%s) values (%s) on conflict (id) do update set %s returning %s`,
		strings.Join(columns, ","),
		strings.Join(placeholder, ","),
		strings.Join(update, ","),
		strings.Join(returns, ","))

	// Выполнение запроса
	teams, err := r.query(ctx, query, values...)
	if err != nil {
		return nil, err
	}

	return &teams[0], nil
}

func (r *repository) query(ctx context.Context, query string, args ...interface{}) ([]models.Team, error) {
	rows, err := r.database.Query(ctx, query, args...)
	if isEmpty(err) {
		return nil, ErrTeamNotFount
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var teams []models.Team

	for rows.Next() {
		var t team
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Slack, &t.Created, &t.Updated); err != nil {
			return nil, err
		}
		team := t.convert()
		teams = append(teams, team)
	}

	return teams, nil
}

func (r *repository) where(filter *v1.TeamFilter) (string, []interface{}) {
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
	for _, id := range filter.Ids {
		i += 1
		values = append(values, id)
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
	}

	// Генерация запроса по ID
	if len(filter.Ids) > 0 {
		where = append(where, fmt.Sprintf(`"t"."id" in (%s)`, strings.Join(placeholder, ",")))
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

type team struct {
	ID          sql.NullInt64
	Name        sql.NullString
	Description sql.NullString
	Slack       sql.NullString
	Created     sql.NullTime
	Updated     sql.NullTime
}

func (t *team) convert() models.Team {
	model := models.Team{
		ID:          t.ID.Int64,
		Name:        t.Name.String,
		Description: t.Description.String,
		Slack:       t.Slack.String,
		Created:     t.Created.Time,
	}

	if t.Updated.Valid {
		model.Updated = t.Updated.Time
	}

	return model
}
