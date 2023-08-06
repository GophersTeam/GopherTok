package logic

import (
	"context"

	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisFavorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDisFavorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisFavorLogic {
	return &DisFavorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DisFavorLogic) DisFavor(in *favor.DisFavorReq) (*favor.DisFavorResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.FavorModel.Delete(l.ctx, in.Userid, in.Videoid)
	if err != nil {
		return nil, err
	}

	return &favor.DisFavorResp{}, nil
}
