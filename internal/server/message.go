package server

import (
	message "github.com/taniko/rin/internal/pb/taniko/rin/message/v1"
	"github.com/taniko/sumire"
)

type messageServer struct {
	message.UnimplementedMessageServiceServer

	logger *sumire.Sumire
}

func NewMessageServer(logger *sumire.Sumire) message.MessageServiceServer {
	return &messageServer{
		logger: logger,
	}
}
