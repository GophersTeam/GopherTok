package logic

import (
	"context"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowCountLogic {
	return &GetFollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowCountLogic) GetFollowCount(in *pb.GetFollowCountReq) (*pb.GetFollowCountResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFollowCountResp{}, nil
}
