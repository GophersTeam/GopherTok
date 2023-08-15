package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/favor/rpc/favorrpc"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/relation/rpc/relationrpc"
	"GopherTok/server/user/model"
	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"
	"GopherTok/server/video/rpc/types/video"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
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

	u := model.User{}
	err := l.svcCtx.MysqlDb.Where("id = ?", in.Id).First(&u).Error
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "mysql查询错误 err：%v", err)
	}
	var followCountResp, followerCountResp, userVideoListResp, isFollowResp, totalFavoritedResp, favoriteCountResp interface{}
	errChan := make(chan error, 6) // Create a channel to collect potential errors

	// Concurrently execute the RPC calls using goroutines
	go func() {
		followCountResp, err = l.svcCtx.RelationRpc.GetFollowCount(l.ctx, &pb.GetFollowCountReq{Userid: in.Id})
		errChan <- err
	}()
	go func() {
		followerCountResp, err = l.svcCtx.RelationRpc.GetFollowerCount(l.ctx, &pb.GetFollowerCountReq{Userid: in.Id})
		errChan <- err
	}()
	go func() {
		userVideoListResp, err = l.svcCtx.VideoRpc.UserVideoList(l.ctx, &video.UserVideoListReq{UserId: in.Id})
		errChan <- err
	}()
	go func() {
		isFollowResp, err = l.svcCtx.RelationRpc.CheckIsFollow(l.ctx, &relationrpc.CheckIsFollowReq{UserId: in.CurrentId, ToUserId: in.Id})
		errChan <- err
	}()
	go func() {
		totalFavoritedResp, err = l.svcCtx.FavorRpc.FavoredNumOfUser(l.ctx, &favorrpc.FavoredNumOfUserReq{UserId: in.Id})
		errChan <- err
	}()
	go func() {
		favoriteCountResp, err = l.svcCtx.FavorRpc.FavorNumOfUser(l.ctx, &favorrpc.FavorNumOfUserReq{UserId: in.Id})
		errChan <- err
	}()

	// Wait for all RPC calls to complete
	for i := 0; i < 6; i++ {
		err := <-errChan // Retrieve errors from the channel
		if err != nil {
			// Handle the error, log, and return if needed
			logc.Error(l.ctx, err, "RPC call error")
			return nil, errors.Wrapf(err, "req: %+v", in)

		}
	}

	return &user.UserInfoResp{
		Id:              u.ID,
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
