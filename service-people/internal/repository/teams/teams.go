package teams

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	v1 "github.com/bgoldovsky/teamer-bot/service-people/internal/generated/rpc/v1"
	"github.com/bgoldovsky/teamer-bot/service-people/internal/models"
	"github.com/jackc/pgx/v4"
)

var (
	ErrTeamNotFount = errors.New("team not found")
)

type team struct {
	ID          sql.NullInt64
	Name        sql.NullString
	Description sql.NullString
	Slack       sql.NullString
	Created     sql.NullTime
	Updated     sql.NullTime
}

type Repository interface {
	Save(ctx context.Context, team *models.Team) (*models.Team, error)
	Update(ctx context.Context, team *models.Team) (*models.Team, error)
	Remove(ctx context.Context, teamID int64) (int64, error)
	Get(ctx context.Context, filter *v1.TeamFilter, limit, offset uint, sort, order string) ([]models.Team, error)
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

func (r *repository) Save(ctx context.Context, team *models.Team) (*models.Team, error) {
	attributes := make(map[string]interface{})
	attributes["name"] = team.Name
	attributes["description"] = team.Description
	attributes["slack"] = team.Slack

	t, err := r.put(ctx, attributes)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *repository) Update(ctx context.Context, t *models.Team) (*models.Team, error) {
	var (
		i           uint8
		values      []interface{}
		placeholder []string
		returns     []string
		attributes  = make(map[string]interface{})
	)

	attributes["name"] = t.Name
	attributes["description"] = t.Description
	attributes["slack"] = t.Slack

	returns = append(returns, "id")
	returns = append(returns, "name")
	returns = append(returns, "description")
	returns = append(returns, "slack")
	returns = append(returns, "updated_at")
	returns = append(returns, "created_at")

	attributes["updated_at"] = time.Now().Local()
	for k, v := range attributes {
		i += 1
		values = append(values, v)
		placeholder = append(placeholder, fmt.Sprintf("%q = $%d", k, i))
	}

	template := `update "teams" set %s where id = %d returning %s`
	query := fmt.Sprintf(template, strings.Join(placeholder, ","), t.ID, strings.Join(returns, ","))

	var data team
	err := r.database.QueryRow(ctx, query, values...).Scan(
		&data.ID,
		&data.Name,
		&data.Description,
		&data.Slack,
		&data.Updated,
		&data.Created)

	if err == pgx.ErrNoRows {
		return nil, ErrTeamNotFount
	}

	if err == nil {
		return nil, err
	}

	return data.convert(), err
}

func (r *repository) Remove(ctx context.Context, teamID int64) (int64, error) {
	var query bytes.Buffer
	query.WriteString("delete from teams where id = $1;")
	_, err := r.database.Query(ctx, query.String(), teamID)

	return teamID, err
}

func (r *repository) Get(
	ctx context.Context,
	filter *v1.TeamFilter,
	limit, offset uint,
	sort, order string,
) ([]models.Team, error) {
	var query bytes.Buffer

	ccc := []string{"id", "name", "description", "slack", "created_at", "updated_at"}
	columns := []string{"id", "name", "description", "slack", "created_at", "updated_at"}

	for i, c := range columns {
		columns[i] = fmt.Sprintf(`"t".%q`, c)
	}

	query.WriteString(fmt.Sprintf(`select %s from "teams" as "t"`, strings.Join(columns, ",")))

	where, args := r.where(filter)
	if where != "" {
		query.WriteString(" ")
		query.WriteString(where)
	}

	query.WriteString(fmt.Sprintf(` order by "t".%q %s limit %d offset %d;`, order, sort, limit, offset))
	return r.query(ctx, ccc, query.String(), args...)
}

func (r *repository) put(ctx context.Context, attributes map[string]interface{}) (*models.Team, error) {
	var (
		i           uint8
		err         error
		values      []interface{}
		columns     []string
		returns     []string
		placeholder []string
	)

	returns = append(returns, "id")
	returns = append(returns, "name")
	returns = append(returns, "description")
	returns = append(returns, "slack")
	returns = append(returns, "created_at")
	returns = append(returns, "updated_at")

	for k, v := range attributes {
		i += 1
		values = append(values, v)
		columns = append(columns, fmt.Sprintf("%q", k))
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
	}

	query := fmt.Sprintf(`insert into "teams" (%s) values (%s) returning %s`,
		strings.Join(columns, ","), strings.Join(placeholder, ","), strings.Join(returns, ","))

	oo, err := r.query(ctx, returns, query, values...)
	if err != nil {
		return nil, err
	}

	return &oo[0], nil
}

func (r *repository) query(ctx context.Context, columns []string, query string, args ...interface{}) ([]models.Team, error) {
	rows, err := r.database.Query(ctx, query, args...)

	if err == pgx.ErrNoRows {
		return []models.Team{}, nil
	}

	if err != nil {
		return nil, err
	}

	var teams []models.Team
	for rows.Next() {
		var t team
		var dest []interface{}

		for _, c := range columns {
			if c == `id` {
				dest = append(dest, &t.ID)
			}
			if c == `name` {
				dest = append(dest, &t.Name)
			}
			if c == `description` {
				dest = append(dest, &t.Description)
			}
			if c == `slack` {
				dest = append(dest, &t.Slack)
			}
			if c == `created_at` {
				dest = append(dest, &t.Created)
			}
			if c == `updated_at` {
				dest = append(dest, &t.Updated)
			}
		}

		if err := rows.Scan(dest...); err != nil {
			return nil, err
		}

		rt := t.convert()
		teams = append(teams, *rt)
	}

	return teams, nil
}

func (r *repository) where(f *v1.TeamFilter) (string, []interface{}) {
	var (
		i           uint8
		where       []string
		values      []interface{}
		placeholder []string
	)

	if f == nil {
		return "", values
	}

	if f.Ids != nil {
		for _, id := range f.Ids {
			i += 1
			values = append(values, id)
			placeholder = append(placeholder, fmt.Sprintf("$%d", i))
		}

		if len(f.Ids) > 0 {
			where = append(where, fmt.Sprintf(`"t"."id" in (%s)`, strings.Join(placeholder, ",")))
			placeholder = nil
		}
	}

	if f.DateFrom != nil && f.DateFrom.Seconds > 0 {
		i += 1
		where = append(where, fmt.Sprintf(`"t"."created_at" > $%d`, i))
		values = append(values, time.Unix(f.DateFrom.Seconds, 0).Format(time.RFC3339Nano))
	}

	if f.DateTo != nil && f.DateTo.Seconds > 0 {
		i += 1
		where = append(where, fmt.Sprintf(`"t"."created_at" < $%d`, i))
		values = append(values, time.Unix(f.DateTo.Seconds, 0).Format(time.RFC3339Nano))
	}

	if len(where) > 0 {
		return "where " + strings.Join(where, " and "), values
	}

	return "", values
}

func (t *team) convert() *models.Team {
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

	return &model
}
