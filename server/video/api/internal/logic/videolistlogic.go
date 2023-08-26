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
	uid, ok := l.ctx.Value(consts.UserId).(int64)
	if !ok {
		logc.Info(l.ctx, "匿名用户")
		uid = 0
	}

	UserVideoListCnt, err := l.svcCtx.VideoRpc.VideoList(l.ctx, &video.VideoListReq{
		LatestTime: req.LatestTime,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	list := UserVideoListCnt.VideoList

	var (
		videoList    []*types.VideoInfo
		nextTime     int64 = math.MaxInt64
		videoResults       = make(chan struct {
			Index int
			Info  *types.VideoInfo
		}, len(list))
		errorChannel = make(chan error, len(list))
	)

	for i := 0; i < len(list); i++ {
		go func(i int) {
			// Perform RPC calls concurrently
			userinfo, uErr := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
				Id:        list[i].UserId,
				CurrentId: uid,
			})
			if uErr != nil {
				errorChannel <- uErr
				return
			}

			commentCount, cErr := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &pb.GetCommentCountRequest{
				VideoId: list[i].Id,
			})
			if cErr != nil {
				errorChannel <- cErr
				return
			}

			favoriteCount, fErr := l.svcCtx.FavorRpc.FavorNum(l.ctx, &favorrpc.FavorNumReq{
				VideoId: list[i].Id,
			})
			if fErr != nil {
				errorChannel <- fErr
				return
			}

			isFavorite := false
			if uid != 0 {
				isFavoriteCnt, favErr := l.svcCtx.FavorRpc.IsFavor(l.ctx, &favorrpc.IsFavorReq{
					UserId:  userinfo.Id,
					VideoId: list[i].Id,
				})
				if favErr != nil {
					errorChannel <- favErr
					return
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

			videoResults <- struct {
				Index int
				Info  *types.VideoInfo
			}{
				Index: i,
				Info:  videoItem,
			}
		}(i)
	}

	for i := 0; i < len(list); i++ {
		select {
		case result := <-videoResults:
			if result.Info != nil {
				videoList[result.Index] = result.Info
			}
		case err := <-errorChannel:
			return nil, err
		}
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
