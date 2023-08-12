package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	//RedisConf struct {
	//	Host string
	//	Pass string
	//}

	RedisCluster struct {
		Cluster1 string
		Cluster2 string
		Cluster3 string
		Cluster4 string
		Cluster5 string
		Cluster6 string
	}

	VideoRpcConf zrpc.RpcClientConf
}
