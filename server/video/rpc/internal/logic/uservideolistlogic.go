package logic

import (
	"context"

	"GopherTok/common/errorx"

	"github.com/pkg/errors"

	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVideoListLogic {
	return &UserVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserVideoListLogic) UserVideoList(in *video.UserVideoListReq) (*video.UserVideoListResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.VideoModel.FindVideosByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql find 错误"+err.Error()), "mysql find err:%v", err)
	}

	videoList := make([]*video.VideoList, 0)

	for i := 0; i < len(list); i++ {
		videoItem := &video.VideoList{
			Id:          list[i].Id,
			UserId:      list[i].UserId,
			Title:       list[i].Title,
			PlayUrl:     list[i].PlayUrl,
			CoverUrl:    list[i].CoverUrl,
			CreateTime:  list[i].CreateTime.Unix(),
			UpdateTime:  list[i].UpdateTime.Unix(),
			VideoSha256: list[i].VideoSha256,
		}
		videoList = append(videoList, videoItem)
	}
	return &video.UserVideoListResp{
		VideoList: videoList,
	}, nil
}
