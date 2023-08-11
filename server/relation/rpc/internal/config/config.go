package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	UserRpc zrpc.RpcClientConf
	ChatRpc zrpc.RpcClientConf
	Mysql   struct {
		DataSource string
	}
	RedisConf struct {
		Cluster1 string
		Cluster2 string
		Cluster3 string
		Cluster4 string
		Cluster5 string
		Cluster6 string
	}

	KqPusherRedisConf struct {
		Brokers []string
		Topic   string
	}
	KqPusherMysqlConf struct {
		Brokers []string
		Topic   string
	}
}
