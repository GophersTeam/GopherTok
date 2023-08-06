package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/video/model"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsExistsVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistsVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistsVideoLogic {
	return &IsExistsVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistsVideoLogic) IsExistsVideo(in *video.IsExistsVideoReq) (*video.IsExistsVideoResp, error) {
	// todo: add your logic here and delete this line
	v := model.Video{}
	isExists := false
	result := l.svcCtx.MysqlDb.Where("id = ?", in.VideoId).First(&v)
	// 判断记录是否存在
	if result.Error == gorm.ErrRecordNotFound {
		logc.Info(l.ctx, "ID %d 的记录不存在\n", in.VideoId)
	} else if result.Error != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql查询出错，err:"+result.Error.Error()), "mysql查询出错，err:%v", result.Error)
	} else {
		logc.Info(l.ctx, "Video_ID %d 的记录存在\n", in.VideoId)
		isExists = true
	}
	return &video.IsExistsVideoResp{
		IsExists: isExists,
	}, nil
}
