package svc

import (
	"GopherTok/server/user/rpc/userclient"
	"GopherTok/server/video/api/internal/config"
	"GopherTok/server/video/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	JWT     rest.Middleware
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		JWT:     middleware.NewJWTMiddleware(c).Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
