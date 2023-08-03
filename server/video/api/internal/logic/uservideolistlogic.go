package logic

import (
	"context"

	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVideoListLogic {
	return &UserVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserVideoListLogic) UserVideoList(req *types.UserVideoListReq) (resp *types.UserVideoListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
