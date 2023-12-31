package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"GopherTok/common/consts"
	"GopherTok/common/init_db"
	"GopherTok/common/utils"
	"GopherTok/server/video/kmq/internal/config"
	"GopherTok/server/video/model"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

const (
	chanCount   = 10   // 通道数量
	bufferCount = 1024 // 每个通道的缓冲区大小
)

type Service struct {
	c             config.Config // 配置信息
	MysqlDb       *gorm.DB      // MySQL 数据库连接对象
	Rdb           *redis.ClusterClient
	Log           logx.LogConf
	waiter        sync.WaitGroup      // 用于等待所有消费者 goroutine 完成的等待组
	msgsChan      []chan *model.Video // 消息通道切片，每个元素是一个通道，用于存放消息
	SensitiveTrie *utils.SensitiveTrie
	SqlConn       sqlx.SqlConn
	VideoModel    model.VideoModel
}

// NewService 创建一个新的 Service 实例
func NewService(c config.Config) *Service {
	mysqlConn := sqlx.NewMysql(c.MysqlCluster.DataSource)

	// 初始化redis
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)

	// 创建 Service 实例
	s := &Service{
		c:             c,
		msgsChan:      make([]chan *model.Video, chanCount),
		SqlConn:       mysqlConn,
		VideoModel:    model.NewVideoModel(mysqlConn, c.CacheRedis),
		Rdb:           redisDb,
		SensitiveTrie: utils.NewSensitiveTrie(),
	}

	// 创建 chanCount 个消费者 goroutine
	for i := 0; i < chanCount; i++ {
		ch := make(chan *model.Video, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}

	return s
}

// consume 是消费者 goroutine 的函数，负责处理从通道中接收的消息
func (s *Service) consume(ch chan *model.Video) {
	defer s.waiter.Done()

	for {
		message, ok := <-ch
		if !ok {
			log.Fatal("接受消息失败")
		}
		m := *message
		fmt.Printf("消费消息: %+v\n", m)

		// 创建 video 对象，用于写入数据库
		v := model.Video{
			Id:          m.Id,
			UserId:      m.UserId,
			Title:       m.Title,
			PlayUrl:     m.PlayUrl,
			CoverUrl:    m.CoverUrl,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
			VideoSha256: m.VideoSha256,
		}
		// 敏感词过滤
		s.SensitiveTrie.AddWords([]string{"傻逼", "死", "你妈", "滚"})
		v.Title = s.SensitiveTrie.Filter(v.Title)
		fmt.Println(v)
		// 并发写入redis,mysql
		err := mr.Finish(func() error {
			err := s.Rdb.Set(context.Background(), consts.VideoPrefix+m.VideoSha256, m.PlayUrl, 0)
			return err.Err()
		}, func() error {
			err := s.Rdb.Set(context.Background(), consts.CoverPrefix+m.VideoSha256, m.CoverUrl, 0)
			return err.Err()
		}, func() error {
			_, err := s.VideoModel.Insert(context.Background(), &v)
			if err != nil {
				logx.Error(err)
			}
			return err
		}, func() error {
			err := s.Rdb.ZAdd(context.Background(), consts.AllVideoIdPrefix, redis.Z{
				Score:  float64(v.CreateTime.Unix()),
				Member: v.Id,
			})
			return err.Err()
		}, func() error {
			err := s.Rdb.SAdd(context.Background(), fmt.Sprintf("%s%v", consts.UserVideoIdsPrefix, v.UserId), v.Id).Err()
			return err
		})

		if err != nil {
			logx.Error("video mq并发写入mysql,redis错误，err:", err)
		}

	}
}

// Consume 是消费者的方法，用于处理消息
func (s *Service) Consume(_, value string) error {
	logx.Infof("消费消息: %s\n", value)

	// 将 JSON 数据解析为 []*model.NewUserFile 对象
	var data []*model.Video
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	// 将解析后的消息根据 UserId 分发到不同的通道
	for _, d := range data {
		s.msgsChan[d.Id%chanCount] <- d
	}

	return nil
}
