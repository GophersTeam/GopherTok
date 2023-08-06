// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoclient

import (
	"context"

	"GopherTok/server/video/rpc/types/video"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonResp        = video.CommonResp
	FindVideoReq      = video.FindVideoReq
	FindVideoResp     = video.FindVideoResp
	IsExistsVideoReq  = video.IsExistsVideoReq
	IsExistsVideoResp = video.IsExistsVideoResp
	PublishVideoReq   = video.PublishVideoReq
	UserVideoListReq  = video.UserVideoListReq
	UserVideoListResp = video.UserVideoListResp
	VideoList         = video.VideoList
	VideoListReq      = video.VideoListReq
	VideoListResp     = video.VideoListResp

	Video interface {
		PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*CommonResp, error)
		UserVideoList(ctx context.Context, in *UserVideoListReq, opts ...grpc.CallOption) (*UserVideoListResp, error)
		VideoList(ctx context.Context, in *VideoListReq, opts ...grpc.CallOption) (*VideoListResp, error)
		IsExistsVideo(ctx context.Context, in *IsExistsVideoReq, opts ...grpc.CallOption) (*IsExistsVideoResp, error)
		FindVideo(ctx context.Context, in *FindVideoReq, opts ...grpc.CallOption) (*FindVideoResp, error)
	}

	defaultVideo struct {
		cli zrpc.Client
	}
)

func NewVideo(cli zrpc.Client) Video {
	return &defaultVideo{
		cli: cli,
	}
}

func (m *defaultVideo) PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}

func (m *defaultVideo) UserVideoList(ctx context.Context, in *UserVideoListReq, opts ...grpc.CallOption) (*UserVideoListResp, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.UserVideoList(ctx, in, opts...)
}

func (m *defaultVideo) VideoList(ctx context.Context, in *VideoListReq, opts ...grpc.CallOption) (*VideoListResp, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.VideoList(ctx, in, opts...)
}

func (m *defaultVideo) IsExistsVideo(ctx context.Context, in *IsExistsVideoReq, opts ...grpc.CallOption) (*IsExistsVideoResp, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.IsExistsVideo(ctx, in, opts...)
}

func (m *defaultVideo) FindVideo(ctx context.Context, in *FindVideoReq, opts ...grpc.CallOption) (*FindVideoResp, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.FindVideo(ctx, in, opts...)
}
