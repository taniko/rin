package handler

import (
	"context"

	"github.com/bufbuild/connect-go"
	v1 "github.com/taniko/rin/internal/gen/taniko/rin/community/v1"
	"github.com/taniko/rin/internal/gen/taniko/rin/community/v1/communityv1connect"
)

type communityServer struct {
}

func NewCommunityServer() communityv1connect.CommunityServiceHandler {
	return &communityServer{}
}

func (c communityServer) CreateCommunity(ctx context.Context, req *connect.Request[v1.CreateCommunityRequest]) (*connect.Response[v1.CreateCommunityResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c communityServer) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c communityServer) CreateInviteToken(ctx context.Context, req *connect.Request[v1.CreateInviteTokenRequest]) (*connect.Response[v1.CreateInviteTokenResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c communityServer) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c communityServer) ChangeUserRole(ctx context.Context, req *connect.Request[v1.ChangeUserRoleRequest]) (*connect.Response[v1.ChangeUserRoleResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c communityServer) ListUsers(ctx context.Context, req *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c communityServer) ListJoinCommunities(ctx context.Context, req *connect.Request[v1.ListJoinCommunitiesRequest]) (*connect.Response[v1.ListJoinCommunitiesResponse], error) {
	//TODO implement me
	panic("implement me")
}
