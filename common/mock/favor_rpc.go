package mock

import (
	"GopherTok/server/favor/rpc/favorrpc"
	"context"
	"google.golang.org/grpc"
)

type FavorRpc struct {
}

func (f FavorRpc) DisFavor(ctx context.Context, in *favorrpc.DisFavorReq, opts ...grpc.CallOption) (*favorrpc.DisFavorResp, error) {

	return &favorrpc.DisFavorResp{}, nil

}

func (f FavorRpc) Favor(ctx context.Context, in *favorrpc.FavorReq, opts ...grpc.CallOption) (*favorrpc.FavorResp, error) {

	return &favorrpc.FavorResp{}, nil
}

func (f FavorRpc) FavorList(ctx context.Context, in *favorrpc.FavorListReq, opts ...grpc.CallOption) (*favorrpc.FavorListResp, error) {
	return &favorrpc.FavorListResp{
		Videos: nil,
	}, nil
}

func (f FavorRpc) FavorNum(ctx context.Context, in *favorrpc.FavorNumReq, opts ...grpc.CallOption) (*favorrpc.FavorNumResp, error) {
	return &favorrpc.FavorNumResp{
		Num: 0,
	}, nil
}

func (f FavorRpc) IsFavor(ctx context.Context, in *favorrpc.IsFavorReq, opts ...grpc.CallOption) (*favorrpc.IsFavorResp, error) {
	return &favorrpc.IsFavorResp{
		IsFavor: false,
	}, nil
}

func (f FavorRpc) FavorNumOfUser(ctx context.Context, in *favorrpc.FavorNumOfUserReq, opts ...grpc.CallOption) (*favorrpc.FavorNumOfUserResp, error) {
	return &favorrpc.FavorNumOfUserResp{
		FavorNumOfUser: 0,
	}, nil
}

func (f FavorRpc) FavoredNumOfUser(ctx context.Context, in *favorrpc.FavoredNumOfUserReq, opts ...grpc.CallOption) (*favorrpc.FavoredNumOfUserResp, error) {
	return &favorrpc.FavoredNumOfUserResp{
		FavoredNumOfUser: 0,
	}, nil
}
