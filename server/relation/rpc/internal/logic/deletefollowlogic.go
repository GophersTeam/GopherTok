package logic

import (
	"context"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFollowLogic {
	return &DeleteFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFollowLogic) DeleteFollow(in *pb.DeleteFollowReq) (*pb.DeleteFollowResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteFollowResp{}, nil
}
