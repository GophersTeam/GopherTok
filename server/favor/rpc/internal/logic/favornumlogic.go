package logic

import (
	"context"

	"GopherTok/common/errorx"
	"GopherTok/server/favor/rpc/internal/svc"

	"github.com/pkg/errors"

	"GopherTok/server/favor/rpc/types/favor"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavorNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorNumLogic {
	return &FavorNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavorNumLogic) FavorNum(in *favor.FavorNumReq) (*favor.FavorNumResp, error) {
	// todo: add your logic here and delete this line
	num, err := l.svcCtx.FavorModel.NumOfFavor(l.ctx, in.VideoId)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "err:%v", err)
	}

	return &favor.FavorNumResp{
		Num: int64(num),
	}, nil
}
