package logic

import (
	"context"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendCountLogic {
	return &GetFriendCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendCountLogic) GetFriendCount(in *pb.GetFriendCountReq) (*pb.GetFriendCountResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFriendCountResp{}, nil
}
