package main

import (
	"GopherTok/common/logs/zapx"
	"GopherTok/common/response/rpcserver"
	"GopherTok/server/favor/rpc/internal/config"
	"GopherTok/server/favor/rpc/internal/server"
	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/favor.yaml", "the config file")

func main() {
	flag.Parse()

	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)
	nacosConf.ListenConfig(func(namespace, group, dataId, data string) {
		fmt.Printf("配置文件发生变化\n")
		fmt.Printf("namespace: %s, group: %s, dataId: %s, data: %s", namespace, group, dataId, data)
	})

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		favor.RegisterFavorRpcServer(grpcServer, server.NewFavorRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 添加rpc 错误返回
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	// zap
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
