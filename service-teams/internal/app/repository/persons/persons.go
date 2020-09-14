package persons

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
	"github.com/jackc/pgx/v4"
)

var (
	ErrPersonNotFount = errors.New("person not found")
)

type person struct {
	ID         sql.NullInt64
	FirstName  sql.NullString
	MiddleName sql.NullString
	LastName   sql.NullString
	Birthday   sql.NullTime
	Email      sql.NullString
	Phone      sql.NullString
	Slack      sql.NullString
	Role       sql.NullInt64
	TeamID     sql.NullInt64
	DutyOrder  sql.NullInt64
	IsActive   sql.NullBool
	Created    sql.NullTime
	Updated    sql.NullTime
}

type Repository interface {
	Save(ctx context.Context, person *models.Person) (*models.Person, error)
	Update(ctx context.Context, person *models.Person) (*models.Person, error)
	Remove(ctx context.Context, personID int64) (int64, error)
	Get(ctx context.Context, filter *v1.PersonFilter, limit, offset uint, sort, order string) ([]models.Person, error)
}

type Queryer interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type repository struct {
	database Queryer
}

func NewRepository(database Queryer) *repository {
	return &repository{
		database: database,
	}
}

func (r *repository) Save(ctx context.Context, person *models.Person) (*models.Person, error) {
	attributes := make(map[string]interface{})

	attributes["team_id"] = person.TeamID
	attributes["first_name"] = person.FirstName
	attributes["middle_name"] = person.MiddleName
	attributes["last_name"] = person.LastName
	attributes["birthday"] = person.Birthday
	attributes["email"] = person.Email
	attributes["phone"] = person.Phone
	attributes["role"] = person.Role
	attributes["slack"] = person.Slack
	attributes["duty_order"] = person.DutyOrder
	attributes["is_active"] = person.IsActive

	t, err := r.put(ctx, attributes)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *repository) Update(ctx context.Context, p *models.Person) (*models.Person, error) {
	var (
		i           uint8
		values      []interface{}
		placeholder []string
		returns     []string
		attributes  = make(map[string]interface{})
	)

	attributes["team_id"] = p.TeamID
	attributes["first_name"] = p.FirstName
	attributes["middle_name"] = p.MiddleName
	attributes["last_name"] = p.LastName
	attributes["birthday"] = p.Birthday
	attributes["email"] = p.Email
	attributes["phone"] = p.Phone
	attributes["slack"] = p.Slack
	attributes["role"] = p.Role
	attributes["duty_order"] = p.DutyOrder
	attributes["is_active"] = p.IsActive

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

	attributes["updated_at"] = time.Now().Local()
	for k, v := range attributes {
		i += 1
		values = append(values, v)
		placeholder = append(placeholder, fmt.Sprintf("%q = $%d", k, i))
	}

	template := `update "persons" set %s where id = %d returning %s`
	query := fmt.Sprintf(template, strings.Join(placeholder, ","), p.ID, strings.Join(returns, ","))

	var data person
	err := r.database.QueryRow(ctx, query, values...).Scan(
		&data.ID,
		&data.TeamID,
		&data.FirstName,
		&data.MiddleName,
		&data.LastName,
		&data.Birthday,
		&data.Email,
		&data.Phone,
		&data.Slack,
		&data.Role,
		&data.IsActive,
		&data.Updated,
		&data.Created)

	if err == pgx.ErrNoRows {
		return nil, ErrPersonNotFount
	}

	if err == nil {
		return nil, err
	}

	return data.convert(), err
}

func (r *repository) Remove(ctx context.Context, personID int64) (int64, error) {
	var query bytes.Buffer
	query.WriteString("delete from teams where id = $1;")
	_, err := r.database.Query(ctx, query.String(), personID)

	return personID, err
}

// TODO: do refactor strict!

