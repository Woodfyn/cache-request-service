package app

import (
	"context"

	"github.com/Woodfyn/cache-service/internal/app/repository/rdb"
	"github.com/Woodfyn/cache-service/internal/app/server"
	"github.com/Woodfyn/cache-service/internal/app/service"
	"github.com/Woodfyn/cache-service/internal/config"
	"github.com/Woodfyn/cache-service/pkg/rdbclient"
)

var (
	ctx = context.Background()
)

func Start() {
	//init config
	cfg := config.NewConfig("configs", "develop")

	//init redis
	redis, err := rdbclient.NewRDBFromConfig(ctx, &rdbclient.RDBConfig{
		Addr:     cfg.RedisConfig.Addr,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.RedisConfig.DB,
	})
	if err != nil {
		panic(err)
	}

	//init dependencies
	handler := server.NewHandler(service.NewCache(rdb.NewCache(redis)))

	//init grpc server
	server.NewServer(handler)
}
