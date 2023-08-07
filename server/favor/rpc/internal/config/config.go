package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	RedisConf struct {
		Host string
		Pass string
	}

	VideoRpcConf zrpc.RpcClientConf
}
