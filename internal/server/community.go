package server

import (
	community "github.com/taniko/rin/internal/pb/taniko/rin/community/v1"
	"github.com/taniko/sumire"
)

type communityServer struct {
	community.UnimplementedCommunityServiceServer

	logger *sumire.Sumire
}

func NewCommunityServer(logger *sumire.Sumire) community.CommunityServiceServer {
	return &communityServer{
		logger: logger,
	}
}
