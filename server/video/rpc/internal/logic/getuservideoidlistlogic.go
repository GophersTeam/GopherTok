package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/video/model"
	"context"
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
	var list []model.Video
	err := l.svcCtx.MysqlDb.Select("id").Where("user_id = ?", in.UserId).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql find 错误"+err.Error()), "mysql find err:%v", err)
	}
	IdList := make([]int64, 0)
	for _, v := range list {
		IdList = append(IdList, v.ID)
	}
	return &video.GetUserVideoIdListResp{
		VideoIdList: IdList,
	}, nil
}