func (r *repository) Get(
	ctx context.Context,
	filter *v1.PersonFilter,
	limit, offset uint,
	sort, order string,
) ([]models.Person, error) {
	var query bytes.Buffer

	ccc := []string{
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
		"created_at",
		"updated_at",
	}

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
		"created_at",
		"updated_at",
	}

	for i, c := range columns {
		columns[i] = fmt.Sprintf(`"p".%q`, c)
	}

	query.WriteString(fmt.Sprintf(`select %s from "persons" as "p"`, strings.Join(columns, ",")))

	where, args := r.where(filter)
	if where != "" {
		query.WriteString(" ")
		query.WriteString(where)
	}

	query.WriteString(fmt.Sprintf(` order by "".%q %s limit %d offset %d;`, order, sort, limit, offset))
	return r.query(ctx, ccc, query.String(), args...)
}

func (r *repository) put(ctx context.Context, attributes map[string]interface{}) (*models.Person, error) {
	var (
		i           uint8
		err         error
		values      []interface{}
		columns     []string
		returns     []string
		placeholder []string
	)

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

	for k, v := range attributes {
		i += 1
		values = append(values, v)
		columns = append(columns, fmt.Sprintf("%q", k))
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
	}

	query := fmt.Sprintf(`insert into "persons" (%s) values (%s) returning %s`,
		strings.Join(columns, ","), strings.Join(placeholder, ","), strings.Join(returns, ","))

	oo, err := r.query(ctx, returns, query, values...)
	if err != nil {
		return nil, err
	}

	return &oo[0], nil
}

func (r *repository) query(ctx context.Context, columns []string, query string, args ...interface{}) ([]models.Person, error) {
	rows, err := r.database.Query(ctx, query, args...)

	if err == pgx.ErrNoRows {
		return []models.Person{}, nil
	}

	if err != nil {
		return nil, err
	}

	var persons []models.Person
	for rows.Next() {
		var p person
		var dest []interface{}

		/*
			returns = append(returns, "")
			returns = append(returns, "updated_at")
			returns = append(returns, "created_at")
		*/

		for _, c := range columns {
			if c == `id` {
				dest = append(dest, &p.ID)
			}
			if c == `team_id` {
				dest = append(dest, &p.TeamID)
			}
			if c == `first_name` {
				dest = append(dest, &p.FirstName)
			}
			if c == `middle_name` {
				dest = append(dest, &p.MiddleName)
			}
			if c == `last_name` {
				dest = append(dest, &p.LastName)
			}
			if c == `birthday` {
				dest = append(dest, &p.Birthday)
			}
			if c == `email` {
				dest = append(dest, &p.Email)
			}
			if c == `phone` {
				dest = append(dest, &p.Phone)
			}
			if c == `slack` {
				dest = append(dest, &p.Slack)
			}
			if c == `role` {
				dest = append(dest, &p.Role)
			}
			if c == `duty_order` {
				dest = append(dest, &p.DutyOrder)
			}
			if c == `is_active` {
				dest = append(dest, &p.IsActive)
			}
			if c == `created_at` {
				dest = append(dest, &p.Created)
			}
			if c == `updated_at` {
				dest = append(dest, &p.Updated)
			}
		}

		if err := rows.Scan(dest...); err != nil {
			return nil, err
		}

		rp := p.convert()
		persons = append(persons, *rp)
	}

	return persons, nil
}

func (r *repository) where(f *v1.PersonFilter) (string, []interface{}) {
	var (
		i           uint8
		where       []string
		values      []interface{}
		placeholder []string
	)

	if f == nil {
		return "", values
	}

	if f.PersonIds != nil {
		for _, id := range f.PersonIds {
			i += 1
			values = append(values, id)
			placeholder = append(placeholder, fmt.Sprintf("$%d", i))
		}

		if len(f.PersonIds) > 0 {
			where = append(where, fmt.Sprintf(`"p"."id" in (%s)`, strings.Join(placeholder, ",")))
			placeholder = nil
		}
	}

	if len(where) > 0 {
		return "where " + strings.Join(where, " and "), values
	}

	return "", values
}

func (p *person) convert() *models.Person {
	model := models.Person{
		ID:        p.ID.Int64,
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
		model.TeamID = &p.TeamID.Int64
	}

	if p.Birthday.Valid {
		model.Birthday = &p.Birthday.Time
	}

	if p.Updated.Valid {
		model.Updated = p.Updated.Time
	}

	return &model
}
