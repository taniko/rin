// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: taniko/rin/channel/v1/channel.proto

package channelv1

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

// ChannelServiceClient is the client API for ChannelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelServiceClient interface {
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*UpdateChannelResponse, error)
	ArchiveChannel(ctx context.Context, in *ArchiveChannelRequest, opts ...grpc.CallOption) (*ArchiveChannelResponse, error)
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error)
	ListRoleUsers(ctx context.Context, in *ListRoleUsersRequest, opts ...grpc.CallOption) (*ListRoleUsersResponse, error)
}

type channelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelServiceClient(cc grpc.ClientConnInterface) ChannelServiceClient {
	return &channelServiceClient{cc}
}

func (c *channelServiceClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.channel.v1.ChannelService/CreateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*UpdateChannelResponse, error) {
	out := new(UpdateChannelResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.channel.v1.ChannelService/UpdateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) ArchiveChannel(ctx context.Context, in *ArchiveChannelRequest, opts ...grpc.CallOption) (*ArchiveChannelResponse, error) {
	out := new(ArchiveChannelResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.channel.v1.ChannelService/ArchiveChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error) {
	out := new(DeleteChannelResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.channel.v1.ChannelService/DeleteChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.channel.v1.ChannelService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error) {
	out := new(GetUserRoleResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.channel.v1.ChannelService/GetUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) ListRoleUsers(ctx context.Context, in *ListRoleUsersRequest, opts ...grpc.CallOption) (*ListRoleUsersResponse, error) {
	out := new(ListRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.channel.v1.ChannelService/ListRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServiceServer is the server API for ChannelService service.
// All implementations must embed UnimplementedChannelServiceServer
// for forward compatibility
type ChannelServiceServer interface {
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelResponse, error)
	ArchiveChannel(context.Context, *ArchiveChannelRequest) (*ArchiveChannelResponse, error)
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error)
	ListRoleUsers(context.Context, *ListRoleUsersRequest) (*ListRoleUsersResponse, error)
	mustEmbedUnimplementedChannelServiceServer()
}

// UnimplementedChannelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChannelServiceServer struct {
}

func (UnimplementedChannelServiceServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedChannelServiceServer) UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannel not implemented")
}
func (UnimplementedChannelServiceServer) ArchiveChannel(context.Context, *ArchiveChannelRequest) (*ArchiveChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveChannel not implemented")
}
func (UnimplementedChannelServiceServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedChannelServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedChannelServiceServer) GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRole not implemented")
}
func (UnimplementedChannelServiceServer) ListRoleUsers(context.Context, *ListRoleUsersRequest) (*ListRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoleUsers not implemented")
}
func (UnimplementedChannelServiceServer) mustEmbedUnimplementedChannelServiceServer() {}

// UnsafeChannelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelServiceServer will
// result in compilation errors.
type UnsafeChannelServiceServer interface {
	mustEmbedUnimplementedChannelServiceServer()
}

func RegisterChannelServiceServer(s grpc.ServiceRegistrar, srv ChannelServiceServer) {
	s.RegisterService(&ChannelService_ServiceDesc, srv)
}

func _ChannelService_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.channel.v1.ChannelService/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_UpdateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).UpdateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.channel.v1.ChannelService/UpdateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).UpdateChannel(ctx, req.(*UpdateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_ArchiveChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).ArchiveChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.channel.v1.ChannelService/ArchiveChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).ArchiveChannel(ctx, req.(*ArchiveChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.channel.v1.ChannelService/DeleteChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.channel.v1.ChannelService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.channel.v1.ChannelService/GetUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_ListRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).ListRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.channel.v1.ChannelService/ListRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).ListRoleUsers(ctx, req.(*ListRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelService_ServiceDesc is the grpc.ServiceDesc for ChannelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taniko.rin.channel.v1.ChannelService",
	HandlerType: (*ChannelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannel",
			Handler:    _ChannelService_CreateChannel_Handler,
		},
		{
			MethodName: "UpdateChannel",
			Handler:    _ChannelService_UpdateChannel_Handler,
		},
		{
			MethodName: "ArchiveChannel",
			Handler:    _ChannelService_ArchiveChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _ChannelService_DeleteChannel_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _ChannelService_ListUsers_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _ChannelService_GetUserRole_Handler,
		},
		{
			MethodName: "ListRoleUsers",
			Handler:    _ChannelService_ListRoleUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taniko/rin/channel/v1/channel.proto",
}
