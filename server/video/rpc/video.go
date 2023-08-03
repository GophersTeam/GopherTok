package main

import (
	"flag"
	"fmt"

	"GopherTok/server/video/rpc/internal/config"
	"GopherTok/server/video/rpc/internal/server"
	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/video.yaml", "the config file")

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
		video.RegisterVideoServer(grpcServer, server.NewVideoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
