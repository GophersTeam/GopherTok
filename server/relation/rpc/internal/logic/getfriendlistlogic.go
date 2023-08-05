package logic

import (
	"context"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *pb.GetFriendListReq) (*pb.GetFriendListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFriendListResp{}, nil
}
