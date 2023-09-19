package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"GopherTok/common/init_db"
	"GopherTok/server/relation/dao"
	"GopherTok/server/relation/kmq/internal/config"
	"GopherTok/server/relation/rpc/pb"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	chanCount   = 10   // 通道数量
	bufferCount = 1024 // 每个通道的缓冲区大小
)

type Service struct {
	c              config.Config // 配置信息
	MysqlDb        *gorm.DB      // MySQL 数据库连接对象
	Rdb            *redis.ClusterClient
	Log            logx.LogConf
	waiter         sync.WaitGroup         // 用于等待所有消费者 goroutine 完成的等待组
	msgsFollowChan []chan *dao.FollowData // 消息通道切片，每个元素是一个通道，用于存放消息
	msgsCountChan  []chan *dao.CountData
}

// NewService 创建一个新的 Service 实例
func NewService(c config.Config) *Service {
	// 初始化 MySQL 数据库连接
	mysqlDb := InitGorm(c.Mysql.DataSource)

	// 初始化redis
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	// 创建 Follow 表
	_ = mysqlDb.Table("follow_subject").AutoMigrate(&pb.FollowSubject{})

	// 创建 Service 实例
	s := &Service{
		c:              c,
		msgsFollowChan: make([]chan *dao.FollowData, chanCount/2),
		msgsCountChan:  make([]chan *dao.CountData, chanCount/2),
		MysqlDb:        mysqlDb,
		Rdb:            redisDb,
	}

	// 创建 chanCount 个消费者 goroutine
	for i := 0; i < chanCount/2; i++ {
		ch := make(chan *dao.FollowData, bufferCount)
		s.msgsFollowChan[i] = ch
		s.waiter.Add(1)
		go s.consumeMysql(ch)
	}

	for i := 0; i < chanCount/2; i++ {
		ch := make(chan *dao.CountData, bufferCount)
		s.msgsCountChan[i] = ch
		s.waiter.Add(1)
		go s.consumeRedis(ch)
	}
	return s
}

// consumeMysql 是消费者 goroutine 的函数，负责处理从通道中接收的消息
func (s *Service) consumeMysql(ch chan *dao.FollowData) {
	defer s.waiter.Done()

	for {
		message, ok := <-ch
		if !ok {
			log.Fatal("接受消息失败")
		}
		m := *message
		fmt.Printf("消费消息: %+v\n", m)

		// 创建 follow 对象，用于写入数据库

		if m.Method == "creat" {
			if err := s.MysqlDb.Table("follow_subject").Create(&pb.FollowSubject{
				UserId:     m.UserId,
				FollowerId: m.FollowerId,
				IsFollow:   m.IsFollow,
			}).Error; err != nil {
				logx.Error(err)
			}
		} else if m.Method == "delete" {
			if err := s.MysqlDb.Table("follow_subject").
				Where("user_id = ? AND follower_id = ?", m.UserId, m.FollowerId).Delete(&pb.FollowSubject{}).Error; err != nil {
				logx.Error(err)
			}
		}
	}
}

func (s *Service) consumeRedis(ch chan *dao.CountData) {
	defer s.waiter.Done()

	for {
		message, ok := <-ch
		if !ok {
			log.Fatal("接受消息失败")
		}
		m := *message
		fmt.Printf("消费消息: %+v\n", m)

		// 更新redis
		if err := s.Rdb.HSet(context.Background(), "followCount", m.FollowCountKey, m.FollowCount).Err(); err != nil {
			logx.Error(err)
		}
		if err := s.Rdb.HSet(context.Background(), "followerCount", m.FollowerCountKey, m.FollowerCount).Err(); err != nil {
			logx.Error(err)
		}
		if err := s.Rdb.HSet(context.Background(), "friendCount", m.FriendCountKey, m.FriendCount).Err(); err != nil {
			logx.Error(err)
		}

	}
}

// ConsumeRedis 是消费者的方法，用于处理消息
func (s *Service) ConsumeMysql(_, value string) error {
	logx.Infof("消费消息: %s\n", value)

	// 将 JSON 数据解析为 []*model.NewUserFile 对象
	var data *dao.FollowData
	if err := jsonx.UnmarshalFromString(value, &data); err != nil {
		return err
	}

	s.msgsFollowChan[data.FollowerId%(chanCount/2)] <- data

	return nil
}

func (s *Service) ConsumeRedis(_, value string) error {
	logx.Infof("消费消息: %s\n", value)

	// 将 JSON 数据解析为 []*model.NewUserFile 对象
	var data *dao.CountData
	if err := jsonx.UnmarshalFromString(value, &data); err != nil {
		return err
	}

	s.msgsCountChan[len(data.FollowerCount)%(chanCount/2)] <- data

	return nil
}

func InitGorm(MysqlDataSourece string) *gorm.DB {
	// 将日志写进kafka
	// logx.SetWriter(*LogxKafka())
	DB, err := gorm.Open(mysql.Open(MysqlDataSourece),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	db, _ := DB.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute)
	return DB
}
