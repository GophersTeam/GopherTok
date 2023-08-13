package model

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

//	func NewFavorModel(c redis.Options) FavorModel {
//		return &customModel{
//			defaultModel: NewDefaultModel(c),
//		}
//	}
func NewFavorModel(c *redis.ClusterClient) FavorModel {
	return &customModel{
		defaultModel: NewDefaultModel(c),
	}
}

//func NewDefaultModel(c redis.Options) defaultModel {
//	client := redis.NewClient(&c)
//	return defaultModel{
//		*client,
//	}
//}

func NewDefaultModel(c *redis.ClusterClient) defaultModel {
	return defaultModel{
		c,
	}
}

type (
	customModel struct {
		defaultModel
	}

	defaultModel struct {
		//redis.Client
		*redis.ClusterClient
	}
	FavorModel interface {
		favorModel
	}
	favorModel interface {
		Insert(ctx context.Context, UserId int64, VideoId int64) error
		Delete(ctx context.Context, UserId int64, VideoId int64) error
		SearchByUid(ctx context.Context, UserId int64) ([]int64, error)
		NumOfFavor(ctx context.Context, VideoId int64) (int, error)
		FavorNumOfUser(ctx context.Context, UserId int64) (int, error)
		IsFavor(ctx context.Context, UserId int64, VideoId int64) (bool, error)
	}
)

// 使用哈希和集合
func (m *defaultModel) Insert(ctx context.Context, UserId int64, VideoId int64) error {
	var err error

	tx := m.TxPipeline()

	// 在事务中执行命令
	if err = tx.SAdd(ctx, strconv.Itoa(int(VideoId)), UserId).Err(); err != nil {
		return err
	}

	if err = tx.HSet(ctx, strconv.Itoa(int(UserId)), strconv.Itoa(int(VideoId)), VideoId).Err(); err != nil {
		return err
	}

	_, err = tx.Exec(ctx)
	return err
}

func (m *defaultModel) Delete(ctx context.Context, UserId int64, VideoId int64) error {
	var err error
	tx := m.TxPipeline()

	// 在事务中执行命令
	if err = tx.SRem(ctx, strconv.Itoa(int(VideoId)), UserId).Err(); err != nil {
		return err
	}

	if err = tx.HDel(ctx, strconv.Itoa(int(UserId)), strconv.Itoa(int(VideoId))).Err(); err != nil {
		return err
	}

	_, err = tx.Exec(ctx)
	return err
}

func (m *defaultModel) SearchByUid(ctx context.Context, UserId int64) ([]int64, error) {
	result, err := m.HVals(ctx, strconv.Itoa(int(UserId))).Result()
	if err != nil {
		return nil, err
	}
	intSlice := make([]int64, len(result))
	for i, str := range result {
		num, err := strconv.Atoi(str)
		if err != nil {
			logx.Error(err.Error())
			return nil, err
		}
		intSlice[i] = int64(num)
	}
	return intSlice, nil
}

func (m *defaultModel) NumOfFavor(ctx context.Context, VideoId int64) (int, error) {
	result, err := m.SCard(ctx, strconv.Itoa(int(VideoId))).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

func (m *defaultModel) IsFavor(ctx context.Context, UserId int64, VideoId int64) (bool, error) {
	get := m.HGet(ctx, strconv.Itoa(int(UserId)), strconv.Itoa(int(VideoId)))
	if get.Err() != nil {
		if get.Err() == redis.Nil {
			return false, nil
		}
		return false, get.Err()
	}
	return true, nil
}

func (m *defaultModel) FavorNumOfUser(ctx context.Context, UserId int64) (int, error) {
	result, err := m.HLen(ctx, strconv.Itoa(int(UserId))).Result()
	if err != nil {
		return 0, err
	}
	return int(result), err
}

//查询用户是否点赞该视频，用户的获赞数目，用户的总点赞数 ，
//用户的获赞数目可以通过 id 插叙所属视频， 再查视频的或赞数目
