package logic

import (
	"context"

	"GopherTok/server/chat/rpc/internal/svc"
	"GopherTok/server/chat/rpc/pb"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageChatLogic) MessageChat(in *pb.MessageChatRequest) (resp *pb.MessageChatResponse, err error) {
	messages, err := l.svcCtx.MessageModel.GetMessages(l.ctx, in.FromUserId, in.ToUserId, in.PreMsgTime)
	if err != nil {
		l.Errorf("MessageChat error: %s", err.Error())
		return nil, err
	}

	messageList := make([]*pb.Message, 0, len(messages))
	resp = new(pb.MessageChatResponse)
	copier.Copy(&messageList, &messages)
	resp.MessageList = messageList

	l.Info(resp.MessageList)

	return
}
