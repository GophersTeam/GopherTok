package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/favor/model"
	"context"
	"github.com/redis/go-redis/v9"

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
		if err == redis.Nil {
			// 查询mysql
			var count int64
			err := l.svcCtx.DB.Model(&model.Info{}).Where("userid = ? AND videoid = ?", in.UserId, in.VideoId).Count(&count).Error
			if err != nil {
				return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "err:%v", err)
			}
			if count > 0 {
				// 插入redis
				l.svcCtx.FavorModel.Insert(l.ctx, in.UserId, in.VideoId)
				return &favor.IsFavorResp{
					IsFavor: true,
				}, nil
			} else {
				return &favor.IsFavorResp{
					IsFavor: false,
				}, nil
			}
		}
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "err:%v", err)
	}
	return &favor.IsFavorResp{
		IsFavor: isFavor,
	}, nil
}
