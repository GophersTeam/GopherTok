package main

import (
	"GopherTok/common/logs/zapx"
	"GopherTok/common/response/rpcserver"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"GopherTok/server/user/rpc/internal/config"
	"GopherTok/server/user/rpc/internal/server"
	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

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
