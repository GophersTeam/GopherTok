package mock

import (
	"context"
	"time"

	"GopherTok/server/video/rpc/types/video"
	"GopherTok/server/video/rpc/videoclient"

	"google.golang.org/grpc"
)

type VideoRpc struct{}

func (v VideoRpc) PublishVideo(ctx context.Context, in *videoclient.PublishVideoReq, opts ...grpc.CallOption) (*videoclient.CommonResp, error) {
	// TODO implement me
	panic("implement me")
}

func (v VideoRpc) UserVideoList(ctx context.Context, in *videoclient.UserVideoListReq, opts ...grpc.CallOption) (*videoclient.UserVideoListResp, error) {
	// TODO implement me
	panic("implement me")
}

func (v VideoRpc) VideoList(ctx context.Context, in *videoclient.VideoListReq, opts ...grpc.CallOption) (*videoclient.VideoListResp, error) {
	// TODO implement me
	panic("implement me")
}

func (v VideoRpc) IsExistsVideo(ctx context.Context, in *videoclient.IsExistsVideoReq, opts ...grpc.CallOption) (*videoclient.IsExistsVideoResp, error) {
	// TODO implement me
	return &video.IsExistsVideoResp{
		IsExists: false,
	}, nil
}

func (v VideoRpc) FindVideo(ctx context.Context, in *videoclient.FindVideoReq, opts ...grpc.CallOption) (*videoclient.FindVideoResp, error) {
	// TODO implement me
	return &video.FindVideoResp{
		Video: &video.VideoList{
			Id:          234,
			UserId:      5234,
			Title:       "今天吃饭了吗",
			PlayUrl:     "http://xxxx.com",
			CoverUrl:    "http://zzzz.com",
			CreateTime:  time.Now().Unix(),
			UpdateTime:  time.Now().Unix(),
			VideoSha256: "ghiuerhguwhgpfjpsefjepruhg",
		},
	}, nil
}
