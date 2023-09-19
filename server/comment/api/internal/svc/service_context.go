package svc

import (
	"GopherTok/server/comment/api/internal/config"
	"GopherTok/server/comment/api/internal/middleware"
	"GopherTok/server/comment/rpc/commentrpc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Auth        rest.Middleware
	RedisClient redis.Redis
	CommentRpc  commentrpc.CommentRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Auth:       middleware.NewAuthMiddleware(c).Handle,
		CommentRpc: commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
