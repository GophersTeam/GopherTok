package svc

import (
	"GopherTok/server/relation/api/internal/config"
	"GopherTok/server/relation/api/internal/middleware"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/relation/rpc/relationRpc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RelationRpc pb.RelationRpcClient
	Jwt         rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Jwt:         middleware.NewJwtMiddleware(c).Handle,
		RelationRpc: relationrpc.NewRelationRpc(zrpc.MustNewClient(c.RelationRpc)),
	}
}
