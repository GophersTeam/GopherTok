package svc

import (
	"GopherTok/common/consts"
	"GopherTok/common/utils"
	"GopherTok/server/comment/model"
	"GopherTok/server/comment/rpc/internal/config"
	"GopherTok/server/relation/rpc/relationrpc"
	"GopherTok/server/user/rpc/userclient"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config              config.Config
	CommentModel        model.CommentModel
	Snowflake           *snowflake.Node
	UserRpc             userclient.User
	SensitiveWordFilter utils.SensitiveWordFilter
	KafkaPusher         *kq.Pusher
	RedisClient         *redis.Redis
	RelationRpc         relationrpc.RelationRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.ChatMachineId)
	trie := utils.NewSensitiveTrie()
	go func() {
		// 从数据库中读取敏感词，采用异步的方式，不影响服务启动
		trie.AddWords([]string{"傻逼", "傻叉", "垃圾", "尼玛", "傻狗", "傻逼吧你", "他妈的", "他妈"})
	}()

	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(c.MongoConf.Url, c.MongoConf.DB, c.MongoConf.Collection),
		Snowflake:    snowflakeNode,
		//UserRpc:             userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		SensitiveWordFilter: trie,
		KafkaPusher:         kq.NewPusher(c.KafkaConf.Addrs, c.KafkaConf.Topic),
		RedisClient:         redis.MustNewRedis(c.RedisConf),
	}
}
