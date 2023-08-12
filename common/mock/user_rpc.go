package mock

import (
	"GopherTok/server/user/rpc/types/user"
	"GopherTok/server/user/rpc/userclient"
	"context"
	"google.golang.org/grpc"
)

type UserRpc struct {
}

func (u UserRpc) UserIsExists(ctx context.Context, in *userclient.UserIsExistsReq, opts ...grpc.CallOption) (*userclient.UserIsExistsResp, error) {
	//TODO implement me
	return &userclient.UserIsExistsResp{Exists: false}, nil
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
	//TODO implement me
	return &userclient.UserInfoResp{
		Id:              in.Id,
		Name:            "hhhh",
		Avatar:          "htpp://xxx",
		BackgroundImage: "htpp://xxx",
		Signature:       "htpp://xxx",
		IsFollow:        false,
		FollowCount:     23,
		FollowerCount:   33,
		TotalFavorited:  "444",
		WorkCount:       23,
		FavoriteCount:   70,
	}, nil
}

func (u UserRpc) AddCount(ctx context.Context, in *userclient.AddCountReq, opts ...grpc.CallOption) (*userclient.AddCountResp, error) {
	//TODO implement me
	return &user.AddCountResp{}, nil

}
