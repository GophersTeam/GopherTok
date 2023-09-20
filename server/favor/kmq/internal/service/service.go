package service

import (
	"GopherTok/common/init_db"
	"GopherTok/server/favor/kmq/internal/config"
	"GopherTok/server/favor/model"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"reflect"
)

type Service struct {
	Config     config.Config
	FavorModel model.FavorModel  // 用于redis插入
	MysqlDb    *gorm.DB      // MySQL 数据库连接对象
}

func structToStringSlice(s interface{}) []string {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		panic("s is not a struct")
	}

	var result []string
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String {
			result = append(result, field.String())
		}
	}

	return result
}


func NewService(c config.Config) *Service {
	//创建mysql表
	MysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	MysqlDb.AutoMigrate(&model.Favor{})
	slice := structToStringSlice(c.RedisCluster)
	redis := init_db.InitRedis(slice)
	return &Service{
		Config:     c,
		FavorModel: model.NewFavorModel(redis),
		MysqlDb :   MysqlDb,
	}
}


func (s *Service) Consume(_ string, value string) error {
	logx.Info("成功消费消息")
	var favorInfo model.FavorInfo
	err := jsonx.UnmarshalFromString(value, favorInfo)
	if err != nil {
		logx.Errorf("MessageAction jsonx.UnmarshalFromString error: %s", err.Error())
		return err
	}
	info := favorInfo.Info
	switch favorInfo.Method{
	case "favor" :
		if err := s.MysqlDb.Create(&info).Error; err != nil{
			logx.Error(err)
		    return err
	    }
		//插入redis
		err := s.FavorModel.Insert(context.Background(),info.UserId, info.VideoId)
		if err != nil {
			logx.Errorf("MessageAction FavorModel.Insert error: %s", err.Error())
			return err
		}

	case "disfavor" :
		//mysql
		if err := s.MysqlDb.Delete(&info).Error; err != nil{
			logx.Error(err)
			return err
		}
		//redis
		err := s.FavorModel.Delete(context.Background(),info.UserId, info.VideoId)
		if err != nil {
			logx.Errorf("MessageAction FavorModel.Delete error: %s", err.Error())
			return err
		}
	}
   return nil
}
