package svc

import (
	"GopherTok/common/consts"
	"GopherTok/common/init_db"
	"GopherTok/server/favor/rpc/favorrpc"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/relation/rpc/relationrpc"
	"GopherTok/server/user/model"
	"GopherTok/server/user/rpc/internal/config"
	"GopherTok/server/video/rpc/videoclient"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	RelationRpc    pb.RelationRpcClient
	VideoRpc       videoclient.Video
	FavorRpc       favorrpc.FavorRpc
	Snowflake      *snowflake.Node
	Rdb            *redis.ClusterClient
	MysqlDb        *gorm.DB
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	snowflakeNode, _ := snowflake.NewNode(consts.UserMachineId)
	mysqlDb.AutoMigrate(&model.User{})
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	return &ServiceContext{
		Config:         c,
		RelationRpc:    relationrpc.NewRelationRpc(zrpc.MustNewClient(c.RelationRpcConf)),
		VideoRpc:       videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		FavorRpc:       favorrpc.NewFavorRpc(zrpc.MustNewClient(c.FavorRpcConf)),
		Snowflake:      snowflakeNode,
		MysqlDb:        mysqlDb,
		Rdb:            redisDb,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
