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

var configFile = "D:\\GopherTok\\server\\relation\\api\\etc\\relation.yaml"

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
