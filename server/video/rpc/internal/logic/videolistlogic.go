package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/video/model"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoListLogic) VideoList(in *video.VideoListReq) (*video.VideoListResp, error) {
	// todo: add your logic here and delete this line

	timestampInt, err := strconv.ParseInt(in.LatestTime, 10, 64)
	if err != nil {
		fmt.Println("解析时间戳出错:", err)
		return nil, errors.Wrapf(errorx.NewDefaultError("解析时间戳出错"+err.Error()), "解析时间戳出错%v", err)

	}

	// 使用Unix秒数创建time.Time类型
	timestampTime := time.Unix(timestampInt, 0)
	fmt.Println("11111111", timestampTime)
	var list []model.Video
	err = l.svcCtx.MysqlDb.Where("create_time <= ?", timestampTime).Order("create_time DESC").Limit(30).Find(&list).Error
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
			CreateTime:  list[i].CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  list[i].UpdateTime.Format("2006-01-02 15:04:05"),
			VideoSha256: list[i].VideoSha256,
		}
		videoList = append(videoList, videoItem)
	}

	return &video.VideoListResp{
		VideoList: videoList,
	}, nil
}
