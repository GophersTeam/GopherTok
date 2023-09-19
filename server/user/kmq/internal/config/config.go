package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	KqConsumerConf kq.KqConf
	service.ServiceConf
	CacheRedis   cache.CacheConf
	MysqlCluster struct {
		DataSource string
	}
}
