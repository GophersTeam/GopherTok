package logic

import (
	"context"

	"GopherTok/server/comment/rpc/internal/svc"
	"GopherTok/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCommentLogic {
	return &DelCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCommentLogic) DelComment(in *pb.DelCommentRequest) (resp *pb.DelCommentResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.DelCommentResponse)

	return
}
