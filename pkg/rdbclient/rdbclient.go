package rdbclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RDBConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRDBFromConfig(ctx context.Context, config *RDBConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
