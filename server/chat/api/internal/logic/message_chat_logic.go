package logic

import (
	"GopherTok/server/chat/rpc/pb"
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"GopherTok/server/chat/api/internal/svc"
	"GopherTok/server/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageChatLogic) MessageChat(req *types.MessageChatRequest) (resp *types.MessageChatResponse, err error) {
	fromUserId := l.ctx.Value("userId").(int64)

	if fromUserId == req.ToUserId {
		return nil, errors.New("不能查看自己的消息记录")
	}

	chatResp, err := l.svcCtx.ChatRpc.MessageChat(l.ctx, &pb.MessageChatRequest{
		FromUserId: fromUserId,
		ToUserId:   req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	})

	if err != nil {
		l.Errorf("MessageChat error: %s", err.Error())
		return nil, err
	}

	resp = new(types.MessageChatResponse)
	resp.MessageList = make([]*types.Message, 0, len(chatResp.MessageList))
	_ = copier.Copy(&resp.MessageList, &chatResp.MessageList)

	return
}
