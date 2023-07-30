package main

import (
	"GopherTok/server/chat/rpc/internal/config"
	"GopherTok/server/chat/rpc/internal/server"
	"GopherTok/server/chat/rpc/internal/svc"
	"GopherTok/server/chat/rpc/pb"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the config file")

func main() {
	flag.Parse()

	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterChatRpcServer(grpcServer, server.NewChatRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logx.DisableStat()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()

}

//func solveNQueens(n int) [][]string {
//	res := [][]string{}
//	// 初始化棋盘
//	board := make([][]string, n)
//	for i := 0; i < n; i++ {
//		board[i] = make([]string, n)
//		for j := 0; j < n; j++ {
//			board[i][j] = ""
//		}
//	}
//
//	var dfs func(board [][]string, row int)  {
//		if row == len(board) {
//			temp := make([]string, len(board))
//			for i := 0; i < len(board); i++ {
//				temp[i] = strings.Join(board[i], "")
//			}
//			res = append(res, temp)
//			return
//		}
//	}
//
//
//
//	dfs(board, 0, &res)
//	return res
//}
