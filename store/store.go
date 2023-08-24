package store

import (
	"cars/store/postgres"
	"cars/store/redis"
	"fmt"
)

const Key = "store"

type Store struct {
	Postgres *postgres.Postgres
	Redis    *redis.Redis
}

func NewStore() (*Store, error) {
	p, err := postgres.NewPostgres()
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres: %v", err)
	}
	r, err := redis.NewRedis()
	if err != nil {
		return nil, fmt.Errorf("failed to create redis: %v", err)
	}
	return &Store{
		Postgres: p,
		Redis:    r,
	}, nil
}
