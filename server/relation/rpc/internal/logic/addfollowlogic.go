package logic

import (
	"GopherTok/common/errorx"
	"context"
	"github.com/pkg/errors"

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
	_, err := l.svcCtx.Rdb.Sadd(string(in.UserId), in.ToUserId)
	if err != nil {
		return &pb.AddFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("redis set err:"+err.Error()), "redis set err ：%v", err)

	}
	err = l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").Create(&pb.FollowSubject{
		UserId:     in.UserId,
		FollowerId: 0,
		IsFollow:   false,
	}).Error
	if err != nil {
		return &pb.AddFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("mysql add err:"+err.Error()), "mysql add err ：%v", err)

	}
	return &pb.AddFollowResp{StatusCode: "0",
		StatusMsg: "add follow successfully"}, nil
}
