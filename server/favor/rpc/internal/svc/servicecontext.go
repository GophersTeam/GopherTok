package svc

import (
	"GopherTok/common/mock"
	"GopherTok/server/favor/model"
	"GopherTok/server/favor/rpc/internal/config"
	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config     config.Config
	VideoRpc   mock.VideoRpc
	FavorModel model.FavorModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 尚未完善
	cnf := redis.Options{
		Network:               "",
		Addr:                  "",
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              "",
		Password:              "",
		CredentialsProvider:   nil,
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
		VideoRpc:   mock.VideoRpc{},
		FavorModel: model.NewFavorModel(cnf),
	}
}
