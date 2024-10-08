// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: interact.proto

package rpc_interact

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	InteractServiceImpl_FavoriteAction_FullMethodName       = "/InteractServiceImpl/FavoriteAction"
	InteractServiceImpl_GetFavoriteVideoList_FullMethodName = "/InteractServiceImpl/GetFavoriteVideoList"
	InteractServiceImpl_CommentAction_FullMethodName        = "/InteractServiceImpl/CommentAction"
	InteractServiceImpl_GetCommentList_FullMethodName       = "/InteractServiceImpl/GetCommentList"
)

// InteractServiceImplClient is the client API for InteractServiceImpl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractServiceImplClient interface {
	FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error)
	GetFavoriteVideoList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error)
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	GetCommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
}

type interactServiceImplClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractServiceImplClient(cc grpc.ClientConnInterface) InteractServiceImplClient {
	return &interactServiceImplClient{cc}
}

func (c *interactServiceImplClient) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FavoriteActionResponse)
	err := c.cc.Invoke(ctx, InteractServiceImpl_FavoriteAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceImplClient) GetFavoriteVideoList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FavoriteListResponse)
	err := c.cc.Invoke(ctx, InteractServiceImpl_GetFavoriteVideoList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceImplClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, InteractServiceImpl_CommentAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceImplClient) GetCommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentListResponse)
	err := c.cc.Invoke(ctx, InteractServiceImpl_GetCommentList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractServiceImplServer is the server API for InteractServiceImpl service.
// All implementations must embed UnimplementedInteractServiceImplServer
// for forward compatibility
type InteractServiceImplServer interface {
	FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionResponse, error)
	GetFavoriteVideoList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error)
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	GetCommentList(context.Context, *CommentListRequest) (*CommentListResponse, error)
	mustEmbedUnimplementedInteractServiceImplServer()
}

// UnimplementedInteractServiceImplServer must be embedded to have forward compatible implementations.
type UnimplementedInteractServiceImplServer struct {
}

func (UnimplementedInteractServiceImplServer) FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedInteractServiceImplServer) GetFavoriteVideoList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteVideoList not implemented")
}
func (UnimplementedInteractServiceImplServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedInteractServiceImplServer) GetCommentList(context.Context, *CommentListRequest) (*CommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedInteractServiceImplServer) mustEmbedUnimplementedInteractServiceImplServer() {}

// UnsafeInteractServiceImplServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractServiceImplServer will
// result in compilation errors.
type UnsafeInteractServiceImplServer interface {
	mustEmbedUnimplementedInteractServiceImplServer()
}

func RegisterInteractServiceImplServer(s grpc.ServiceRegistrar, srv InteractServiceImplServer) {
	s.RegisterService(&InteractServiceImpl_ServiceDesc, srv)
}

func _InteractServiceImpl_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceImplServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractServiceImpl_FavoriteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceImplServer).FavoriteAction(ctx, req.(*FavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractServiceImpl_GetFavoriteVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceImplServer).GetFavoriteVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractServiceImpl_GetFavoriteVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceImplServer).GetFavoriteVideoList(ctx, req.(*FavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractServiceImpl_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceImplServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractServiceImpl_CommentAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceImplServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractServiceImpl_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceImplServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractServiceImpl_GetCommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceImplServer).GetCommentList(ctx, req.(*CommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractServiceImpl_ServiceDesc is the grpc.ServiceDesc for InteractServiceImpl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractServiceImpl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InteractServiceImpl",
	HandlerType: (*InteractServiceImplServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _InteractServiceImpl_FavoriteAction_Handler,
		},
		{
			MethodName: "GetFavoriteVideoList",
			Handler:    _InteractServiceImpl_GetFavoriteVideoList_Handler,
		},
		{
			MethodName: "CommentAction",
			Handler:    _InteractServiceImpl_CommentAction_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _InteractServiceImpl_GetCommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interact.proto",
}
