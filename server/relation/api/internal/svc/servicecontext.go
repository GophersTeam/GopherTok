package svc

import (
	"GopherTok/server/relation/api/internal/config"
	"GopherTok/server/relation/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Jwt    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Jwt:    middleware.NewJwtMiddleware().Handle,
	}
}
