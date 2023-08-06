package mock

import (
	"GopherTok/server/video/rpc/types/video"
	"GopherTok/server/video/rpc/videoclient"
	"context"
	"google.golang.org/grpc"
	"time"
)

type VideoRpc struct {
}

func (v VideoRpc) PublishVideo(ctx context.Context, in *videoclient.PublishVideoReq, opts ...grpc.CallOption) (*videoclient.CommonResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) UserVideoList(ctx context.Context, in *videoclient.UserVideoListReq, opts ...grpc.CallOption) (*videoclient.UserVideoListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) VideoList(ctx context.Context, in *videoclient.VideoListReq, opts ...grpc.CallOption) (*videoclient.VideoListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) IsExistsVideo(ctx context.Context, in *videoclient.IsExistsVideoReq, opts ...grpc.CallOption) (*videoclient.IsExistsVideoResp, error) {
	//TODO implement me
	return &video.IsExistsVideoResp{
		IsExists: false,
	}, nil
}

func (v VideoRpc) FindVideo(ctx context.Context, in *videoclient.FindVideoReq, opts ...grpc.CallOption) (*videoclient.FindVideoResp, error) {
	//TODO implement me
	return &video.FindVideoResp{
		Video: &video.VideoList{
			Id:          234,
			UserId:      5234,
			Title:       "今天吃饭了吗",
			PlayUrl:     "http://xxxx.com",
			CoverUrl:    "http://zzzz.com",
			CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
			UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
			VideoSha256: "ghiuerhguwhgpfjpsefjepruhg",
		},
	}, nil
}
