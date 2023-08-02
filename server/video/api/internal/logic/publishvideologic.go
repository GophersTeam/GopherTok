package logic

import (
	"context"

	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishVideoReq) (resp *types.PublishVideoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
