package svc

import (
	"GopherTok/common/init_db"
	"GopherTok/server/chat/rpc/chatrpc"
	"GopherTok/server/relation/rpc/internal/config"
	"GopherTok/server/user/rpc/userclient"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	Rdb            *redis.ClusterClient
	MysqlDb        *gorm.DB
	UserRpc        userclient.User
	ChatRpc        chatrpc.ChatRpc
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	rc := make([]string, 1)
	rc = append(rc, c.RedisConf.Cluster1,
		c.RedisConf.Cluster2,
		c.RedisConf.Cluster3,
		c.RedisConf.Cluster4,
		c.RedisConf.Cluster5,
		c.RedisConf.Cluster6)
	rdb := init_db.InitRedis(rc)
	return &ServiceContext{
		Config:         c,
		MysqlDb:        mysqlDb,
		Rdb:            rdb,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ChatRpc:        chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpc)),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
