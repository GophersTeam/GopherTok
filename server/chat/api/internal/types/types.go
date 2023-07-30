// Code generated by goctl. DO NOT EDIT.
package types

type MessageActionRequest struct {
	ToUserId   int64  `form:"to_user_id" vd:"$>0;msg:'to_user_id error'"`
	ActionType int64  `form:"action_type" vd:"$==1;msg:'action_type error'"`
	Content    string `form:"content"`
}

type MessageActionResponse struct {
}

type MessageChatRequest struct {
	ToUserId   int64 `form:"to_user_id"`
	PreMsgTime int64 `form:"pre_msg_time,optional"`
}

type MessageChatResponse struct {
	MessageList []*Message `json:"message_list"`
}

type Message struct {
	Id         int64  `json:"id"`
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
