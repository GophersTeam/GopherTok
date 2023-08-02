package logic

import (
	"context"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *pb.GetFollowerReq) (*pb.GetFollowerResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFollowerResp{}, nil
}
