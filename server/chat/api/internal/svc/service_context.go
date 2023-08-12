package svc

import (
	"GopherTok/server/chat/api/internal/config"
	"GopherTok/server/chat/api/internal/middleware"
	"GopherTok/server/chat/rpc/chatrpc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Auth    rest.Middleware
	ChatRpc chatrpc.ChatRpc
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:  c,
		Auth:    middleware.NewAuthMiddleware(c).Handle,
		ChatRpc: chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf)),
	}
}
