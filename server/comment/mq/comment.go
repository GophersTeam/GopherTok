package main

import (
	"flag"
	"fmt"

	"GopherTok/server/comment/mq/internal/config"
	"GopherTok/server/comment/mq/internal/service"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the etc file")

func main() {
	flag.Parse()

	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)

	s := service.NewService(c)

	queue := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(s.Consume))
	defer queue.Stop()

	fmt.Println("comment-mq started!!!")
	queue.Start()
}
