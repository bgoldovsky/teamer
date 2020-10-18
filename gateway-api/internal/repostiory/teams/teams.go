package teams

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/gomodule/redigo/redis"
)

const (
	teamsKey = "teams"
)

var (
	ErrTeamsNotFound = errors.New("teams not found")
)

type Repository interface {
	Save(teams []models.TeamView) error
	Get() ([]models.TeamView, error)
	Clear() error
}

type repository struct {
	pool *redis.Pool
}

func NewRepository(address string) (*repository, error) {
	pool, err := newPool(address)
	if err != nil {
		return nil, err
	}

	return &repository{
		pool: pool,
	}, nil
}

func newPool(address string) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:   50,
		MaxActive: 10000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address)
			return conn, err
		},
	}, nil
}

func (r *repository) Save(teams []models.TeamView) error {
	conn := r.pool.Get()
	defer conn.Close()

	bytes, err := json.Marshal(teams)
	if err != nil {
		return fmt.Errorf("marshal teams error: %w. data: %v", err, bytes)
	}

	_, err = conn.Do("SET", teamsKey, bytes)
	if err != nil {
		return fmt.Errorf("save teams error: %w", err)
	}

	return nil
}

func (r *repository) Get() ([]models.TeamView, error) {
	conn := r.pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", teamsKey))
	if err == redis.ErrNil {
		return nil, ErrTeamsNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("retrieve teams error: %w", err)
	}

	var teams []models.TeamView
	err = json.Unmarshal([]byte(s), &teams)
	if err != nil {
		return nil, fmt.Errorf("unmarshal teams error: %w. data: %v", err, teams)
	}

	return teams, nil
}

func (r *repository) Clear() error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", teamsKey)
	if err != nil {
		return fmt.Errorf("clear teams error: %w", err)
	}

	return nil
}
