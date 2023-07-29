package svc

import (
	"GopherTok/server/chat/api/internal/config"
	"GopherTok/server/chat/api/internal/middleware"
	"GopherTok/server/chat/rpc/chatrpc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Auth        rest.Middleware
	RedisClient *redis.Redis
	ChatRpc     chatrpc.ChatRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.MustNewRedis(c.RedisConf)

	return &ServiceContext{
		Config:      c,
		Auth:        middleware.NewAuthMiddleware(redisClient).Handle,
		RedisClient: redisClient,
		ChatRpc:     chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf)),
	}
}
