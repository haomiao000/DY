// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: redis_svr.proto

package redis_svr

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
	RedisSvr_Get_FullMethodName           = "/redis_svr/Get"
	RedisSvr_BatchGet_FullMethodName      = "/redis_svr/BatchGet"
	RedisSvr_Set_FullMethodName           = "/redis_svr/Set"
	RedisSvr_SetWithExpire_FullMethodName = "/redis_svr/SetWithExpire"
	RedisSvr_BatchSet_FullMethodName      = "/redis_svr/BatchSet"
	RedisSvr_SetIfNotExist_FullMethodName = "/redis_svr/SetIfNotExist"
	RedisSvr_Delete_FullMethodName        = "/redis_svr/Delete"
	RedisSvr_LPush_FullMethodName         = "/redis_svr/LPush"
	RedisSvr_RPush_FullMethodName         = "/redis_svr/RPush"
	RedisSvr_LPop_FullMethodName          = "/redis_svr/LPop"
	RedisSvr_RPop_FullMethodName          = "/redis_svr/RPop"
	RedisSvr_LRange_FullMethodName        = "/redis_svr/LRange"
	RedisSvr_SAdd_FullMethodName          = "/redis_svr/SAdd"
	RedisSvr_SRem_FullMethodName          = "/redis_svr/SRem"
	RedisSvr_SISMember_FullMethodName     = "/redis_svr/SISMember"
	RedisSvr_SMembers_FullMethodName      = "/redis_svr/SMembers"
	RedisSvr_SCard_FullMethodName         = "/redis_svr/SCard"
)

// RedisSvrClient is the client API for RedisSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisSvrClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error)
	BatchGet(ctx context.Context, in *BatchGetReq, opts ...grpc.CallOption) (*BatchGetRsp, error)
	Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetRsp, error)
	SetWithExpire(ctx context.Context, in *SetWithExpireReq, opts ...grpc.CallOption) (*SetWithExpireRsp, error)
	BatchSet(ctx context.Context, in *BatchSetReq, opts ...grpc.CallOption) (*BatchSetRsp, error)
	SetIfNotExist(ctx context.Context, in *SetIfNotExistReq, opts ...grpc.CallOption) (*SetIfNotExistRsp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRsp, error)
	// List----------------------------------<
	LPush(ctx context.Context, in *LPushRequest, opts ...grpc.CallOption) (*ListResponse, error)
	RPush(ctx context.Context, in *RPushRequest, opts ...grpc.CallOption) (*ListResponse, error)
	LPop(ctx context.Context, in *LPopRequest, opts ...grpc.CallOption) (*PopResponse, error)
	RPop(ctx context.Context, in *RPopRequest, opts ...grpc.CallOption) (*PopResponse, error)
	LRange(ctx context.Context, in *LRangeRequest, opts ...grpc.CallOption) (*LRangeResponse, error)
	// Set----------------------------------<
	SAdd(ctx context.Context, in *SAddRequest, opts ...grpc.CallOption) (*SAddResponse, error)
	SRem(ctx context.Context, in *SRemRequest, opts ...grpc.CallOption) (*SRemResponse, error)
	SISMember(ctx context.Context, in *SISMemberRequest, opts ...grpc.CallOption) (*SISMemberResponse, error)
	SMembers(ctx context.Context, in *SMembersRequest, opts ...grpc.CallOption) (*SMembersResponse, error)
	SCard(ctx context.Context, in *SCardRequest, opts ...grpc.CallOption) (*SCardResponse, error)
}

type redisSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisSvrClient(cc grpc.ClientConnInterface) RedisSvrClient {
	return &redisSvrClient{cc}
}

