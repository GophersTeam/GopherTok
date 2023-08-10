package main

import (
	"flag"
	"fmt"

	"GopherTok/server/relation/api/internal/config"
	"GopherTok/server/relation/api/internal/handler"
	"GopherTok/server/relation/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the config file")

func main() {
	flag.Parse()
	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
