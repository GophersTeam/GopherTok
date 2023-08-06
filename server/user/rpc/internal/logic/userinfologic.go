package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/user/model"
	"GopherTok/server/video/rpc/types/video"
	"context"
	"github.com/pkg/errors"
	"strconv"

	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	u := model.User{}
	err := l.svcCtx.MysqlDb.Where("id = ?", in.Id).First(&u).Error
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "mysql查询错误 err：%v", err)
	}
	followCount, err := l.svcCtx.RelationRpc.GetFollowCount(l.ctx, &pb.GetFollowCountReq{
		Userid: strconv.FormatInt(in.Id, 10),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}
	followerCount, err := l.svcCtx.RelationRpc.GetFollowerCount(l.ctx, &pb.GetFollowerCountReq{
		Userid: strconv.FormatInt(in.Id, 10),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}
	userVideoList, err := l.svcCtx.VideoRpc.UserVideoList(l.ctx, &video.UserVideoListReq{
		UserId: in.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}

	return &user.UserInfoResp{
		Id:              u.ID,
		Name:            u.Username,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		FollowCount:     followCount.Count,
		FollowerCount:   followerCount.Count,
		WorkCount:       int64(len(userVideoList.VideoList)),
	}, nil
}
