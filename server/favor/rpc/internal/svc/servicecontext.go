package svc

import (
	"GopherTok/common/init_db"
	"GopherTok/server/favor/model"
	"GopherTok/server/favor/rpc/internal/config"
	"GopherTok/server/video/rpc/videoclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//VideoRpc mock.VideoRpc
	VideoRpc   videoclient.Video
	FavorModel model.FavorModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//cnf := redis.Options{
	//	Addr:     c.RedisConf.Host,
	//	Password: "",
	//}
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	return &ServiceContext{
		Config: c,
		//VideoRpc:   mock.VideoRpc{},
		VideoRpc:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		FavorModel: model.NewFavorModel(redisDb),
	}
}
