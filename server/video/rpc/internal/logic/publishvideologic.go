package logic

import (
	"GopherTok/common/batcher"
	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/video/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"strconv"
	"time"

	"GopherTok/server/video/rpc/internal/svc"
	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	batcher             *batcher.Batcher
	SensitiveWordFilter utils.SensitiveWordFilter
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	trie := utils.NewSensitiveTrie()
	f := &PublishVideoLogic{
		ctx:                 ctx,
		svcCtx:              svcCtx,
		Logger:              logx.WithContext(ctx),
		SensitiveWordFilter: trie,
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
	// 标题敏感过滤
	in.Title = l.SensitiveWordFilter.Filter(in.Title)
	v := model.Video{
		ID:          in.Id,
		UserID:      in.UserId,
		Title:       in.Title,
		PlayURL:     in.PlayUrl,
		CoverURL:    in.CoverUrl,
		CreateTime:  CreateTime,
		UpdateTime:  UpdateTime,
		DeleteTime:  gorm.DeletedAt{},
		VideoSha256: in.VideoSha256,
	}
	fmt.Println(000000, v)
	// kafka异步处理file元数据
	err = l.batcher.Add(strconv.FormatInt(in.UserId, 10), &v)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(40003, errorx.ErrKafkaUserFileMeta+err.Error()), "kafka异步UserFileMeta失败 err:%v", err)
	}
	fmt.Println(v)
	return &video.CommonResp{}, nil
}
