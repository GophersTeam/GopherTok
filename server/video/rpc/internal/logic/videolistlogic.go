package logic

import (
	"context"

	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoListLogic) VideoList(in *video.VideoListReq) (*video.VideoListResp, error) {
	// todo: add your logic here and delete this line

	return &video.VideoListResp{}, nil
}
