package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf  zrpc.RpcClientConf
	VideoRpcConf zrpc.RpcClientConf
	Token        struct {
		AccessToken  string
		RefreshToken string
	}
	MinioCluster struct {
		Endpoint  string
		AccessKey string
		SecretKey string
	}
	CurrentStoreType int
}
