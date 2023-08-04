package logic

import (
	"context"

	"GopherTok/server/comment/rpc/internal/svc"
	"GopherTok/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *pb.AddCommentRequest) (resp *pb.AddCommentResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.AddCommentResponse)
	resp.Comment = &pb.Comment{
		Id:         123,
		VideoId:    in.VideoId,
		Content:    in.Content,
		CreateDate: "2023-08-08 08:08:08",
	}

	return
}
