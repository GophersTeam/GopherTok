package svc

import (
	"GopherTok/common/init_db"
	"GopherTok/server/relation/rpc/internal/config"
	"GopherTok/server/user/rpc/userclient"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	Rdb     *redis.ClusterClient
	MysqlDb *gorm.DB
	UserRpc userclient.User
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
	fmt.Println("连接redis数据库成功")
	return &ServiceContext{
		Config:  c,
		MysqlDb: mysqlDb,
		Rdb:     rdb,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
