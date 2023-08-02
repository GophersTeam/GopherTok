package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type Config struct {
	rest.RestConf
	Redis       redis.RedisConf
	Mysql       *gorm.DB
	UserRpc     zrpc.RpcClientConf
	RelationRpc zrpc.RpcClientConf
}
