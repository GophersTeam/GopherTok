package logic

import (
	"context"
	"sync"
	"time"

	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"
	"GopherTok/server/user/rpc/userclient"
	"GopherTok/server/video/rpc/videoclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorListLogic {
	return &FavorListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//func (l *FavorListLogic) FavorList(in *favor.FavorListReq) (*favor.FavorListResp, error) {
//	// todo: add your logic here and delete this line
//	uids, err := l.svcCtx.FavorModel.SearchByUid(l.ctx, in.Userid)
//	if err != nil {
//		return nil, err
//	}
//	videos := make([]*favor.Video, 0)
//
//	for _, id := range uids {
//		video, err := l.svcCtx.VideoRpc.FindVideo(l.ctx, &videoclient.FindVideoReq{
//			Id: id,
//		})
//		if err != nil {
//			//
//			continue
//		}
//		info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoReq{
//			Id:        video.Video.UserId,
//			CurrentId: in.Userid,
//		})
//		if err != nil {
//			//
//			continue
//		}
//
//		videos = append(videos, &favor.Video{
//			Author: &favor.User{
//				Id:              info.Id,
//				Name:            info.Name,
//				Avatar:          info.Avatar,
//				BackgroundImage: info.BackgroundImage,
//				Signature:       info.Signature,
//				IsFollow:        info.IsFollow,
//				FollowCount:     info.FollowerCount,
//				FollowerCount:   info.FollowCount,
//				TotalFavorited:  info.TotalFavorited,
//				WorkCount:       info.WorkCount,
//				FavoriteCount:   info.FavoriteCount,
//			},
//			Id:          video.Video.Id,
//			Title:       video.Video.Title,
//			PlayUrl:     video.Video.PlayUrl,
//			CoverUrl:    video.Video.CoverUrl,
//			CreateTime:  time.Unix(video.Video.CreateTime, 0).Format("2006-01-02 15:04:05"),
//			UpdateTime:  time.Unix(video.Video.UpdateTime, 0).Format("2006-01-02 15:04:05"),
//			VideoSha256: video.Video.VideoSha256,
//		})
//
//	}
//
//	return &favor.FavorListResp{
//		Videos: videos,
//	}, nil
//}

func (l *FavorListLogic) FavorList(in *favor.FavorListReq) (*favor.FavorListResp, error) {
	uids, err := l.svcCtx.FavorModel.SearchByUid(l.ctx, in.Userid)
	if err != nil {
		return nil, err
	}
	videos := make([]*favor.Video, 0)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, id := range uids {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			video, err := l.svcCtx.VideoRpc.FindVideo(l.ctx, &videoclient.FindVideoReq{
				Id: id,
			})
			if err != nil {
				return
			}
			info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoReq{
				Id:        video.Video.UserId,
				CurrentId: in.Userid,
			})
			if err != nil {
				return
			}

			mu.Lock()
			videos = append(videos, &favor.Video{
				Author: &favor.User{
					Id:              info.Id,
					Name:            info.Name,
					Avatar:          info.Avatar,
					BackgroundImage: info.BackgroundImage,
					Signature:       info.Signature,
					IsFollow:        info.IsFollow,
					FollowCount:     info.FollowerCount,
					FollowerCount:   info.FollowCount,
					TotalFavorited:  info.TotalFavorited,
					WorkCount:       info.WorkCount,
					FavoriteCount:   info.FavoriteCount,
				},
				Id:          video.Video.Id,
				Title:       video.Video.Title,
				PlayUrl:     video.Video.PlayUrl,
				CoverUrl:    video.Video.CoverUrl,
				CreateTime:  time.Unix(video.Video.CreateTime, 0).Format("2006-01-02 15:04:05"),
				UpdateTime:  time.Unix(video.Video.UpdateTime, 0).Format("2006-01-02 15:04:05"),
				VideoSha256: video.Video.VideoSha256,
			})
			mu.Unlock()
		}(id)
	}

	wg.Wait()

	return &favor.FavorListResp{
		Videos: videos,
	}, nil
}
