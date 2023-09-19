package svc

import (
	"GopherTok/common/consts"
	"GopherTok/common/init_db"
	"GopherTok/server/favor/rpc/favorrpc"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/relation/rpc/relationrpc"
	"GopherTok/server/user/model"
	"GopherTok/server/user/rpc/internal/config"
	"GopherTok/server/video/rpc/videoclient"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type ServiceContext struct {
	Config         config.Config
	RelationRpc    pb.RelationRpcClient
	VideoRpc       videoclient.Video
	FavorRpc       favorrpc.FavorRpc
	Snowflake      *snowflake.Node
	Rdb            *redis.ClusterClient
	MysqlDb        *gorm.DB
	KqPusherClient *kq.Pusher
	SqlConn        sqlx.SqlConn
	UserModel      model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	//mysqlDb.AutoMigrate(&model.User{})
	snowflakeNode, _ := snowflake.NewNode(consts.UserMachineId)

	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.Cluster1, c.RedisCluster.Cluster2, c.RedisCluster.Cluster3, c.RedisCluster.Cluster4, c.RedisCluster.Cluster5, c.RedisCluster.Cluster6)
	redisDb := init_db.InitRedis(rc)
	mysqlConn := sqlx.NewMysql(c.MysqlCluster.DataSource)
	return &ServiceContext{
		Config:      c,
		SqlConn:     mysqlConn,
		UserModel:   model.NewUserModel(mysqlConn, c.CacheRedis),
		RelationRpc: relationrpc.NewRelationRpc(zrpc.MustNewClient(c.RelationRpcConf)),
		VideoRpc:    videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpcConf)),
		FavorRpc:    favorrpc.NewFavorRpc(zrpc.MustNewClient(c.FavorRpcConf)),
		Snowflake:   snowflakeNode,
		//MysqlDb:        mysqlDb,
		Rdb:            redisDb,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
func InitGorm(MysqlDataSourece string) *gorm.DB {
	// 将日志写进kafka
	//logx.SetWriter(*LogxKafka())
	DB, err := gorm.Open(mysql.Open(MysqlDataSourece),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	db, _ := DB.DB()
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(200)
	db.SetConnMaxLifetime(time.Minute)
	return DB
}
