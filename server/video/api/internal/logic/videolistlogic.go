package logic

import (
	"GopherTok/common/consts"
	"GopherTok/server/comment/rpc/pb"
	"GopherTok/server/favor/rpc/favorrpc"
	"GopherTok/server/user/rpc/types/user"
	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"
	"GopherTok/server/video/rpc/types/video"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
)

type VideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoListLogic) VideoList(req *types.VideoListReq) (resp *types.VideoListResp, err error) {
	// todo: add your logic here and delete this line
	uid, ok := l.ctx.Value(consts.UserId).(int64)
	if !ok {
		logc.Info(l.ctx, "匿名用户")
		uid = 0
		//return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", uid)
	}

	UserVideoListCnt, err := l.svcCtx.VideoRpc.VideoList(l.ctx, &video.VideoListReq{
		LatestTime: req.LatestTime,
	})
	list := UserVideoListCnt.VideoList
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	videoList := make([]*types.VideoInfo, 0) // Assuming VideoList is a struct that matches your needs
	nextTime := int64(math.MaxInt64)
	for i := 0; i < len(list); i++ {
		// 查看视频的作者信息
		userinfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
			Id:        list[i].UserId,
			CurrentId: uid,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}

		commentCount, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &pb.GetCommentCountRequest{
			VideoId: list[i].Id,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		favoriteCount, err := l.svcCtx.FavorRpc.FavorNum(l.ctx, &favorrpc.FavorNumReq{
			VideoId: list[i].Id,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		// 未登陆的用户,直接设置未点赞该视频,并且未关注该用户
		isFavorite := false
		if uid == 0 {
			userinfo.IsFollow = false
		} else {
			isFavoriteCnt, err := l.svcCtx.FavorRpc.IsFavor(l.ctx, &favorrpc.IsFavorReq{
				UserId:  userinfo.Id,
				VideoId: list[i].Id,
			})
			if err != nil {
				return nil, errors.Wrapf(err, "req: %+v", req)
			}
			isFavorite = isFavoriteCnt.IsFavor
		}
		if list[i].CreateTime < nextTime {
			nextTime = list[i].CreateTime
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
			FavoriteCount: favoriteCount.Num,
			CommentCount:  commentCount.Count,
			IsFavorite:    isFavorite,
		}
		videoList = append(videoList, videoItem)
	}

	return &types.VideoListResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success!",
		},
		NextTime:  nextTime * 1000,
		VideoList: types.VideoList{List: videoList},
	}, nil
}
