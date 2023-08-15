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
	"sync"
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
	}

	UserVideoListCnt, err := l.svcCtx.VideoRpc.VideoList(l.ctx, &video.VideoListReq{
		LatestTime: req.LatestTime,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	list := UserVideoListCnt.VideoList

	// ...

	videoList := make([]*types.VideoInfo, 0)
	nextTime := int64(math.MaxInt64)
	errChan := make(chan error, len(list)) // Create a channel to collect potential errors

	var wg sync.WaitGroup
	for i := 0; i < len(list); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// 查看视频的作者信息
			userInfoChan := make(chan *user.UserInfoResp)
			errChanUserInfo := make(chan error)
			go func() {
				userinfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
					Id:        list[i].UserId,
					CurrentId: uid,
				})
				if err != nil {
					errChanUserInfo <- err
				} else {
					userInfoChan <- userinfo
				}
			}()

			// 获取评论数
			commentCountChan := make(chan *pb.GetCommentCountResponse)
			errChanCommentCount := make(chan error)
			go func() {
				commentCount, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &pb.GetCommentCountRequest{
					VideoId: list[i].Id,
				})
				if err != nil {
					errChanCommentCount <- err
				} else {
					commentCountChan <- commentCount
				}
			}()

			// Wait for goroutines to complete
			userinfo := <-userInfoChan
			errUserInfo := <-errChanUserInfo
			commentCountResp := <-commentCountChan
			errCommentCount := <-errChanCommentCount

			if errUserInfo != nil || errCommentCount != nil {
				errChan <- errors.Wrapf(errUserInfo, "req: %+v", req)
				return
			}

			// ... Other parts of your loop ...

			// 获取 favoriteCount 和 isFavorite
			favoriteCount, err := l.svcCtx.FavorRpc.FavorNum(l.ctx, &favorrpc.FavorNumReq{
				VideoId: list[i].Id,
			})
			if err != nil {
				errChan <- errors.Wrapf(err, "req: %+v", req)
				return
			}

			isFavorite := false
			if uid != 0 {
				isFavoriteCnt, err := l.svcCtx.FavorRpc.IsFavor(l.ctx, &favorrpc.IsFavorReq{
					UserId:  userinfo.Id,
					VideoId: list[i].Id,
				})
				if err != nil {
					errChan <- errors.Wrapf(err, "req: %+v", req)
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
				CommentCount:  commentCountResp.Count,
				IsFavorite:    isFavorite,
			}
			videoList = append(videoList, videoItem)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errChan)

	// Collect errors
	for err := range errChan {
		if err != nil {
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
