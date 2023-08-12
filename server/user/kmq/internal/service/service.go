package service

import (
	"GopherTok/common/init_db"
	"GopherTok/common/utils"
	"GopherTok/server/user/kmq/internal/config"
	"GopherTok/server/user/model"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type Service struct {
	c       config.Config // 配置信息
	MysqlDb *gorm.DB      // MySQL 数据库连接对象
	Rdb     *redis.ClusterClient
	Log     logx.LogConf
}

// NewService 创建一个新的 Service 实例
func NewService(c config.Config) *Service {
	// 初始化 MySQL 数据库连接
	mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)

	// 创建 video 表
	mysqlDb.AutoMigrate(&model.User{})

	// 创建 Service 实例
	s := &Service{
		c:       c,
		MysqlDb: mysqlDb,
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
	// 创建 user 对象，用于写入数据库
	v := model.User{
		ID:              m.ID,
		Username:        m.Username,
		Password:        m.Password,
		Avatar:          utils.GetRandomImageUrl(),
		BackgroundImage: utils.GetRandomImageUrl(),
		Signature:       utils.GetRandomYiYan(),
	}
	fmt.Println(v)

	// 写入 user 表
	if err := s.MysqlDb.Create(&v).Error; err != nil {
		logx.Error(err)
		return err
	}
	return nil

}
