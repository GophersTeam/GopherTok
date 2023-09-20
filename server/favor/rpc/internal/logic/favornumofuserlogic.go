package logic

import (
	"GopherTok/server/favor/model"
	"context"
	"github.com/redis/go-redis/v9"

	"GopherTok/common/errorx"

	"github.com/pkg/errors"

	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorNumOfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavorNumOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorNumOfUserLogic {
	return &FavorNumOfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavorNumOfUserLogic) FavorNumOfUser(in *favor.FavorNumOfUserReq) (*favor.FavorNumOfUserResp, error) {
	// todo: add your logic here and delete this line
	num, err := l.svcCtx.FavorModel.FavorNumOfUser(l.ctx, in.UserId)
	if err != nil {
		if err == redis.Nil {
			var count int64
			err := l.svcCtx.DB.Model(&model.Info{}).Where("userid = ?", in.UserId).Count(&count).Error
			if err != nil {
				return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "err:%v", err)
			}
			return &favor.FavorNumOfUserResp{
				FavorNumOfUser: int64(count),
			}, nil
		}
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "err:%v", err)
	}

	return &favor.FavorNumOfUserResp{
		FavorNumOfUser: int64(num),
	}, nil
}
