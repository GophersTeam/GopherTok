package mock

import (
	"context"

	"GopherTok/server/comment/rpc/commentrpc"

	"google.golang.org/grpc"
)

type CommentRpc struct{}

func (c CommentRpc) AddComment(ctx context.Context, in *commentrpc.AddCommentRequest, opts ...grpc.CallOption) (*commentrpc.AddCommentResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c CommentRpc) DelComment(ctx context.Context, in *commentrpc.DelCommentRequest, opts ...grpc.CallOption) (*commentrpc.DelCommentResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c CommentRpc) GetCommentList(ctx context.Context, in *commentrpc.GetCommentListRequest, opts ...grpc.CallOption) (*commentrpc.GetCommentListResponse, error) {
	// TODO implement me
	panic("implement me")
}

// GetCommentCount 获取视频的评论数
func (c CommentRpc) GetCommentCount(ctx context.Context, in *commentrpc.GetCommentCountRequest, opts ...grpc.CallOption) (*commentrpc.GetCommentCountResponse, error) {
	return &commentrpc.GetCommentCountResponse{
		Count: 100,
	}, nil
}
