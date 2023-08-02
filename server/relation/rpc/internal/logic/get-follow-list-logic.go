package logic

import (
	"context"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *pb.GetFollowListReq) (*pb.GetFollowListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFollowListResp{}, nil
}
