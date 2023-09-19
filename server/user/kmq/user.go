package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"GopherTok/common/errorx"
	"GopherTok/common/logs/zapx"
	"GopherTok/server/user/kmq/internal/config"
	"GopherTok/server/user/kmq/internal/service"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

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

	srv := service.NewService(c)
	queue := kq.MustNewQueue(c.KqConsumerConf, kq.WithHandle(srv.Consume))
	defer queue.Stop()
	// 自定义错误
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
	fmt.Println("seckill started!!!")
	queue.Start()
}
