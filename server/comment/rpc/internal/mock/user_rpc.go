package mock

import (
	"GopherTok/server/user/rpc/userclient"
	"context"
	"google.golang.org/grpc"
)

type UserRpc struct {
}

func (u UserRpc) Register(ctx context.Context, in *userclient.RegisterReq, opts ...grpc.CallOption) (*userclient.RegisterResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRpc) Login(ctx context.Context, in *userclient.LoginReq, opts ...grpc.CallOption) (*userclient.LoginResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRpc) UserInfo(ctx context.Context, in *userclient.UserInfoReq, opts ...grpc.CallOption) (*userclient.UserInfoResp, error) {
	return &userclient.UserInfoResp{
		Id:              in.Id,
		Name:            "zhangsan",
		Avatar:          "https://picsum.photos/200",
		BackgroundImage: "https://picsum.photos/400/500",
		Signature:       "这个人很懒，什么都没有留下",
		IsFollow:        true,
		FollowCount:     120,
		FollowerCount:   200,
		WorkCount:       24,
		FavoriteCount:   100000,
	}, nil
}

func (u UserRpc) AddCount(ctx context.Context, in *userclient.AddCountReq, opts ...grpc.CallOption) (*userclient.AddCountResp, error) {
	//TODO implement me
	panic("implement me")
}
