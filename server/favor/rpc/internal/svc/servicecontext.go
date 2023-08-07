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
	cnf := redis.Options{
		Addr:     c.RedisConf.Host,
		Password: "",
	}
	return &ServiceContext{
		Config:     c,
		VideoRpc:   mock.VideoRpc{},
		FavorModel: model.NewFavorModel(cnf),
	}
}