func (c *redisSvrClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRsp)
	err := c.cc.Invoke(ctx, RedisSvr_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) BatchGet(ctx context.Context, in *BatchGetReq, opts ...grpc.CallOption) (*BatchGetRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetRsp)
	err := c.cc.Invoke(ctx, RedisSvr_BatchGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetRsp)
	err := c.cc.Invoke(ctx, RedisSvr_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) SetWithExpire(ctx context.Context, in *SetWithExpireReq, opts ...grpc.CallOption) (*SetWithExpireRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetWithExpireRsp)
	err := c.cc.Invoke(ctx, RedisSvr_SetWithExpire_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) BatchSet(ctx context.Context, in *BatchSetReq, opts ...grpc.CallOption) (*BatchSetRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchSetRsp)
	err := c.cc.Invoke(ctx, RedisSvr_BatchSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) SetIfNotExist(ctx context.Context, in *SetIfNotExistReq, opts ...grpc.CallOption) (*SetIfNotExistRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetIfNotExistRsp)
	err := c.cc.Invoke(ctx, RedisSvr_SetIfNotExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRsp)
	err := c.cc.Invoke(ctx, RedisSvr_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) LPush(ctx context.Context, in *LPushRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, RedisSvr_LPush_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) RPush(ctx context.Context, in *RPushRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, RedisSvr_RPush_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) LPop(ctx context.Context, in *LPopRequest, opts ...grpc.CallOption) (*PopResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PopResponse)
	err := c.cc.Invoke(ctx, RedisSvr_LPop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) RPop(ctx context.Context, in *RPopRequest, opts ...grpc.CallOption) (*PopResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PopResponse)
	err := c.cc.Invoke(ctx, RedisSvr_RPop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) LRange(ctx context.Context, in *LRangeRequest, opts ...grpc.CallOption) (*LRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LRangeResponse)
	err := c.cc.Invoke(ctx, RedisSvr_LRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) SAdd(ctx context.Context, in *SAddRequest, opts ...grpc.CallOption) (*SAddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SAddResponse)
	err := c.cc.Invoke(ctx, RedisSvr_SAdd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) SRem(ctx context.Context, in *SRemRequest, opts ...grpc.CallOption) (*SRemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SRemResponse)
	err := c.cc.Invoke(ctx, RedisSvr_SRem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) SISMember(ctx context.Context, in *SISMemberRequest, opts ...grpc.CallOption) (*SISMemberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SISMemberResponse)
	err := c.cc.Invoke(ctx, RedisSvr_SISMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) SMembers(ctx context.Context, in *SMembersRequest, opts ...grpc.CallOption) (*SMembersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SMembersResponse)
	err := c.cc.Invoke(ctx, RedisSvr_SMembers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisSvrClient) SCard(ctx context.Context, in *SCardRequest, opts ...grpc.CallOption) (*SCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCardResponse)
	err := c.cc.Invoke(ctx, RedisSvr_SCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedisSvrServer is the server API for RedisSvr service.
// All implementations must embed UnimplementedRedisSvrServer
// for forward compatibility
type RedisSvrServer interface {
	Get(context.Context, *GetReq) (*GetRsp, error)
	BatchGet(context.Context, *BatchGetReq) (*BatchGetRsp, error)
	Set(context.Context, *SetReq) (*SetRsp, error)
	SetWithExpire(context.Context, *SetWithExpireReq) (*SetWithExpireRsp, error)
	BatchSet(context.Context, *BatchSetReq) (*BatchSetRsp, error)
	SetIfNotExist(context.Context, *SetIfNotExistReq) (*SetIfNotExistRsp, error)
	Delete(context.Context, *DeleteReq) (*DeleteRsp, error)
	// List----------------------------------<
	LPush(context.Context, *LPushRequest) (*ListResponse, error)
	RPush(context.Context, *RPushRequest) (*ListResponse, error)
	LPop(context.Context, *LPopRequest) (*PopResponse, error)
	RPop(context.Context, *RPopRequest) (*PopResponse, error)
	LRange(context.Context, *LRangeRequest) (*LRangeResponse, error)
	// Set----------------------------------<
	SAdd(context.Context, *SAddRequest) (*SAddResponse, error)
	SRem(context.Context, *SRemRequest) (*SRemResponse, error)
	SISMember(context.Context, *SISMemberRequest) (*SISMemberResponse, error)
	SMembers(context.Context, *SMembersRequest) (*SMembersResponse, error)
	SCard(context.Context, *SCardRequest) (*SCardResponse, error)
	mustEmbedUnimplementedRedisSvrServer()
}

// UnimplementedRedisSvrServer must be embedded to have forward compatible implementations.
type UnimplementedRedisSvrServer struct {
}

func (UnimplementedRedisSvrServer) Get(context.Context, *GetReq) (*GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRedisSvrServer) BatchGet(context.Context, *BatchGetReq) (*BatchGetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (UnimplementedRedisSvrServer) Set(context.Context, *SetReq) (*SetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedRedisSvrServer) SetWithExpire(context.Context, *SetWithExpireReq) (*SetWithExpireRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWithExpire not implemented")
}
func (UnimplementedRedisSvrServer) BatchSet(context.Context, *BatchSetReq) (*BatchSetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchSet not implemented")
}
func (UnimplementedRedisSvrServer) SetIfNotExist(context.Context, *SetIfNotExistReq) (*SetIfNotExistRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIfNotExist not implemented")
}
func (UnimplementedRedisSvrServer) Delete(context.Context, *DeleteReq) (*DeleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRedisSvrServer) LPush(context.Context, *LPushRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LPush not implemented")
}
func (UnimplementedRedisSvrServer) RPush(context.Context, *RPushRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RPush not implemented")
}
func (UnimplementedRedisSvrServer) LPop(context.Context, *LPopRequest) (*PopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LPop not implemented")
}
func (UnimplementedRedisSvrServer) RPop(context.Context, *RPopRequest) (*PopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RPop not implemented")
}
func (UnimplementedRedisSvrServer) LRange(context.Context, *LRangeRequest) (*LRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LRange not implemented")
}
func (UnimplementedRedisSvrServer) SAdd(context.Context, *SAddRequest) (*SAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SAdd not implemented")
}
func (UnimplementedRedisSvrServer) SRem(context.Context, *SRemRequest) (*SRemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SRem not implemented")
}
func (UnimplementedRedisSvrServer) SISMember(context.Context, *SISMemberRequest) (*SISMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SISMember not implemented")
}
func (UnimplementedRedisSvrServer) SMembers(context.Context, *SMembersRequest) (*SMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SMembers not implemented")
}
func (UnimplementedRedisSvrServer) SCard(context.Context, *SCardRequest) (*SCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SCard not implemented")
}
func (UnimplementedRedisSvrServer) mustEmbedUnimplementedRedisSvrServer() {}

// UnsafeRedisSvrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedisSvrServer will
// result in compilation errors.
type UnsafeRedisSvrServer interface {
	mustEmbedUnimplementedRedisSvrServer()
}

func RegisterRedisSvrServer(s grpc.ServiceRegistrar, srv RedisSvrServer) {
	s.RegisterService(&RedisSvr_ServiceDesc, srv)
}

func _RedisSvr_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_BatchGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).BatchGet(ctx, req.(*BatchGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).Set(ctx, req.(*SetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_SetWithExpire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWithExpireReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).SetWithExpire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_SetWithExpire_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).SetWithExpire(ctx, req.(*SetWithExpireReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_BatchSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).BatchSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_BatchSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).BatchSet(ctx, req.(*BatchSetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_SetIfNotExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetIfNotExistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).SetIfNotExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_SetIfNotExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).SetIfNotExist(ctx, req.(*SetIfNotExistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_LPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LPushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).LPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_LPush_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).LPush(ctx, req.(*LPushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_RPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).RPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_RPush_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).RPush(ctx, req.(*RPushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_LPop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LPopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).LPop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_LPop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).LPop(ctx, req.(*LPopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_RPop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).RPop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_RPop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).RPop(ctx, req.(*RPopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_LRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).LRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_LRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).LRange(ctx, req.(*LRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_SAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).SAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_SAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).SAdd(ctx, req.(*SAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_SRem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SRemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).SRem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_SRem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).SRem(ctx, req.(*SRemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_SISMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SISMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).SISMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_SISMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).SISMember(ctx, req.(*SISMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_SMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).SMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_SMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).SMembers(ctx, req.(*SMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisSvr_SCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisSvrServer).SCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisSvr_SCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisSvrServer).SCard(ctx, req.(*SCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RedisSvr_ServiceDesc is the grpc.ServiceDesc for RedisSvr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisSvr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "redis_svr",
	HandlerType: (*RedisSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RedisSvr_Get_Handler,
		},
		{
			MethodName: "BatchGet",
			Handler:    _RedisSvr_BatchGet_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _RedisSvr_Set_Handler,
		},
		{
			MethodName: "SetWithExpire",
			Handler:    _RedisSvr_SetWithExpire_Handler,
		},
		{
			MethodName: "BatchSet",
			Handler:    _RedisSvr_BatchSet_Handler,
		},
		{
			MethodName: "SetIfNotExist",
			Handler:    _RedisSvr_SetIfNotExist_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RedisSvr_Delete_Handler,
		},
		{
			MethodName: "LPush",
			Handler:    _RedisSvr_LPush_Handler,
		},
		{
			MethodName: "RPush",
			Handler:    _RedisSvr_RPush_Handler,
		},
		{
			MethodName: "LPop",
			Handler:    _RedisSvr_LPop_Handler,
		},
		{
			MethodName: "RPop",
			Handler:    _RedisSvr_RPop_Handler,
		},
		{
			MethodName: "LRange",
			Handler:    _RedisSvr_LRange_Handler,
		},
		{
			MethodName: "SAdd",
			Handler:    _RedisSvr_SAdd_Handler,
		},
		{
			MethodName: "SRem",
			Handler:    _RedisSvr_SRem_Handler,
		},
		{
			MethodName: "SISMember",
			Handler:    _RedisSvr_SISMember_Handler,
		},
		{
			MethodName: "SMembers",
			Handler:    _RedisSvr_SMembers_Handler,
		},
		{
			MethodName: "SCard",
			Handler:    _RedisSvr_SCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redis_svr.proto",
}
