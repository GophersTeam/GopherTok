package svc

import (
	"GopherTok/server/chat/rpc/chatrpc"
	"GopherTok/server/relation/api/internal/config"
	"GopherTok/server/relation/api/internal/middleware"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/relation/rpc/relationrpc"
	"GopherTok/server/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RelationRpc pb.RelationRpcClient
	UserRpc     userclient.User
	//UserRpc mock.UserRpc
	ChatRpc chatrpc.ChatRpc
	Jwt     rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Jwt:         middleware.NewJwtMiddleware(c).Handle,
		RelationRpc: relationrpc.NewRelationRpc(zrpc.MustNewClient(c.RelationRpc)),
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		//UserRpc: mock.UserRpc{},
		ChatRpc: chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpc)),
	}
}
