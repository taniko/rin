package server

import (
	channel "github.com/taniko/rin/internal/pb/taniko/rin/channel/v1"
	"github.com/taniko/sumire"
)

type channelServer struct {
	channel.UnimplementedChannelServiceServer

	logger *sumire.Sumire
}

func NewChannelServer(logger *sumire.Sumire) channel.ChannelServiceServer {
	return &channelServer{
		logger: logger,
	}
}
