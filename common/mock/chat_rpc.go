package mock

import (
	"context"
	"strconv"

	"GopherTok/server/chat/rpc/chatrpc"

	"google.golang.org/grpc"
)

type ChatRpc struct{}

func (c ChatRpc) MessageAction(ctx context.Context, in *chatrpc.MessageActionRequest, opts ...grpc.CallOption) (*chatrpc.MessageActionResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c ChatRpc) MessageChat(ctx context.Context, in *chatrpc.MessageChatRequest, opts ...grpc.CallOption) (*chatrpc.MessageChatResponse, error) {
	// TODO implement me
	panic("implement me")
}

// MessageChatLast 获取和每个好友的最后一条消息
func (c ChatRpc) MessageChatLast(ctx context.Context, in *chatrpc.MessageChatLastRequest, opts ...grpc.CallOption) (*chatrpc.MessageChatLastResponse, error) {
	resp := new(chatrpc.MessageChatLastResponse)
	resp.LastMessageList = make([]*chatrpc.LastMessage, len(in.ToUserIdList))
	msgType := int64(1)
	for i := 0; i < len(resp.LastMessageList); i++ {
		resp.LastMessageList[i] = &chatrpc.LastMessage{
			Content: "hello " + strconv.Itoa(int(in.ToUserIdList[i])),
			MsgType: msgType,
		}

		if i%2 == 0 {
			msgType = 1
		} else {
			msgType = 0
		}
	}
	return resp, nil
}
