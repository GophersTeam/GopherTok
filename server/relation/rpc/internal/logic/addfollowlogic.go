package logic

import (
	"context"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowLogic {
	return &AddFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFollowLogic) AddFollow(in *pb.AddFollowReq) (*pb.AddFollowResp, error) {
	sadd, err := l.svcCtx.Rdb.Sadd(in.UserId, in.ToUserId)
	if err != nil {
		return nil, err
	}

	return &pb.AddFollowResp{}, nil
}
