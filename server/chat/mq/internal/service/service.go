package service

import (
	"GopherTok/common/consts"
	"GopherTok/server/chat/model"
	"GopherTok/server/chat/mq/internal/config"
	"GopherTok/server/chat/rpc/pb"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/protobuf/proto"
	"strconv"
)

type Service struct {
	Config       config.Config
	SqlConn      sqlx.SqlConn
	MessageModel model.MessageModel
	RedisClient  *redis.Redis
}

func NewService(c config.Config) *Service {
	mysqlConn := sqlx.NewMysql(c.MySQLConf.DataSource)
	return &Service{
		Config:       c,
		MessageModel: model.NewMessageModel(mysqlConn, c.CacheRedis),
		SqlConn:      mysqlConn,
		RedisClient:  redis.MustNewRedis(c.RedisConf),
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Info("成功消费消息")
	var message model.Message
	err := jsonx.UnmarshalFromString(value, &message)
	if err != nil {
		logx.Errorf("MessageAction jsonx.UnmarshalFromString error: %s", err.Error())
		return err
	}

	_, err = s.MessageModel.Insert(context.Background(), &message)
	if err != nil {
		logx.Errorf("MessageAction MessageModel.Insert error: %s", err.Error())
		return err
	}

	// 保存最新消息到redis
	fromUserID := strconv.Itoa(int(message.FromUserId))
	toUserID := strconv.Itoa(int(message.ToUserId))

	lastMessage := &pb.LastMessage{Content: message.Content}
	lastMessage.MsgType = consts.MsgTypeRecv
	lastMessageRecvBytes, _ := proto.Marshal(lastMessage)
	//lastMessageRecvStr, _ := jsonx.MarshalToString(lastMessage)
	lastMessage.MsgType = consts.MsgTypeSend
	lastMessageSendBytes, _ := proto.Marshal(lastMessage)
	//lastMessageSendStr, err := jsonx.MarshalToString(lastMessage)

	_ = s.RedisClient.Hset(consts.LastMessagePrefix+fromUserID, toUserID, string(lastMessageSendBytes))
	_ = s.RedisClient.Hset(consts.LastMessagePrefix+toUserID, fromUserID, string(lastMessageRecvBytes))

	return nil
}
