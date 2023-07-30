package logic

import (
	"GopherTok/common/consts"
	"context"
	"google.golang.org/protobuf/proto"
	"strconv"

	"GopherTok/server/chat/rpc/internal/svc"
	"GopherTok/server/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLastLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageChatLastLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLastLogic {
	return &MessageChatLastLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageChatLastLogic) MessageChatLast(in *pb.MessageChatLastRequest) (resp *pb.MessageChatLastResponse, err error) {
	fromUserIdStr := strconv.Itoa(int(in.FromUserId))
	allLastMessages, err := l.svcCtx.RedisClient.HgetallCtx(l.ctx, consts.LastMessagePrefix+fromUserIdStr)
	if err != nil {
		l.Errorf("MessageChatLast RedisClient.HgetallCtx error: %s", err.Error())
		return nil, err
	}

	resp = new(pb.MessageChatLastResponse)
	resp.LastMessageList = make([]*pb.LastMessage, 0, len(in.ToUserIdList))

	for _, toUserId := range in.ToUserIdList {
		toUserIdStr := strconv.Itoa(int(toUserId))
		lastMessageStr := allLastMessages[toUserIdStr]

		var lastMessage pb.LastMessage
		//_ = jsonx.UnmarshalFromString(lastMessageStr, &lastMessage)
		_ = proto.Unmarshal([]byte(lastMessageStr), &lastMessage)

		resp.LastMessageList = append(resp.LastMessageList, &pb.LastMessage{
			Content: lastMessage.Content,
			MsgType: lastMessage.MsgType,
		})
	}

	return
}
