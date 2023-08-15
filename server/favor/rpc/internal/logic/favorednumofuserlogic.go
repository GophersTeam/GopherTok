package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/video/rpc/videoclient"
	"context"
	"github.com/pkg/errors"

	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoredNumOfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoredNumOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoredNumOfUserLogic {
	return &FavoredNumOfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoredNumOfUserLogic) FavoredNumOfUser(in *favor.FavoredNumOfUserReq) (*favor.FavoredNumOfUserResp, error) {

	list, err := l.svcCtx.VideoRpc.UserVideoList(l.ctx, &videoclient.UserVideoListReq{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "err:%v", err)

	}
	var sum int = 0
	for _, j := range list.VideoList {
		num, _ := l.svcCtx.FavorModel.NumOfFavor(l.ctx, j.Id)
		sum += num
	}

	return &favor.FavoredNumOfUserResp{
		FavoredNumOfUser: int64(sum),
	}, nil
}
