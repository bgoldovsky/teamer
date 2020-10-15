package persons

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/gomodule/redigo/redis"
)

const (
	personsKey = "persons"
)

var (
	ErrPersonsNotFound = errors.New("persons not found")
)

type Repository interface {
	Save(persons []models.PersonView) error
	Get() ([]models.PersonView, error)
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

func (r *repository) Save(persons []models.PersonView) error {
	conn := r.pool.Get()
	defer conn.Close()

	bytes, err := json.Marshal(persons)
	if err != nil {
		return fmt.Errorf("marshal persons error: %w. data: %v", err, bytes)
	}

	_, err = conn.Do("SET", personsKey, bytes)
	if err != nil {
		return fmt.Errorf("save persons error: %w", err)
	}

	return nil
}

func (r *repository) Get() ([]models.PersonView, error) {
	conn := r.pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", personsKey))
	if err == redis.ErrNil {
		return nil, ErrPersonsNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("retrieve persons error: %w", err)
	}

	var persons []models.PersonView
	err = json.Unmarshal([]byte(s), &persons)
	if err != nil {
		return nil, fmt.Errorf("unmarshal persons error: %w. data: %v", err, persons)
	}

	return persons, nil
}

func (r *repository) Clear() error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", personsKey)
	if err != nil {
		return fmt.Errorf("clear persons error: %w", err)
	}

	return nil
}
