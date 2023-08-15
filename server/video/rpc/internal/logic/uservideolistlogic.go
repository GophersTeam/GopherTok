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

type UserVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVideoListLogic {
	return &UserVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserVideoListLogic) UserVideoList(in *video.UserVideoListReq) (*video.UserVideoListResp, error) {
	// todo: add your logic here and delete this line
	var list []model.Video
	err := l.svcCtx.MysqlDb.Where("user_id = ?", in.UserId).Order("create_time DESC").Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql find 错误"+err.Error()), "mysql find err:%v", err)
	}

	videoList := make([]*video.VideoList, 0) // Assuming VideoList is a struct that matches your needs

	for i := 0; i < len(list); i++ {
		videoItem := &video.VideoList{
			Id:          list[i].ID,
			UserId:      list[i].UserID,
			Title:       list[i].Title,
			PlayUrl:     list[i].PlayURL,
			CoverUrl:    list[i].CoverURL,
			CreateTime:  list[i].CreateTime.Unix(),
			UpdateTime:  list[i].UpdateTime.Unix(),
			VideoSha256: list[i].VideoSha256,
		}
		videoList = append(videoList, videoItem)
	}
	return &video.UserVideoListResp{
		VideoList: videoList,
	}, nil
}
