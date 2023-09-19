package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"GopherTok/common/batcher"
	"GopherTok/common/errorx"
	"GopherTok/server/video/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	batcher *batcher.Batcher
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	f := &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
	// batcher配置
	options := batcher.Options{
		5,
		100,
		100,
		5 * time.Second,
	}
	// 实现batcher
	b := batcher.New(options)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % options.Worker
	}
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*model.Video
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*model.Video))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Errorf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		if err = f.svcCtx.KqPusherClient.Push(string(kd)); err != nil {
			logx.Errorf("KafkaPusher.Push kd: %s error: %v", string(kd), err)
		}
	}
	f.batcher = b
	f.batcher.Start()
	return f
}

func (l *PublishVideoLogic) PublishVideo(in *video.PublishVideoReq) (*video.CommonResp, error) {
	// todo: add your logic here and delete this line
	CreateTime, err := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	if err != nil {
		logc.Error(l.ctx, "Error parsing time:", err)
		return nil, errors.Wrapf(errorx.NewDefaultError("Error parsing time"), "Error parsing time:%v", err)
	}
	UpdateTime, err := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	if err != nil {
		logc.Error(l.ctx, "Error parsing time:", err)
		return nil, errors.Wrapf(errorx.NewDefaultError("Error parsing time"), "Error parsing time:%v", err)
	}
	fmt.Println(CreateTime)
	fmt.Println("---")
	fmt.Println(UpdateTime)

	v := model.Video{
		Id:          in.Id,
		UserId:      in.UserId,
		Title:       in.Title,
		PlayUrl:     in.PlayUrl,
		CoverUrl:    in.CoverUrl,
		CreateTime:  CreateTime,
		UpdateTime:  UpdateTime,
		VideoSha256: in.VideoSha256,
	}

	// kafka异步处理file元数据
	err = l.batcher.Add(strconv.FormatInt(in.UserId, 10), &v)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(40003, errorx.ErrKafkaUserFileMeta+err.Error()), "kafka异步UserFileMeta失败 err:%v", err)
	}
	fmt.Println(v)
	return &video.CommonResp{}, nil
}
