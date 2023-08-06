package mock

import (
	"GopherTok/server/relation/rpc/relationrpc"
	"context"
	"google.golang.org/grpc"
)

type RelationRpc struct {
}

func (r RelationRpc) AddFollow(ctx context.Context, in *relationrpc.AddFollowReq, opts ...grpc.CallOption) (*relationrpc.AddFollowResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.AddFollowResp{StatusCode: "",
		StatusMsg: ""}, nil
}

func (r RelationRpc) DeleteFollow(ctx context.Context, in *relationrpc.DeleteFollowReq, opts ...grpc.CallOption) (*relationrpc.DeleteFollowResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.DeleteFollowResp{StatusCode: "",
		StatusMsg: ""}, nil
}

func (r RelationRpc) GetFollowList(ctx context.Context, in *relationrpc.GetFollowListReq, opts ...grpc.CallOption) (*relationrpc.GetFollowListResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.GetFollowListResp{StatusMsg: "",
		StatusCode: "",
		UserList:   nil}, nil
}

func (r RelationRpc) GetFollowerList(ctx context.Context, in *relationrpc.GetFollowerReq, opts ...grpc.CallOption) (*relationrpc.GetFollowerResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.GetFollowerResp{StatusMsg: "",
		StatusCode: "",
		UserList:   nil}, nil
}

func (r RelationRpc) GetFriendList(ctx context.Context, in *relationrpc.GetFriendListReq, opts ...grpc.CallOption) (*relationrpc.GetFriendListResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.GetFriendListResp{StatusMsg: "",
		StatusCode: "",
		UserList:   nil}, nil
}

func (r RelationRpc) GetFollowerCount(ctx context.Context, in *relationrpc.GetFollowerCountReq, opts ...grpc.CallOption) (*relationrpc.GetFollowerCountResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.GetFollowerCountResp{StatusCode: "",
		StatusMsg: "",
		Count:     0}, nil
}

func (r RelationRpc) GetFollowCount(ctx context.Context, in *relationrpc.GetFollowCountReq, opts ...grpc.CallOption) (*relationrpc.GetFollowCountResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.GetFollowCountResp{StatusCode: "",
		StatusMsg: "",
		Count:     0}, nil
}

func (r RelationRpc) GetFriendCount(ctx context.Context, in *relationrpc.GetFriendCountReq, opts ...grpc.CallOption) (*relationrpc.GetFriendCountResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.GetFriendCountResp{StatusCode: "",
		StatusMsg: "",
		Count:     0}, nil
}

func (r RelationRpc) CheckIsFollow(ctx context.Context, in *relationrpc.CheckIsFollowReq, opts ...grpc.CallOption) (*relationrpc.CheckIsFollowResp, error) {
	//TODO implement me
	panic("implement me")
	return &relationrpc.CheckIsFollowResp{StatusCode: "",
		StatusMsg: "",
		IsFollow:  true}, nil
}
