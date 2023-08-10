package service

import (
	"GopherTok/common/init_db"
	"GopherTok/server/relation/kmq/internal/config"
	"GopherTok/server/relation/rpc/pb"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"log"
	"sync"
)

const (
	chanCount   = 10   // 通道数量
	bufferCount = 1024 // 每个通道的缓冲区大小
)

type Service struct {
	c        config.Config // 配置信息
	MysqlDb  *gorm.DB      // MySQL 数据库连接对象
	Rdb      *redis.ClusterClient
	Log      logx.LogConf
	waiter   sync.WaitGroup           // 用于等待所有消费者 goroutine 完成的等待组
	msgsChan []chan *pb.FollowSubject // 消息通道切片，每个元素是一个通道，用于存放消息
}

// NewService 创建一个新的 Service 实例
func NewService(c config.Config) *Service {
	// 初始化 MySQL 数据库连接
	mysqlDb := init_db.InitGorm(c.Mysql.DataSource)

	// 初始化redis
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	// 创建 video 表
	_ = mysqlDb.AutoMigrate(&pb.FollowSubject{})

	// 创建 Service 实例
	s := &Service{
		c:        c,
		msgsChan: make([]chan *pb.FollowSubject, chanCount),
		MysqlDb:  mysqlDb,
		Rdb:      redisDb,
	}

	// 创建 chanCount 个消费者 goroutine
	for i := 0; i < chanCount; i++ {
		ch := make(chan *pb.FollowSubject, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}

	return s
}

// consume 是消费者 goroutine 的函数，负责处理从通道中接收的消息
func (s *Service) consume(ch chan *pb.FollowSubject) {
	defer s.waiter.Done()

	for {
		message, ok := <-ch
		if !ok {
			log.Fatal("接受消息失败")
		}
		m := *message
		fmt.Printf("消费消息: %+v\n", m)

		// 创建 follow 对象，用于写入数据库
		f := pb.FollowSubject{}
		_ = copier.Copy(f, m)
		fmt.Println(f)

		if err := s.MysqlDb.Table("follow_subject").Create(&f).Error; err != nil {
			logx.Error(err)
		}

	}
}

// Consume 是消费者的方法，用于处理消息
func (s *Service) Consume(_ string, value string) error {
	logx.Infof("消费消息: %s\n", value)

	// 将 JSON 数据解析为 []*model.NewUserFile 对象
	var data []*pb.FollowSubject
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	i := 0
	for _, d := range data {
		s.msgsChan[i%chanCount] <- d
		i++
	}

	return nil
}
