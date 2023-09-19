package logic

import (
	"context"
	"fmt"
	"sync"

	"GopherTok/server/comment/rpc/commentrpc"
	"GopherTok/server/favor/api/internal/svc"
	"GopherTok/server/favor/api/internal/types"
	"GopherTok/server/favor/rpc/types/favor"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FavorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorListLogic {
	return &FavorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
//func (l *FavorListLogic) FavorList(req *types.FavorlistReq) (resp *types.FavorlistResp, err error) {
//	// todo: add your logic here and delete this line
//	userId := req.UserId
//
//	fmt.Println("user id =", userId)
//
//	list, err := l.svcCtx.FavorRpc.FavorList(l.ctx, &favor.FavorListReq{
//		Userid: userId,
//	})
//	if err != nil {
//		return nil, errors.Wrapf(err, "req: %+v", req)
//	}
//
//	videos := make([]types.Video, 0, len(list.Videos))
//
//	for _, video := range list.Videos {
//		videoItem := types.Video{
//			ID: video.Id,
//			Author: types.Author{
//				ID:              video.Author.Id,
//				Name:            video.Author.Name,
//				FollowCount:     video.Author.FollowCount,
//				FollowerCount:   video.Author.FollowerCount,
//				IsFollow:        video.Author.IsFollow,
//				Avatar:          video.Author.Avatar,
//				BackgroundImage: video.Author.BackgroundImage,
//				Signature:       video.Author.Signature,
//				TotalFavorited:  video.Author.TotalFavorited,
//				WorkCount:       video.Author.WorkCount,
//				FavoriteCount:   video.Author.FavoriteCount,
//			},
//			PlayURL:  video.PlayUrl,
//			CoverURL: video.CoverUrl,
//			Title:    video.Title,
//		}
//
//		num, err := l.svcCtx.FavorRpc.FavorNum(l.ctx, &favor.FavorNumReq{
//			VideoId: req.UserId,
//		})
//		if err != nil {
//			// 错误处理
//			// ...
//		} else {
//			videoItem.FavoriteCount = num.Num
//		}
//
//		count, err := l.svcCtx.CommenRpc.GetCommentCount(l.ctx, &commentrpc.GetCommentCountRequest{
//			VideoId: video.Id,
//		})
//		if err != nil {
//			// 错误处理
//			// ...
//		} else {
//			videoItem.CommentCount = count.Count
//		}
//
//		isFavor, err := l.svcCtx.FavorRpc.IsFavor(l.ctx, &favor.IsFavorReq{
//			UserId:  req.UserId,
//			VideoId: video.Id,
//		})
//		if err != nil {
//			// 错误处理
//			// ...
//		} else {
//			videoItem.IsFavorite = isFavor.IsFavor
//		}
//		videos = append(videos, videoItem)
//	}
//
//	return &types.FavorlistResp{
//		BaseResponse: types.BaseResponse{
//			Code:    0,
//			Message: "success",
//		},
//		Videos: videos,
//	}, nil
//}

func (l *FavorListLogic) FavorList(req *types.FavorlistReq) (resp *types.FavorlistResp, err error) {
	userId := req.UserId

	fmt.Println("user id =", userId)

	list, err := l.svcCtx.FavorRpc.FavorList(l.ctx, &favor.FavorListReq{
		Userid: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	videos := make([]types.Video, 0, len(list.Videos))
	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for _, video := range list.Videos {
		waitGroup.Add(1)
		go func(video *favor.Video) {
			defer waitGroup.Done()

			videoItem := types.Video{
				ID: video.Id,
				Author: types.Author{
					ID:              video.Author.Id,
					Name:            video.Author.Name,
					FollowCount:     video.Author.FollowCount,
					FollowerCount:   video.Author.FollowerCount,
					IsFollow:        video.Author.IsFollow,
					Avatar:          video.Author.Avatar,
					BackgroundImage: video.Author.BackgroundImage,
					Signature:       video.Author.Signature,
					TotalFavorited:  video.Author.TotalFavorited,
					WorkCount:       video.Author.WorkCount,
					FavoriteCount:   video.Author.FavoriteCount,
				},
				PlayURL:  video.PlayUrl,
				CoverURL: video.CoverUrl,
				Title:    video.Title,
			}

			num, err := l.svcCtx.FavorRpc.FavorNum(l.ctx, &favor.FavorNumReq{
				VideoId: video.Id,
			})
			if err != nil {
				videoItem.FavoriteCount = 0
			} else {
				videoItem.FavoriteCount = num.Num
			}

			count, err := l.svcCtx.CommenRpc.GetCommentCount(l.ctx, &commentrpc.GetCommentCountRequest{
				VideoId: video.Id,
			})
			if err != nil {
				videoItem.CommentCount = 0
			} else {
				videoItem.CommentCount = count.Count
			}

			isFavor, err := l.svcCtx.FavorRpc.IsFavor(l.ctx, &favor.IsFavorReq{
				UserId:  req.UserId,
				VideoId: video.Id,
			})
			if err != nil {
				videoItem.IsFavorite = false
			} else {
				videoItem.IsFavorite = isFavor.IsFavor
			}

			mutex.Lock()
			videos = append(videos, videoItem)
			mutex.Unlock()
		}(video)
	}

	waitGroup.Wait()

	return &types.FavorlistResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success",
		},
		Videos: videos,
	}, nil
}
