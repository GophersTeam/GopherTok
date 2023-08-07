package logic

import (
	"GopherTok/common/errorx"
	"context"
	"github.com/pkg/errors"

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

	_, err := l.svcCtx.Rdb.Srem(string(in.UserId), in.ToUserId)
	if err != nil {
		return &pb.DeleteFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("redis srem err:"+err.Error()), "redis srem err ：%v", err)

	}

	err = l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("user_id = ? AND follower_id = ?", in.ToUserId, in.UserId).
		Delete(&pb.FollowSubject{}).
		Error

	if err != nil {
		return &pb.DeleteFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("mysql delete err:"+err.Error()), "mysql delete err ：%v", err)

	}
	return &pb.DeleteFollowResp{
		StatusCode: "0",
		StatusMsg:  "delete follow successfully",
	}, nil
}
