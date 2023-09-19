package logic

import (
	"GopherTok/server/favor/rpc/favorrpc"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/relation/rpc/relationrpc"
	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"
	"GopherTok/server/video/rpc/types/video"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/mr"
	"strconv"

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

	//u := model.User{}
	//err := l.svcCtx.MysqlDb.Where("id = ?", in.Id).First(&u).Error
	//if err != nil {
	//	return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "mysql查询错误 err：%v", err)
	//}
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	// 并发调用rpc
	var followCountResp, followerCountResp, userVideoListResp, isFollowResp, totalFavoritedResp, favoriteCountResp interface{}
	err = mr.Finish(func() error {
		followCountResp, err = l.svcCtx.RelationRpc.GetFollowCount(l.ctx, &pb.GetFollowCountReq{Userid: in.Id})
		return err
	}, func() error {
		followerCountResp, err = l.svcCtx.RelationRpc.GetFollowerCount(l.ctx, &pb.GetFollowerCountReq{Userid: in.Id})
		return err

	}, func() error {
		userVideoListResp, err = l.svcCtx.VideoRpc.UserVideoList(l.ctx, &video.UserVideoListReq{UserId: in.Id})
		return err

	}, func() error {
		isFollowResp, err = l.svcCtx.RelationRpc.CheckIsFollow(l.ctx, &relationrpc.CheckIsFollowReq{UserId: in.CurrentId, ToUserId: in.Id})
		return err

	}, func() error {
		totalFavoritedResp, err = l.svcCtx.FavorRpc.FavoredNumOfUser(l.ctx, &favorrpc.FavoredNumOfUserReq{UserId: in.Id})
		return err

	}, func() error {
		favoriteCountResp, err = l.svcCtx.FavorRpc.FavorNumOfUser(l.ctx, &favorrpc.FavorNumOfUserReq{UserId: in.Id})
		return err

	})

	if err != nil {
		// Handle the error, log, and return if needed
		logc.Error(l.ctx, err, "RPC call error")
		return nil, errors.Wrapf(err, "req: %+v", in)

	}
	return &user.UserInfoResp{
		Id:              u.Id,
		Name:            u.Username,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		IsFollow:        isFollowResp.(*relationrpc.CheckIsFollowResp).IsFollow,
		FollowCount:     followCountResp.(*pb.GetFollowCountResp).Count,
		FollowerCount:   followerCountResp.(*pb.GetFollowerCountResp).Count,
		TotalFavorited:  strconv.FormatInt(totalFavoritedResp.(*favorrpc.FavoredNumOfUserResp).FavoredNumOfUser, 10),
		WorkCount:       int64(len(userVideoListResp.(*video.UserVideoListResp).VideoList)),
		FavoriteCount:   favoriteCountResp.(*favorrpc.FavorNumOfUserResp).FavorNumOfUser,
	}, nil
}
