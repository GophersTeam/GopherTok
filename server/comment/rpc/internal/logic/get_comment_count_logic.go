package logic

import (
	"context"

	"GopherTok/server/comment/rpc/internal/svc"
	"GopherTok/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentCountLogic {
	return &GetCommentCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentCountLogic) GetCommentCount(in *pb.GetCommentCountRequest) (resp *pb.GetCommentCountResponse, err error) {
	count, err := l.svcCtx.CommentModel.GetCountByVideoId(l.ctx, in.VideoId)
	if err != nil {
		l.Errorf("Get comment count error: %v", err)
		return
	}

	resp = new(pb.GetCommentCountResponse)
	resp.Count = count

	return
}
