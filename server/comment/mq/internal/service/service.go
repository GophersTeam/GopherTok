package service

import (
	"GopherTok/common/consts"
	"GopherTok/server/comment/model"
	"GopherTok/server/comment/mq/internal/config"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
)

type Service struct {
	Config       config.Config
	RedisClient  *redis.Redis
	CommentModel model.CommentModel
}

func NewService(c config.Config) *Service {
	return &Service{
		Config:       c,
		RedisClient:  redis.MustNewRedis(c.RedisConf),
		CommentModel: model.NewCommentModel(c.MongoConf.Url, c.MongoConf.DB, c.MongoConf.Collection),
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Info("成功消费消息")
	var comment model.Comment
	err := jsonx.UnmarshalFromString(value, &comment)
	if err != nil {
		logx.Errorf("MessageAction jsonx.UnmarshalFromString error: %s", err.Error())
		return err
	}

	err = s.CommentModel.Insert(context.Background(), &comment)
	if err != nil {
		logx.Errorf("MessageAction MessageModel.Insert error: %s", err.Error())
		return err
	}

	// 视频评论数+1
	res, err := s.RedisClient.Incr(consts.VideoCommentPrefix + strconv.Itoa(int(comment.VideoId)))
	if err != nil {
		logx.Errorf("MessageAction RedisClient.Incr error: %s", err.Error())
		return err
	}

	logx.Info("视频评论数+1", res)

	return nil
}
