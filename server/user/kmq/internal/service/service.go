package service

import (
	"GopherTok/common/utils"
	"GopherTok/server/user/kmq/internal/config"
	"GopherTok/server/user/model"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Service struct {
	c         config.Config // 配置信息
	Rdb       *redis.ClusterClient
	Log       logx.LogConf
	SqlConn   sqlx.SqlConn
	UserModel model.UserModel
}

// NewService 创建一个新的 Service 实例
func NewService(c config.Config) *Service {
	mysqlConn := sqlx.NewMysql(c.MysqlCluster.DataSource)
	// 创建 Service 实例
	s := &Service{
		c:         c,
		SqlConn:   mysqlConn,
		UserModel: model.NewUserModel(mysqlConn, c.CacheRedis),
	}

	return s
}

// consume 是消费者
func (s *Service) Consume(_ string, value string) error {

	var m model.User
	err := jsonx.UnmarshalFromString(value, &m)
	if err != nil {
		logx.Errorf("MessageAction jsonx.UnmarshalFromString error: %s", err.Error())
		return err
	}
	var avatar, backGroundImage, signature string
	// 并发请求网站
	mr.Finish(func() error {
		avatar = utils.GetRandomImageUrl()
		return nil
	}, func() error {
		backGroundImage = utils.GetRandomImageUrl()
		return nil
	}, func() error {
		signature = utils.GetRandomYiYan()
		return nil
	})

	// 创建 user 对象，用于写入数据库
	u := model.User{
		Id:              m.Id,
		Username:        m.Username,
		Password:        m.Password,
		Avatar:          avatar,
		BackgroundImage: backGroundImage,
		Signature:       signature,
	}
	fmt.Println(u)
	_, err = s.UserModel.Insert(context.Background(), &u)
	if err != nil {
		logx.Errorf(" user服务写入信息时候错误,err: ", err)
		return err
	}
	return nil

}
