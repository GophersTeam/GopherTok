package model

import (
	"context"
	"fmt"
	"strconv"

	"GopherTok/common/consts"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindVideosByUserId(ctx context.Context, UserId int64) ([]*Video, error)
		FindVideoIdsByUserId(ctx context.Context, rdb *redis.ClusterClient, UserId int64) ([]int64, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c, opts...),
	}
}

func (m *defaultVideoModel) FindVideosByUserId(ctx context.Context, UserId int64) ([]*Video, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? order by create_time desc", videoRows, m.table)
	var v []*Video
	err := m.QueryRowsNoCacheCtx(ctx, v, query, UserId, UserId)
	return v, err
}

func (m *defaultVideoModel) FindVideoIdsByUserId(ctx context.Context, rdb *redis.ClusterClient, UserId int64) ([]int64, error) {
	// 先从redis里面查
	vIds := make([]int64, 0)

	res, err := rdb.SMembers(context.Background(), fmt.Sprintf("%s%v", consts.UserVideoIdsPrefix, UserId)).Result()
	// 没有走mysql
	if err != nil {
		logc.Error(ctx, "redis查询用户全部视频的id错误：err", err)
		query := fmt.Sprintf("select id from %s where user_id = ?", m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &vIds, query, UserId, UserId)
		if err != nil {
			return nil, err
		}
		return vIds, nil
	}
	// 缓存查询成功，走redis
	for _, v := range res {
		num, _ := strconv.ParseInt(v, 10, 64)
		vIds = append(vIds, num)
	}
	return vIds, nil
}
