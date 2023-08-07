package svc

import (
	"GopherTok/common/consts"
	"GopherTok/common/init_db"
	"GopherTok/common/mock"
	"GopherTok/server/video/api/internal/config"
	"GopherTok/server/video/api/internal/middleware"
	"GopherTok/server/video/rpc/videoclient"
	"github.com/bwmarrin/snowflake"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	JWT        rest.Middleware
	VideoJWT   rest.Middleware
	UserRpc    mock.UserRpc
	VideoRpc   videoclient.Video
	FavorRpc   mock.FavorRpc
	CommentRpc mock.CommentRpc

	Snowflake *snowflake.Node
	MinioDb   *minio.Client
	Rdb       *redis.ClusterClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.VideoMachineId)
	minioDb := init_db.InitMinio(c.MinioCluster.Endpoint, c.MinioCluster.AccessKey, c.MinioCluster.SecretKey)
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	return &ServiceContext{
		Config:     c,
		JWT:        middleware.NewJWTMiddleware(c).Handle,
		VideoJWT:   middleware.NewVideoJWTMiddleware(c).Handle,
		UserRpc:    mock.UserRpc{},
		FavorRpc:   mock.FavorRpc{},
		VideoRpc:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		CommentRpc: mock.CommentRpc{},
		Snowflake:  snowflakeNode,
		MinioDb:    minioDb,
		Rdb:        redisDb,
	}
}
