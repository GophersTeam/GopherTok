package logic

import (
	"context"

	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorLogic {
	return &FavorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavorLogic) Favor(in *favor.FavorReq) (*favor.FavorResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.FavorModel.Insert(l.ctx, in.Userid, in.Videoid)
	if err != nil {
		return nil, err
	}

	return &favor.FavorResp{}, nil
}
