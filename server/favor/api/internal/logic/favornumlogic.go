package logic

import (
	"GopherTok/server/favor/rpc/types/favor"
	"context"
	"github.com/pkg/errors"

	"GopherTok/server/favor/api/internal/svc"
	"GopherTok/server/favor/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorNumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorNumLogic {
	return &FavorNumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavorNumLogic) FavorNum(req *types.FavorNumReq) (resp *types.FavorNumResp, err error) {
	// todo: add your logic here and delete this line

	num, err := l.svcCtx.FavorRpc.FavorNum(l.ctx, &favor.FavorNumReq{
		Vedioid: req.Video_id,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.FavorNumResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success",
		},
		FavorNum: num.Num,
	}, nil
}
