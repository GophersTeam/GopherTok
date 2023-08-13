package logic

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/favor/api/internal/svc"
	"GopherTok/server/favor/api/internal/types"
	"GopherTok/server/favor/rpc/types/favor"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorLogic {
	return &FavorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavorLogic) Favor(req *types.FavorReq) (resp *types.FavorResp, err error) {
	// todo: add your logic here and delete this line
	userId, ok := l.ctx.Value(consts.UserId).(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	fmt.Println("user id =", userId)

	switch req.Action_type {
	case 1:
		//没有点赞，点赞操作
		_, err = l.svcCtx.FavorRpc.Favor(l.ctx, &favor.FavorReq{
			Userid:  userId,
			VideoId: req.Video_id,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
	case 2:
		//已经点赞，取消点赞
		_, err = l.svcCtx.FavorRpc.DisFavor(l.ctx, &favor.DisFavorReq{
			UserId:  userId,
			VideoId: req.Video_id,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
	}

	return &types.FavorResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success",
		},
	}, nil
}
