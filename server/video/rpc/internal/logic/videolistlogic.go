package logic

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/mr"

	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/video/model"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logc"

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

	// 使用ZREVRANGEBYSCORE命令查询score小于timestampTime的最大30个成员
	vIds, err := l.svcCtx.Rdb.ZRevRangeByScoreWithScores(l.ctx, consts.AllVideoIdPrefix, &redis.ZRangeBy{
		Min:    "0",
		Max:    strconv.FormatInt(timestampTime.Unix(), 10),
		Offset: 0,
		Count:  30,
	}).Result()
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("redis ZREVRANGEBYSCORE错误"+err.Error()), "redis ZREVRANGEBYSCORE错误%v", err)
	}

	// 根据id走缓存查询
	// 使用MapReduce并发查询缓存
	list, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, v := range vIds {
			source <- v.Member
		}
	}, func(item interface{}, writer mr.Writer[*model.Video], cancel func(error)) {
		oneVideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, item.(int64))
		if err != nil {
			cancel(err)
		}
		writer.Write(oneVideo)
	}, func(pipe <-chan *model.Video, writer mr.Writer[[]*model.Video], cancel func(error)) {
		lists := make([]*model.Video, 0)
		for p := range pipe {
			lists = append(lists, p)
		}
		writer.Write(lists)
	})
	// 时间逆序排序
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreateTime.After(list[j].CreateTime)
	})
	//for _, v := range vIds {
	//	oneVideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, v.Member.(int64))
	//	if err != nil {
	//		return nil, errors.Wrapf(errorx.NewDefaultError("mysql find 错误"+err.Error()), "mysql find err:%v", err)
	//	}
	//	list = append(list, oneVideo)
	//}

	videoList := make([]*video.VideoList, 0)

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
