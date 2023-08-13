package logic

import (
	"GopherTok/server/comment/rpc/commentrpc"
	"GopherTok/server/favor/rpc/types/favor"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"GopherTok/server/favor/api/internal/svc"
	"GopherTok/server/favor/api/internal/types"

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

func (l *FavorListLogic) FavorList(req *types.FavorlistReq) (resp *types.FavorlistResp, err error) {
	// todo: add your logic here and delete this line
	userId := req.UserId

	fmt.Println("user id =", userId)

	videos := make([]types.Video, 0)

	list, err := l.svcCtx.FavorRpc.FavorList(l.ctx, &favor.FavorListReq{
		Userid: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	for i, video := range list.Videos {
		videos[i] = types.Video{
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
			//FavoriteCount: num.Num,
			//CommentCount:  0,
			//IsFavorite:    false,
			Title: video.Title,
		}

		num, err := l.svcCtx.FavorRpc.FavorNum(l.ctx, &favor.FavorNumReq{
			VideoId: req.UserId,
		})
		if err == nil {
			videos[i].FavoriteCount = num.Num
		}
		count, err := l.svcCtx.CommenRpc.GetCommentCount(l.ctx, &commentrpc.GetCommentCountRequest{
			VideoId: video.Id,
		})
		if err == nil {
			videos[i].CommentCount = count.Count
		}
		isFavor, err := l.svcCtx.FavorRpc.IsFavor(l.ctx, &favor.IsFavorReq{
			UserId:  req.UserId,
			VideoId: video.Id,
		})
		if err == nil {
			videos[i].IsFavorite = isFavor.IsFavor
		}
	}

	return &types.FavorlistResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success",
		},
		Videos: videos,
	}, nil
}
