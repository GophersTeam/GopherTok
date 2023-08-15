package logic

import (
	"GopherTok/common/errorx"
	"context"
	"github.com/pkg/errors"

	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFavorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFavorLogic {
	return &IsFavorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFavorLogic) IsFavor(in *favor.IsFavorReq) (*favor.IsFavorResp, error) {

	isFavor, err := l.svcCtx.FavorModel.IsFavor(l.ctx, in.UserId, in.VideoId)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "err:%v", err)

	}
	return &favor.IsFavorResp{
		IsFavor: isFavor,
	}, nil
}
