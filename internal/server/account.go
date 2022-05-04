package server

import (
	account "github.com/taniko/rin/internal/pb/taniko/rin/account/v1"
	"github.com/taniko/sumire"
)

type accountServer struct {
	account.UnimplementedAccountServiceServer

	logger *sumire.Sumire
}

func NewAccountServer(logger *sumire.Sumire) account.AccountServiceServer {
	return &accountServer{
		logger: logger,
	}
}
