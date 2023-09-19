package logic

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/video/model"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
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
	// 毫秒
	timestampTime := time.Now()
	if in.LatestTime == "" {
		logc.Info(l.ctx, "未选择lastime,req:", in)
	} else {
		timestampInt, err := strconv.ParseInt(in.LatestTime, 10, 64)
		if err != nil {
			fmt.Println("解析时间戳出错:", err)
			return nil, errors.Wrapf(errorx.NewDefaultError("解析时间戳出错"+err.Error()), "解析时间戳出错%v", err)
		}

		// 使用Unix秒数创建time.Time类型
		timestampTime = time.Unix(0, timestampInt*int64(time.Millisecond))

	}
	// 在redis中查询出时间戳最大的30个视频id
	VideoIds, err := l.svcCtx.Rdb.ZRevRangeWithScores(context.Background(), consts.AllVideoId, 0, 29).Result()
	if err != nil {
		// 处理错误
		return nil, errors.Wrapf(errorx.NewDefaultError("redis 查询错误"+err.Error()), "redis 查询错误%v", err)

	}
	list := make([]*model.Video, 0)
	for _, v := range VideoIds {
		oneVideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, v.Member.(int64))
		if err != nil {
			return nil, errors.Wrapf(errorx.NewDefaultError("mysql find 错误"+err.Error()), "mysql find err:%v", err)
		}
		list = append(list, oneVideo)
	}

	videoList := make([]*video.VideoList, 0) // Assuming VideoList is a struct that matches your needs

	for i := 0; i < len(list); i++ {
		videoItem := &video.VideoList{
			Id:          list[i].Id,
			UserId:      list[i].UserId,
			Title:       list[i].Title,
			PlayUrl:     list[i].PlayUrl,
			CoverUrl:    list[i].CoverUrl,
			CreateTime:  list[i].CreateTime.Unix(),
			UpdateTime:  list[i].UpdateTime.Unix(),
			VideoSha256: list[i].VideoSha256,
		}
		videoList = append(videoList, videoItem)
	}

	return &video.VideoListResp{
		VideoList: videoList,
	}, nil
}
