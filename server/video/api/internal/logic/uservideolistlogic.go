package logic

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/comment/rpc/pb"
	"GopherTok/server/user/rpc/types/user"
	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"
	"GopherTok/server/video/rpc/types/video"
	"context"
	"github.com/pkg/errors"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVideoListLogic {
	return &UserVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserVideoListLogic) UserVideoList(req *types.UserVideoListReq) (resp *types.UserVideoListResp, err error) {
	// todo: add your logic here and delete this line
	// 获取user id
	uid, ok := l.ctx.Value(consts.UserId).(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", uid)
	}
	userId, _ := strconv.ParseInt(req.UserId, 10, 64)
	UserVideoListCnt, err := l.svcCtx.VideoRpc.UserVideoList(l.ctx, &video.UserVideoListReq{
		UserId: userId,
	})
	list := UserVideoListCnt.VideoList
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	userinfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
		Id:        userId,
		CurrentId: uid,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	videoList := make([]*types.VideoInfo, 0) // Assuming VideoList is a struct that matches your needs

	for i := 0; i < len(list); i++ {
		commentCount, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &pb.GetCommentCountRequest{
			VideoId: list[i].Id,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		videoItem := &types.VideoInfo{
			ID: list[i].Id,
			Author: types.AuthorInfo{
				ID:              userinfo.Id,
				Name:            userinfo.Name,
				FollowCount:     userinfo.FollowCount,
				FollowerCount:   userinfo.FollowerCount,
				IsFollow:        userinfo.IsFollow,
				Avatar:          userinfo.Avatar,
				BackgroundImage: userinfo.BackgroundImage,
				Signature:       userinfo.Signature,
				TotalFavorited:  userinfo.TotalFavorited,
				WorkCount:       userinfo.WorkCount,
				FavoriteCount:   userinfo.FavoriteCount,
			},
			Title:         list[i].Title,
			PlayURL:       list[i].PlayUrl,
			CoverURL:      list[i].CoverUrl,
			FavoriteCount: 0,
			CommentCount:  commentCount.Count,
		}
		videoList = append(videoList, videoItem)
	}
	return &types.UserVideoListResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success!",
		},
		VideoList: types.VideoList{List: videoList},
	}, nil
}
