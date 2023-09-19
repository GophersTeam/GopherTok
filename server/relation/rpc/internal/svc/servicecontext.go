package svc

import (
	"GopherTok/common/init_db"
	"GopherTok/server/chat/rpc/chatrpc"
	"GopherTok/server/relation/rpc/internal/config"
	"GopherTok/server/user/rpc/userclient"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type ServiceContext struct {
	Config              config.Config
	Rdb                 *redis.ClusterClient
	MysqlDb             *gorm.DB
	UserRpc             userclient.User
	ChatRpc             chatrpc.ChatRpc
	KqPusherRedisClient *kq.Pusher
	KqPusherMysqlClient *kq.Pusher
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
		Config:              c,
		MysqlDb:             mysqlDb,
		Rdb:                 rdb,
		UserRpc:             userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ChatRpc:             chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpc)),
		KqPusherRedisClient: kq.NewPusher(c.KqPusherRedisConf.Brokers, c.KqPusherRedisConf.Topic),
		KqPusherMysqlClient: kq.NewPusher(c.KqPusherMysqlConf.Brokers, c.KqPusherMysqlConf.Topic),
	}
}
func InitGorm(MysqlDataSourece string) *gorm.DB {
	// 将日志写进kafka
	//logx.SetWriter(*LogxKafka())
	DB, err := gorm.Open(mysql.Open(MysqlDataSourece),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	db, _ := DB.DB()
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(200)
	db.SetConnMaxLifetime(time.Minute)
	return DB
}
