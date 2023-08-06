package svc

import (
	"GopherTok/server/favor/model"
	"GopherTok/server/favor/rpc/internal/config"
	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config config.Config

	FavorModel model.FavorModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conf := redis.Options{
		Addr:                  "",
		Username:              "",
		Password:              "",
		DB:                    0,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
	}

	return &ServiceContext{
		Config:     c,
		FavorModel: model.NewFavorModel(conf),
	}
}
