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
		l.Errorf("MessageChat MessageModel.MGetMessages error: %s", err.Error())
		return nil, err
	}

	resp = new(pb.MessageChatResponse)
	resp.MessageList = make([]*pb.Message, 0, len(messages))
	_ = copier.Copy(&resp.MessageList, &messages)

	return
}
