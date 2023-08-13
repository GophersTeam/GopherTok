package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	KqConsumerConf kq.KqConf
	service.ServiceConf
	MysqlCluster struct {
		DataSource string
	}
}
