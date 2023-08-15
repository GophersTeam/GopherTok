package svc

import (
	"GopherTok/server/comment/rpc/commentrpc"
	"GopherTok/server/favor/api/internal/config"
	"GopherTok/server/favor/api/internal/middleware"
	"GopherTok/server/favor/rpc/favorrpc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	JWT       rest.Middleware
	FavorRpc  favorrpc.FavorRpc
	CommenRpc commentrpc.CommentRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		JWT:       middleware.NewJWTMiddleware(c).Handle,
		FavorRpc:  favorrpc.NewFavorRpc(zrpc.MustNewClient(c.FavorRpcConf)),
		CommenRpc: commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
