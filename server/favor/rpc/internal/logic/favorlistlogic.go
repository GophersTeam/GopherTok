package logic

import (
	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/user/rpc/userclient"
	"GopherTok/server/video/rpc/videoclient"
	"context"

	"GopherTok/server/favor/rpc/types/favor"

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

func (l *FavorListLogic) FavorList(in *favor.FavorListReq) (*favor.FavorListResp, error) {
	// todo: add your logic here and delete this line
	uids, err := l.svcCtx.FavorModel.SearchByUid(l.ctx, in.Userid)
	if err != nil {
		return nil, err
	}
	videos := make([]*favor.Video, 0)

	for i, id := range uids {
		video, err := l.svcCtx.VideoRpc.FindVideo(l.ctx, &videoclient.FindVideoReq{
			Id: id,
		})
		if err != nil {
			videos[i] = nil
			continue
		}
		videos[i] = &favor.Video{
			Id:          video.Video.Id,
			Title:       video.Video.Title,
			PlayUrl:     video.Video.PlayUrl,
			CoverUrl:    video.Video.CoverUrl,
			CreateTime:  video.Video.CreateTime,
			UpdateTime:  video.Video.UpdateTime,
			VideoSha256: video.Video.VideoSha256,
		}

		info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoReq{
			Id:        video.Video.UserId,
			CurrentId: in.Userid,
		})
		if err != nil {
			videos[i].Author = nil
			continue
		}
		videos[i].Author = &favor.User{
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
		}

	}

	return &favor.FavorListResp{
		Videos: videos,
	}, nil
}
