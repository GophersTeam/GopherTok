package logic

import (
	"GopherTok/common/errorx"
	"context"
	"github.com/pkg/errors"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsFollowLogic {
	return &CheckIsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsFollowLogic) CheckIsFollow(in *pb.CheckIsFollowReq) (*pb.CheckIsFollowResp, error) {
	isFollow, err := l.svcCtx.Rdb.Sismember(string(in.UserId), in.ToUserId)
	if err != nil {
		return &pb.CheckIsFollowResp{StatusCode: "-1",
				StatusMsg: err.Error(),
				IsFollow:  false},
			errors.Wrapf(errorx.NewDefaultError("redis sismember err:"+err.Error()), "redis sismember err ：%v", err)
	}

	return &pb.CheckIsFollowResp{StatusCode: "0",
		StatusMsg: "check isFollow successfully",
		IsFollow:  isFollow}, nil
}
