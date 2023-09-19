package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RelationRpcConf zrpc.RpcClientConf
	VideoRpcConf    zrpc.RpcClientConf
	FavorRpcConf    zrpc.RpcClientConf
	CacheRedis      cache.CacheConf

	MysqlCluster struct {
		DataSource string
	}
	RedisCluster struct {
		Cluster1 string
		Cluster2 string
		Cluster3 string
		Cluster4 string
		Cluster5 string
		Cluster6 string
	}
	Token struct {
		AccessToken  string
		RefreshToken string
	}
	Salt         string
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
