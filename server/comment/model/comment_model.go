package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindByVideoId(ctx context.Context, videoId int64) (comments []*Comment, err error)
		GetCountByVideoId(ctx context.Context, videoId int64) (count int64, err error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

func (m *customCommentModel) GetCountByVideoId(ctx context.Context, videoId int64) (count int64, err error) {
	count, err = m.conn.CountDocuments(ctx, bson.M{"video_id": videoId})
	if err != nil {
		return 0, err
	}
	return
}

func (m *customCommentModel) FindByVideoId(ctx context.Context, videoId int64) (comments []*Comment, err error) {
	comments = []*Comment{}
	err = m.conn.Find(ctx, &comments, bson.M{"video_id": videoId}, &options.FindOptions{Sort: bson.M{"create_date": -1}})
	return
}

// NewCommentModel returns a model for the mongo.
func NewCommentModel(url, db, collection string) CommentModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customCommentModel{
		defaultCommentModel: newDefaultCommentModel(conn),
	}
}
