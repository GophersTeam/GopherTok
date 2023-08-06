package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	KafkaConf kq.KqConf

	RedisConf redis.RedisConf

	MongoConf struct {
		Url        string
		DB         string
		Collection string
	}
}
