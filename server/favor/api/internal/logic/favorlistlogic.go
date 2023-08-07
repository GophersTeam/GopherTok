package logic

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/favor/rpc/types/favor"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"GopherTok/server/favor/api/internal/svc"
	"GopherTok/server/favor/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorListLogic {
	return &FavorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavorListLogic) FavorList(req *types.FavorlistReq) (resp *types.FavorlistResp, err error) {
	// todo: add your logic here and delete this line
	userId, ok := l.ctx.Value(consts.UserId).(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	fmt.Println("user id =", userId)

	list, err := l.svcCtx.FavorRpc.FavorList(l.ctx, &favor.FavorListReq{
		Userid: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.FavorlistResp{
		BaseResponse: types.BaseResponse{},
		Video_ids:    list.VideoIds,
	}, nil
}
