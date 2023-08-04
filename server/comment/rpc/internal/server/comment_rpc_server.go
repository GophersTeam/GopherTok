// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"GopherTok/server/comment/rpc/internal/logic"
	"GopherTok/server/comment/rpc/internal/svc"
	"GopherTok/server/comment/rpc/pb"
)

type CommentRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCommentRpcServer
}

func NewCommentRpcServer(svcCtx *svc.ServiceContext) *CommentRpcServer {
	return &CommentRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentRpcServer) AddComment(ctx context.Context, in *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	l := logic.NewAddCommentLogic(ctx, s.svcCtx)
	return l.AddComment(in)
}

func (s *CommentRpcServer) DelComment(ctx context.Context, in *pb.DelCommentRequest) (*pb.DelCommentResponse, error) {
	l := logic.NewDelCommentLogic(ctx, s.svcCtx)
	return l.DelComment(in)
}

func (s *CommentRpcServer) GetCommentList(ctx context.Context, in *pb.GetCommentListRequest) (*pb.GetCommentListResponse, error) {
	l := logic.NewGetCommentListLogic(ctx, s.svcCtx)
	return l.GetCommentList(in)
}

func (s *CommentRpcServer) GetCommentCount(ctx context.Context, in *pb.GetCommentCountRequest) (*pb.GetCommentCountResponse, error) {
	l := logic.NewGetCommentCountLogic(ctx, s.svcCtx)
	return l.GetCommentCount(in)
}
