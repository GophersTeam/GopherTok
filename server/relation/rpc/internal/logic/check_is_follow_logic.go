package logic

import (
	"context"
	"fmt"

	"GopherTok/common/errorx"
	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/pkg/errors"

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
	boolCmd := l.svcCtx.Rdb.SIsMember(l.ctx, fmt.Sprintf("cache:gopherTok:follow:id:%d", in.ToUserId), in.UserId)
	if boolCmd.Err() != nil {
		return &pb.CheckIsFollowResp{
				StatusCode: -1,
				StatusMsg:  boolCmd.Err().Error(),
				IsFollow:   false,
			},
			errors.Wrapf(errorx.NewDefaultError("redis sismember err:"+boolCmd.Err().Error()), "redis sismember err ：%v", boolCmd.Err())
	}

	return &pb.CheckIsFollowResp{
		StatusCode: 0,
		StatusMsg:  "check isFollow successfully",
		IsFollow:   boolCmd.Val(),
	}, nil
}
