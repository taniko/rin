package handler

import (
	"context"

	"github.com/bufbuild/connect-go"
	v1 "github.com/taniko/rin/internal/gen/taniko/rin/channel/v1"
	"github.com/taniko/rin/internal/gen/taniko/rin/channel/v1/channelv1connect"
	"github.com/taniko/sumire"
)

type channelHandler struct {
	logger *sumire.Sumire
}

func NewChannel(logger *sumire.Sumire) channelv1connect.ChannelServiceHandler {
	return &channelHandler{
		logger: logger,
	}
}

func (c channelHandler) CreateChannel(ctx context.Context, c2 *connect.Request[v1.CreateChannelRequest]) (*connect.Response[v1.CreateChannelResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c channelHandler) UpdateChannel(ctx context.Context, c2 *connect.Request[v1.UpdateChannelRequest]) (*connect.Response[v1.UpdateChannelResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c channelHandler) ArchiveChannel(ctx context.Context, c2 *connect.Request[v1.ArchiveChannelRequest]) (*connect.Response[v1.ArchiveChannelResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c channelHandler) DeleteChannel(ctx context.Context, c2 *connect.Request[v1.DeleteChannelRequest]) (*connect.Response[v1.DeleteChannelResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c channelHandler) ListUsers(ctx context.Context, c2 *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c channelHandler) GetUserRole(ctx context.Context, c2 *connect.Request[v1.GetUserRoleRequest]) (*connect.Response[v1.GetUserRoleResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (c channelHandler) ListRoleUsers(ctx context.Context, c2 *connect.Request[v1.ListRoleUsersRequest]) (*connect.Response[v1.ListRoleUsersResponse], error) {
	//TODO implement me
	panic("implement me")
}
