package model

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

func NewFavorModel(c redis.Options) FavorModel {
	return &customModel{
		defaultModel: NewDefaultModel(c),
	}

}

func NewDefaultModel(c redis.Options) defaultModel {
	client := redis.NewClient(&c)
	return defaultModel{
		*client,
	}
}

type (
	customModel struct {
		defaultModel
	}

	defaultModel struct {
		redis.Client
	}
	FavorModel interface {
		favorModel
	}
	favorModel interface {
		Insert(ctx context.Context, UserId int64, VideoId int64) error
		Delete(ctx context.Context, UserId int64, VideoId int64) error
		SearchByUid(ctx context.Context, UserId int64) ([]int64, error)
		NumOfFavor(ctx context.Context, VideoId int64) (int, error)
	}
)

// 使用哈希和集合
func (m *defaultModel) Insert(ctx context.Context, UserId int64, VideoId int64) error {
	var err error
	tx := m.Client.TxPipeline()

	// 在事务中执行命令
	if err = tx.SAdd(ctx, strconv.FormatInt(VideoId, 10), UserId).Err(); err != nil {
		return err
	}

	if err = tx.HSet(ctx, strconv.FormatInt(UserId, 10), VideoId, VideoId).Err(); err != nil {
		return err
	}

	_, err = tx.Exec(ctx)
	return err
}

func (m *defaultModel) Delete(ctx context.Context, UserId int64, VideoId int64) error {
	var err error
	tx := m.Client.TxPipeline()

	// 在事务中执行命令
	if err = tx.SRem(ctx, strconv.FormatInt(VideoId, 10), UserId).Err(); err != nil {
		return err
	}

	if err = tx.HDel(ctx, strconv.FormatInt(UserId, 10), strconv.Itoa(int(VideoId))).Err(); err != nil {
		return err
	}

	_, err = tx.Exec(ctx)
	return err
}

func (m *defaultModel) SearchByUid(ctx context.Context, UserId int64) ([]int64, error) {
	result, err := m.Client.HVals(ctx, strconv.FormatInt(UserId, 10)).Result()
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
	result, err := m.Client.SCard(ctx, strconv.FormatInt(VideoId, 10)).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}
