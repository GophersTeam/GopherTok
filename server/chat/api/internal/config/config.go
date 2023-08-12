package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ChatRpcConf zrpc.RpcClientConf
	Token       struct {
		AccessToken  string
		RefreshToken string
	}
}
