package handler

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/taniko/rin/internal/gen/taniko/rin/message/v1"
	"github.com/taniko/rin/internal/gen/taniko/rin/message/v1/messagev1connect"
	"github.com/taniko/sumire"
)

type messageServer struct {
	logger *sumire.Sumire
}

func (m messageServer) CreateMessage(ctx context.Context, c *connect_go.Request[v1.CreateMessageRequest]) (*connect_go.Response[v1.CreateMessageResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (m messageServer) DeleteMessage(ctx context.Context, c *connect_go.Request[v1.DeleteMessageRequest]) (*connect_go.Response[v1.DeleteMessageResponse], error) {
	//TODO implement me
	panic("implement me")
}

func NewMessageServer(logger *sumire.Sumire) messagev1connect.MessageServiceHandler {
	return &messageServer{
		logger: logger,
	}
}
