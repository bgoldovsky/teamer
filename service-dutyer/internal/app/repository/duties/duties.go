//go:generate mockgen -destination duties_mock/duties_mock.go -source duties.go

package duties

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	pgx "github.com/jackc/pgx/v4"
)

const (
	orderNotConsistent = "null value in column \"duty_order\""
)

var (
	ErrInvalidTeam  = errors.New("invalid team or persons ID")
	ErrDutyNotFound = errors.New("duty not found")
)

type Repository interface {
	Get(ctx context.Context, teamID int64) (*models.Duty, error)
	Save(ctx context.Context, duty *models.Duty) error
	Swap(ctx context.Context, teamID, firstPersonID, secondPersonID int64) error
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

func (r *repository) Swap(ctx context.Context, teamID, firstPersonID, secondPersonID int64) error {
	query := `update persons p set duty_order = case id
		when $1 then (select duty_order from persons where id = $2 and team_id = $3)
		when $2 then (select duty_order from persons where id = $1 and team_id = $3)
		end
		where id in ($1,$2);`

	_, err := r.database.Query(ctx, query, firstPersonID, secondPersonID, teamID)
	if err != nil && strings.Contains(err.Error(), orderNotConsistent) {
		return ErrInvalidTeam
	}

	return err
}

func (r *repository) Get(ctx context.Context, teamID int64) (*models.Duty, error) {
	columns := []string{
		`"d"."team_id"`,
		`"d"."person_id"`,
		`"p"."first_name"`,
		`"p"."last_name"`,
		`"p"."slack"`,
		`"t"."slack"`,
		`"p"."duty_order"`,
		`"d"."month"`,
		`"d"."day"`,
		`"d"."updated_at"`,
		`"d"."created_at"`,
	}

	// Формирование запроса
	query := fmt.Sprintf(`select %s from "duties" as "d" left join persons as "p" on "d"."person_id" = "p"."id" left join "teams" as "t" on "p".team_id = "t"."id" where "d"."team_id" = $1;`,
		strings.Join(columns, ","))

	// Выполнение запроса
	row := r.database.QueryRow(ctx, query, teamID)
	var d duty
	err := row.Scan(
		&d.TeamID,
		&d.PersonID,
		&d.FirstName,
		&d.LastName,
		&d.Slack,
		&d.Channel,
		&d.DutyOrder,
		&d.Month,
		&d.Day,
		&d.Updated,
		&d.Created,
	)

	if isEmpty(err) {
		return nil, ErrDutyNotFound
	}

	duty := d.convert()
	return &duty, nil
}

func (r *repository) Save(ctx context.Context, d *models.Duty) error {
	var (
		i           uint8
		err         error
		values      []interface{}
		columns     []string
		placeholder []string
		update      []string
	)

	attributes := map[string]interface{}{
		"team_id":   d.TeamID,
		"person_id": d.PersonID,
		"month":     d.Month,
		"day":       d.Day,
	}

	for k, v := range attributes {
		i += 1
		values = append(values, v)
		columns = append(columns, fmt.Sprintf("%q", k))
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
		update = append(update, fmt.Sprintf("%s=excluded.%s", k, k))
	}

	query := fmt.Sprintf(`insert into "duties" (%s) values (%s) on conflict (team_id) do update set %s`,
		strings.Join(columns, ","),
		strings.Join(placeholder, ","),
		strings.Join(update, ","))

	_, err = r.database.Query(ctx, query, values...)
	return err
}

type duty struct {
	TeamID    sql.NullInt64
	PersonID  sql.NullInt64
	FirstName sql.NullString
	LastName  sql.NullString
	Slack     sql.NullString
	Channel   sql.NullString
	DutyOrder sql.NullInt64
	Month     sql.NullInt64
	Day       sql.NullInt64
	Created   sql.NullTime
	Updated   sql.NullTime
}

func (d *duty) convert() models.Duty {
	model := models.Duty{
		TeamID:    d.TeamID.Int64,
		PersonID:  d.PersonID.Int64,
		FirstName: d.FirstName.String,
		LastName:  d.LastName.String,
		Slack:     d.Slack.String,
		Channel:   d.Channel.String,
		Order:     d.DutyOrder.Int64,
		Month:     time.Month(d.Month.Int64),
		Day:       d.Day.Int64,
		Created:   d.Created.Time,
	}

	if d.Updated.Valid {
		model.Updated = d.Updated.Time
	}

	return model
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
