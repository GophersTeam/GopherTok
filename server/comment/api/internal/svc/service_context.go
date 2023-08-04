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
	redisClient := redis.MustNewRedis(c.RedisConf)

	return &ServiceContext{
		Config:     c,
		Auth:       middleware.NewAuthMiddleware(redisClient).Handle,
		CommentRpc: commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
