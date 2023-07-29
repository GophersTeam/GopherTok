package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"GopherTok/server/chat/api/internal/config"
	"GopherTok/server/chat/api/internal/handler"
	"GopherTok/server/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	httpx.SetValidator(svc.NewValidator())

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 取消检测日志打印
	logx.DisableStat()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
