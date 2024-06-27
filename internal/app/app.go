package app

import (
	"context"

	"github.com/Woodfyn/cache-service/internal/config"
	"github.com/Woodfyn/cache-service/internal/handler"
	"github.com/Woodfyn/cache-service/internal/repository/rdb"
	"github.com/Woodfyn/cache-service/internal/server"
	"github.com/Woodfyn/cache-service/internal/service"
	"github.com/Woodfyn/cache-service/pkg/rdbclient"
	"github.com/Woodfyn/cache-service/pkg/signaler"
)

var (
	ctx = context.Background()
)

const (
	CFG_FOLDER = "configs"
	CFG_FILE   = "develop"
)

func Start() {
	//init config
	cfg := config.NewConfig(CFG_FOLDER, CFG_FILE)

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
	handler := handler.NewHandler(service.NewCache(rdb.NewCache(redis)))

	//init grpc server
	srv := server.NewServer(handler)

	// start grpc server
	go srv.Run(cfg.ServerConfig.Port)

	signaler.Wait()

	if err := redis.Close(); err != nil {
		panic(err)
	}
}
