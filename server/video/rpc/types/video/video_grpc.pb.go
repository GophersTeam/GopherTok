// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: video.proto

package video

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Video_PublishVideo_FullMethodName       = "/video.video/PublishVideo"
	Video_UserVideoList_FullMethodName      = "/video.video/UserVideoList"
	Video_VideoList_FullMethodName          = "/video.video/VideoList"
	Video_IsExistsVideo_FullMethodName      = "/video.video/IsExistsVideo"
	Video_FindVideo_FullMethodName          = "/video.video/FindVideo"
	Video_GetUserVideoIdList_FullMethodName = "/video.video/GetUserVideoIdList"
)

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*CommonResp, error)
	UserVideoList(ctx context.Context, in *UserVideoListReq, opts ...grpc.CallOption) (*UserVideoListResp, error)
	VideoList(ctx context.Context, in *VideoListReq, opts ...grpc.CallOption) (*VideoListResp, error)
	IsExistsVideo(ctx context.Context, in *IsExistsVideoReq, opts ...grpc.CallOption) (*IsExistsVideoResp, error)
	FindVideo(ctx context.Context, in *FindVideoReq, opts ...grpc.CallOption) (*FindVideoResp, error)
	GetUserVideoIdList(ctx context.Context, in *GetUserVideoIdListReq, opts ...grpc.CallOption) (*GetUserVideoIdListResp, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Video_PublishVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) UserVideoList(ctx context.Context, in *UserVideoListReq, opts ...grpc.CallOption) (*UserVideoListResp, error) {
	out := new(UserVideoListResp)
	err := c.cc.Invoke(ctx, Video_UserVideoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) VideoList(ctx context.Context, in *VideoListReq, opts ...grpc.CallOption) (*VideoListResp, error) {
	out := new(VideoListResp)
	err := c.cc.Invoke(ctx, Video_VideoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) IsExistsVideo(ctx context.Context, in *IsExistsVideoReq, opts ...grpc.CallOption) (*IsExistsVideoResp, error) {
	out := new(IsExistsVideoResp)
	err := c.cc.Invoke(ctx, Video_IsExistsVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) FindVideo(ctx context.Context, in *FindVideoReq, opts ...grpc.CallOption) (*FindVideoResp, error) {
	out := new(FindVideoResp)
	err := c.cc.Invoke(ctx, Video_FindVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) GetUserVideoIdList(ctx context.Context, in *GetUserVideoIdListReq, opts ...grpc.CallOption) (*GetUserVideoIdListResp, error) {
	out := new(GetUserVideoIdListResp)
	err := c.cc.Invoke(ctx, Video_GetUserVideoIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	PublishVideo(context.Context, *PublishVideoReq) (*CommonResp, error)
	UserVideoList(context.Context, *UserVideoListReq) (*UserVideoListResp, error)
	VideoList(context.Context, *VideoListReq) (*VideoListResp, error)
	IsExistsVideo(context.Context, *IsExistsVideoReq) (*IsExistsVideoResp, error)
	FindVideo(context.Context, *FindVideoReq) (*FindVideoResp, error)
	GetUserVideoIdList(context.Context, *GetUserVideoIdListReq) (*GetUserVideoIdListResp, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) PublishVideo(context.Context, *PublishVideoReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}
func (UnimplementedVideoServer) UserVideoList(context.Context, *UserVideoListReq) (*UserVideoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserVideoList not implemented")
}
func (UnimplementedVideoServer) VideoList(context.Context, *VideoListReq) (*VideoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoList not implemented")
}
func (UnimplementedVideoServer) IsExistsVideo(context.Context, *IsExistsVideoReq) (*IsExistsVideoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsExistsVideo not implemented")
}
func (UnimplementedVideoServer) FindVideo(context.Context, *FindVideoReq) (*FindVideoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindVideo not implemented")
}
func (UnimplementedVideoServer) GetUserVideoIdList(context.Context, *GetUserVideoIdListReq) (*GetUserVideoIdListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserVideoIdList not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_PublishVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).PublishVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_PublishVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).PublishVideo(ctx, req.(*PublishVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_UserVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVideoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).UserVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_UserVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).UserVideoList(ctx, req.(*UserVideoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_VideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).VideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_VideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).VideoList(ctx, req.(*VideoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_IsExistsVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsExistsVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).IsExistsVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_IsExistsVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).IsExistsVideo(ctx, req.(*IsExistsVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_FindVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).FindVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_FindVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).FindVideo(ctx, req.(*FindVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_GetUserVideoIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserVideoIdListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).GetUserVideoIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_GetUserVideoIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).GetUserVideoIdList(ctx, req.(*GetUserVideoIdListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishVideo",
			Handler:    _Video_PublishVideo_Handler,
		},
		{
			MethodName: "UserVideoList",
			Handler:    _Video_UserVideoList_Handler,
		},
		{
			MethodName: "VideoList",
			Handler:    _Video_VideoList_Handler,
		},
		{
			MethodName: "IsExistsVideo",
			Handler:    _Video_IsExistsVideo_Handler,
		},
		{
			MethodName: "FindVideo",
			Handler:    _Video_FindVideo_Handler,
		},
		{
			MethodName: "GetUserVideoIdList",
			Handler:    _Video_GetUserVideoIdList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
