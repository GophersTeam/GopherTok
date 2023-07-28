package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Message, error)
		GetMessages(ctx context.Context, fromUserId int64, toUserId int64, preMsgTime int64) ([]*Message, error)
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

func (m *customMessageModel) FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Message, error) {
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", messageRows, m.table)
	var resp []*Message
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (pageNum-1)*pageSize, pageSize)

	return resp, err
}

func (m *customMessageModel) GetMessages(ctx context.Context, fromUserId int64, toUserId int64, preMsgTime int64) ([]*Message, error) {
	query := fmt.Sprintf(`select %s from %s where 
		((from_user_id = ? and to_user_id = ?) or (from_user_id = ? and to_user_id = ?)) 
        and create_time < ? order by create_time`, messageRows, m.table)
	var resp []*Message
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, fromUserId, toUserId, toUserId, fromUserId, preMsgTime)

	return resp, err
}

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn, c, opts...),
	}
}
