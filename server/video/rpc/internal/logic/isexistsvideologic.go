package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

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
	isExists := false
	_, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
	// 判断记录是否存在
	if errors.Is(err, sqlc.ErrNotFound) {
		logc.Info(l.ctx, "ID %d 的记录不存在\n", in.VideoId)
	} else if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql查询出错，err:"+err.Error()), "mysql查询出错，err:%v", err.Error)
	} else {
		logc.Info(l.ctx, "Video_ID %d 的记录存在\n", in.VideoId)
		isExists = true
	}
	return &video.IsExistsVideoResp{
		IsExists: isExists,
	}, nil
}
