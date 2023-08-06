package logic

import (
	"context"

	"GopherTok/server/favor/rpc/internal/svc"
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

	return &favor.FavorListResp{
		Videoids: uids,
	}, nil
}
