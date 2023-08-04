package svc

import (
	"GopherTok/common/consts"
	"GopherTok/common/init_db"
	"GopherTok/server/video/model"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-queue/kq"

	"GopherTok/server/video/rpc/internal/config"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	Snowflake      *snowflake.Node
	Rdb            *redis.ClusterClient
	MysqlDb        *gorm.DB
	MinioDb        *minio.Client
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	snowflakeNode, _ := snowflake.NewNode(consts.VideoMachineId)
	mysqlDb.AutoMigrate(&model.Video{})
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	minioDb := init_db.InitMinio(c.MinioCluster.Endpoint, c.MinioCluster.AccessKey, c.MinioCluster.SecretKey)

	return &ServiceContext{
		Config:         c,
		Snowflake:      snowflakeNode,
		MysqlDb:        mysqlDb,
		Rdb:            redisDb,
		MinioDb:        minioDb,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
