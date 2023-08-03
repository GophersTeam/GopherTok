package svc

import (
	"GopherTok/common/consts"
	"GopherTok/common/init_db"
	"GopherTok/server/user/rpc/userclient"
	"GopherTok/server/video/api/internal/config"
	"GopherTok/server/video/api/internal/middleware"
	"GopherTok/server/video/rpc/videoclient"
	"github.com/bwmarrin/snowflake"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	JWT       rest.Middleware
	UserRpc   userclient.User
	VideoRpc  videoclient.Video
	Snowflake *snowflake.Node
	MinioDb   *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.VideoMachineId)
	minioDb := init_db.InitMinio(c.MinioCluster.Endpoint, c.MinioCluster.AccessKey, c.MinioCluster.SecretKey)
	return &ServiceContext{
		Config:    c,
		JWT:       middleware.NewJWTMiddleware(c).Handle,
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		VideoRpc:  videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		Snowflake: snowflakeNode,
		MinioDb:   minioDb,
	}
}
