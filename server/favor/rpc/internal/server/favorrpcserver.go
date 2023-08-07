// Code generated by goctl. DO NOT EDIT.
// Source: favor.proto

package server

import (
	"context"

	"GopherTok/server/favor/rpc/internal/logic"
	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"
)

type FavorRpcServer struct {
	svcCtx *svc.ServiceContext
	favor.UnimplementedFavorRpcServer
}

func NewFavorRpcServer(svcCtx *svc.ServiceContext) *FavorRpcServer {
	return &FavorRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *FavorRpcServer) DisFavor(ctx context.Context, in *favor.DisFavorReq) (*favor.DisFavorResp, error) {
	l := logic.NewDisFavorLogic(ctx, s.svcCtx)
	return l.DisFavor(in)
}

func (s *FavorRpcServer) Favor(ctx context.Context, in *favor.FavorReq) (*favor.FavorResp, error) {
	l := logic.NewFavorLogic(ctx, s.svcCtx)
	return l.Favor(in)
}

func (s *FavorRpcServer) FavorList(ctx context.Context, in *favor.FavorListReq) (*favor.FavorListResp, error) {
	l := logic.NewFavorListLogic(ctx, s.svcCtx)
	return l.FavorList(in)
}

func (s *FavorRpcServer) FavorNum(ctx context.Context, in *favor.FavorNumReq) (*favor.FavorNumResp, error) {
	l := logic.NewFavorNumLogic(ctx, s.svcCtx)
	return l.FavorNum(in)
}

func (s *FavorRpcServer) IsFavor(ctx context.Context, in *favor.IsFavorReq) (*favor.IsFavorResp, error) {
	l := logic.NewIsFavorLogic(ctx, s.svcCtx)
	return l.IsFavor(in)
}

func (s *FavorRpcServer) FavorNumOfUser(ctx context.Context, in *favor.FavorNumOfUserReq) (*favor.FavorNumOfUserResp, error) {
	l := logic.NewFavorNumOfUserLogic(ctx, s.svcCtx)
	return l.FavorNumOfUser(in)
}

func (s *FavorRpcServer) FavoredNumOfUser(ctx context.Context, in *favor.FavoredNumOfUserReq) (*favor.FavoredNumOfUserResp, error) {
	l := logic.NewFavoredNumOfUserLogic(ctx, s.svcCtx)
	return l.FavoredNumOfUser(in)
}
