package logic

import (
	"context"

	"GopherTok/common/errorx"

	"github.com/pkg/errors"

	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserVideoIdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserVideoIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVideoIdListLogic {
	return &GetUserVideoIdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserVideoIdListLogic) GetUserVideoIdList(in *video.GetUserVideoIdListReq) (*video.GetUserVideoIdListResp, error) {
	// todo: add your logic here and delete this line

	vIds, err := l.svcCtx.VideoModel.FindVideoIdsByUserId(l.ctx, l.svcCtx.Rdb, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("服务查询出错，err:"+err.Error()), "mysql查询出错，err:%v", err)
	}
	IdList := make([]int64, 0)
	for _, v := range vIds {
		IdList = append(IdList, v)
	}
	return &video.GetUserVideoIdListResp{
		VideoIdList: IdList,
	}, nil
}
