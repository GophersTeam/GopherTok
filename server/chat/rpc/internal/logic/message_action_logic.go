package logic

import (
	"context"
	"time"

	"GopherTok/server/chat/model"
	"GopherTok/server/chat/rpc/internal/svc"
	"GopherTok/server/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageActionLogic) MessageAction(in *pb.MessageActionRequest) (resp *pb.MessageActionResponse, err error) {
	message := &model.Message{
		Id:         l.svcCtx.Snowflake.Generate().Int64(),
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Content:    in.Content,
		CreateTime: time.Now().Unix(),
	}
	messageStr, err := jsonx.MarshalToString(message)
	err = l.svcCtx.KafkaPusher.Push(messageStr)
	if err != nil {
		return nil, err
	}

	resp = new(pb.MessageActionResponse)

	return
}
