package teams

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	v1 "github.com/bgoldovsky/teamer-bot/service-people/internal/generated/rpc/v1"
	"github.com/bgoldovsky/teamer-bot/service-people/internal/models"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	Save(ctx context.Context, team *models.Team) (*models.Team, error)
	Update(ctx context.Context, team *models.Team) (*models.Team, error)
	Remove(ctx context.Context, teamID int64) (int64, error)
	Query(ctx context.Context, filter *v1.TeamFilter, limit, offset uint, sort, order string) ([]models.Team, error)
}

type Queryer interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type RepositoryPgsql struct {
	database Queryer
}

func NewRepository(database Queryer) Repository {
	return &RepositoryPgsql{
		database: database,
	}
}

func (r *RepositoryPgsql) Save(ctx context.Context, team *models.Team) (*models.Team, error) {
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

func (r *RepositoryPgsql) Update(ctx context.Context, team *models.Team) (*models.Team, error) {
	var (
		i           uint8
		values      []interface{}
		placeholder []string
		returns     []string
		attributes  = make(map[string]interface{})
	)

	attributes["name"] = team.Name
	attributes["description"] = team.Description
	attributes["slack"] = team.Slack

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
	query := fmt.Sprintf(template, strings.Join(placeholder, ","), team.ID, strings.Join(returns, ","))

	var replyTeam models.Team
	err := r.database.QueryRow(ctx, query, values...).Scan(
		&replyTeam.ID,
		&replyTeam.Name,
		&replyTeam.Description,
		&replyTeam.Slack,
		&replyTeam.Updated,
		&replyTeam.Created)

	return &replyTeam, err
}

func (r *RepositoryPgsql) Remove(ctx context.Context, teamID int64) (int64, error) {
	var query bytes.Buffer
	query.WriteString("delete from teams where id = $1;")
	_, err := r.database.Query(ctx, query.String(), teamID)

	return teamID, err
}

func (r *RepositoryPgsql) Query(
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

func (r *RepositoryPgsql) put(ctx context.Context, attributes map[string]interface{}) (*models.Team, error) {
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

func (r *RepositoryPgsql) query(ctx context.Context, columns []string, query string, args ...interface{}) ([]models.Team, error) {
	rows, err := r.database.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var tt []models.Team

	for rows.Next() {
		var t models.Team
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

		tt = append(tt, t)
	}

	return tt, nil
}

func (r *RepositoryPgsql) where(f *v1.TeamFilter) (string, []interface{}) {
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
