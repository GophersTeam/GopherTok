package svc

import (
	"GopherTok/common/init_db"
	"GopherTok/server/relation/rpc/internal/config"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	Rdb     *redis.Redis
	MysqlDb *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	rdb := redis.MustNewRedis(c.RedisConf)
	fmt.Println("连接redis数据库成功")
	return &ServiceContext{
		Config:  c,
		MysqlDb: mysqlDb,
		Rdb:     rdb,
	}
}
