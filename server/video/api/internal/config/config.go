package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf    zrpc.RpcClientConf
	VideoRpcConf   zrpc.RpcClientConf
	CommentRpcConf zrpc.RpcClientConf
	FavorRpcConf   zrpc.RpcClientConf
	Token          struct {
		AccessToken  string
		RefreshToken string
	}
	MinioCluster struct {
		Endpoint  string
		AccessKey string
		SecretKey string
	}
	CurrentStoreType int
	VideReflection   struct {
		Host string
		Port string
	}
	TencentCOS struct {
		Url       string
		SecretId  string
		SecretKey string
	}
	RedisCluster struct {
		Cluster1 string
		Cluster2 string
		Cluster3 string
		Cluster4 string
		Cluster5 string
		Cluster6 string
	}
}
