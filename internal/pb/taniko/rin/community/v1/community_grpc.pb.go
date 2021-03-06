// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: taniko/rin/community/v1/community.proto

package communityv1

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

// CommunityServiceClient is the client API for CommunityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunityServiceClient interface {
	CreateCommunity(ctx context.Context, in *CreateCommunityRequest, opts ...grpc.CallOption) (*CreateCommunityResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CreateInviteToken(ctx context.Context, in *CreateInviteTokenRequest, opts ...grpc.CallOption) (*CreateInviteTokenResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ChangeUserRole(ctx context.Context, in *ChangeUserRoleRequest, opts ...grpc.CallOption) (*ChangeUserRoleResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	ListJoinCommunities(ctx context.Context, in *ListJoinCommunitiesRequest, opts ...grpc.CallOption) (*ListJoinCommunitiesResponse, error)
}

type communityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunityServiceClient(cc grpc.ClientConnInterface) CommunityServiceClient {
	return &communityServiceClient{cc}
}

func (c *communityServiceClient) CreateCommunity(ctx context.Context, in *CreateCommunityRequest, opts ...grpc.CallOption) (*CreateCommunityResponse, error) {
	out := new(CreateCommunityResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.community.v1.CommunityService/CreateCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.community.v1.CommunityService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) CreateInviteToken(ctx context.Context, in *CreateInviteTokenRequest, opts ...grpc.CallOption) (*CreateInviteTokenResponse, error) {
	out := new(CreateInviteTokenResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.community.v1.CommunityService/CreateInviteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.community.v1.CommunityService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) ChangeUserRole(ctx context.Context, in *ChangeUserRoleRequest, opts ...grpc.CallOption) (*ChangeUserRoleResponse, error) {
	out := new(ChangeUserRoleResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.community.v1.CommunityService/ChangeUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.community.v1.CommunityService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) ListJoinCommunities(ctx context.Context, in *ListJoinCommunitiesRequest, opts ...grpc.CallOption) (*ListJoinCommunitiesResponse, error) {
	out := new(ListJoinCommunitiesResponse)
	err := c.cc.Invoke(ctx, "/taniko.rin.community.v1.CommunityService/ListJoinCommunities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunityServiceServer is the server API for CommunityService service.
// All implementations must embed UnimplementedCommunityServiceServer
// for forward compatibility
type CommunityServiceServer interface {
	CreateCommunity(context.Context, *CreateCommunityRequest) (*CreateCommunityResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CreateInviteToken(context.Context, *CreateInviteTokenRequest) (*CreateInviteTokenResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ChangeUserRole(context.Context, *ChangeUserRoleRequest) (*ChangeUserRoleResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	ListJoinCommunities(context.Context, *ListJoinCommunitiesRequest) (*ListJoinCommunitiesResponse, error)
	mustEmbedUnimplementedCommunityServiceServer()
}

// UnimplementedCommunityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunityServiceServer struct {
}

func (UnimplementedCommunityServiceServer) CreateCommunity(context.Context, *CreateCommunityRequest) (*CreateCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunity not implemented")
}
func (UnimplementedCommunityServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedCommunityServiceServer) CreateInviteToken(context.Context, *CreateInviteTokenRequest) (*CreateInviteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInviteToken not implemented")
}
func (UnimplementedCommunityServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedCommunityServiceServer) ChangeUserRole(context.Context, *ChangeUserRoleRequest) (*ChangeUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserRole not implemented")
}
func (UnimplementedCommunityServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedCommunityServiceServer) ListJoinCommunities(context.Context, *ListJoinCommunitiesRequest) (*ListJoinCommunitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJoinCommunities not implemented")
}
func (UnimplementedCommunityServiceServer) mustEmbedUnimplementedCommunityServiceServer() {}

// UnsafeCommunityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunityServiceServer will
// result in compilation errors.
type UnsafeCommunityServiceServer interface {
	mustEmbedUnimplementedCommunityServiceServer()
}

func RegisterCommunityServiceServer(s grpc.ServiceRegistrar, srv CommunityServiceServer) {
	s.RegisterService(&CommunityService_ServiceDesc, srv)
}

func _CommunityService_CreateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.community.v1.CommunityService/CreateCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreateCommunity(ctx, req.(*CreateCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.community.v1.CommunityService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_CreateInviteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInviteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreateInviteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.community.v1.CommunityService/CreateInviteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreateInviteToken(ctx, req.(*CreateInviteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.community.v1.CommunityService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_ChangeUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).ChangeUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.community.v1.CommunityService/ChangeUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).ChangeUserRole(ctx, req.(*ChangeUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.community.v1.CommunityService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_ListJoinCommunities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJoinCommunitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).ListJoinCommunities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taniko.rin.community.v1.CommunityService/ListJoinCommunities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).ListJoinCommunities(ctx, req.(*ListJoinCommunitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunityService_ServiceDesc is the grpc.ServiceDesc for CommunityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taniko.rin.community.v1.CommunityService",
	HandlerType: (*CommunityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommunity",
			Handler:    _CommunityService_CreateCommunity_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _CommunityService_CreateUser_Handler,
		},
		{
			MethodName: "CreateInviteToken",
			Handler:    _CommunityService_CreateInviteToken_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _CommunityService_GetUser_Handler,
		},
		{
			MethodName: "ChangeUserRole",
			Handler:    _CommunityService_ChangeUserRole_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _CommunityService_ListUsers_Handler,
		},
		{
			MethodName: "ListJoinCommunities",
			Handler:    _CommunityService_ListJoinCommunities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taniko/rin/community/v1/community.proto",
}
