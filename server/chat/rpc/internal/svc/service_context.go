package svc

import (
	"GopherTok/common/consts"
	"GopherTok/server/chat/model"
	"GopherTok/server/chat/rpc/internal/config"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	SqlConn      sqlx.SqlConn
	MessageModel model.MessageModel
	Snowflake    *snowflake.Node
	RedisClient  *redis.Redis
	KafkaPusher  *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQLConf.DataSource)
	snowflakeNode, _ := snowflake.NewNode(consts.ChatMachineId)

	return &ServiceContext{
		Config:       c,
		SqlConn:      mysqlConn,
		MessageModel: model.NewMessageModel(mysqlConn, c.CacheRedis),
		Snowflake:    snowflakeNode,
		RedisClient:  redis.MustNewRedis(c.RedisConf),
		KafkaPusher:  kq.NewPusher(c.KafkaConf.Addrs, c.KafkaConf.Topic),
	}
}
