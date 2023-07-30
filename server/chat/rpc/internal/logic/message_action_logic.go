package logic

import (
	"GopherTok/common/consts"
	"GopherTok/server/chat/model"
	"GopherTok/server/chat/rpc/internal/svc"
	"GopherTok/server/chat/rpc/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
	"strconv"
	"time"
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
	_, err = l.svcCtx.MessageModel.Insert(l.ctx, message)
	if err != nil {
		l.Errorf("MessageAction MessageModel.Insert error: %s", err.Error())
		return nil, err
	}

	// 保存最新消息到redis
	fromUserID := strconv.Itoa(int(message.FromUserId))
	toUserID := strconv.Itoa(int(message.ToUserId))

	lastMessage := &pb.LastMessage{Content: message.Content}
	lastMessage.MsgType = consts.MsgTypeRecv
	lastMessageRecvBytes, _ := proto.Marshal(lastMessage)
	//lastMessageRecvStr, _ := jsonx.MarshalToString(lastMessage)
	lastMessage.MsgType = consts.MsgTypeSend
	lastMessageSendBytes, _ := proto.Marshal(lastMessage)
	//lastMessageSendStr, err := jsonx.MarshalToString(lastMessage)

	_ = l.svcCtx.RedisClient.HsetCtx(l.ctx, consts.LastMessagePrefix+fromUserID, toUserID, string(lastMessageSendBytes))
	_ = l.svcCtx.RedisClient.HsetCtx(l.ctx, consts.LastMessagePrefix+toUserID, fromUserID, string(lastMessageRecvBytes))

	resp = new(pb.MessageActionResponse)

	return
}
