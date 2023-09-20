package svc

import (
	"GopherTok/common/init_db"
	"GopherTok/server/favor/model"
	"GopherTok/server/favor/rpc/internal/config"
	"GopherTok/server/user/rpc/userclient"
	"GopherTok/server/video/rpc/videoclient"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	VideoRpc    videoclient.Video
	FavorModel  model.FavorModel
	UserRpc     userclient.User
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	return &ServiceContext{
		KafkaPusher: kq.NewPusher(c.KafkaConf.Addrs, c.KafkaConf.Topic),
		Config:      c,
		VideoRpc:    videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		FavorModel:  model.NewFavorModel(redisDb),
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
