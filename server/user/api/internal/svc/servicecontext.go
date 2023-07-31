package svc

import (
	"GopherTok/server/user/api/internal/config"
	"GopherTok/server/user/api/internal/middleware"
	"GopherTok/server/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	JWT     rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		JWT:     middleware.NewJWTMiddleware(c).Handle,
	}
}
