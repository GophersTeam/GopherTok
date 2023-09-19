package logic

import (
	"GopherTok/common/errorx"
	"context"
	"github.com/pkg/errors"

	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVideoLogic {
	return &FindVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindVideoLogic) FindVideo(in *video.FindVideoReq) (*video.FindVideoResp, error) {
	// todo: add your logic here and delete this line
	v, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql查询出错，err:"+err.Error()), "mysql查询出错，err:%v", err)
	}
	return &video.FindVideoResp{
		Video: &video.VideoList{
			Id:          v.Id,
			UserId:      v.UserId,
			Title:       v.Title,
			PlayUrl:     v.PlayUrl,
			CoverUrl:    v.CoverUrl,
			CreateTime:  v.CreateTime.Unix(),
			UpdateTime:  v.UpdateTime.Unix(),
			VideoSha256: v.VideoSha256,
		},
	}, nil
}
