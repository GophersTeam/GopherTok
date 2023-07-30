package main

import (
	"GopherTok/server/chat/api/internal/config"
	"GopherTok/server/chat/api/internal/handler"
	"GopherTok/server/chat/api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the config file")

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
