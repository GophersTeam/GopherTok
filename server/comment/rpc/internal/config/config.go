package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MongoConf struct {
		Url        string
		DB         string
		Collection string
	}
	UserRpcConf zrpc.RpcClientConf
}
