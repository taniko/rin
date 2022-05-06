package server

import (
	"context"

	"github.com/taniko/rin/internal/dto"
	account "github.com/taniko/rin/internal/pb/taniko/rin/account/v1"
	"github.com/taniko/rin/internal/usecase"
	"github.com/taniko/sumire"
)

type accountServer struct {
	account.UnimplementedAccountServiceServer

	logger  *sumire.Sumire
	usecase usecase.Account
}

func NewAccountServer(logger *sumire.Sumire, usecase usecase.Account) account.AccountServiceServer {
	return &accountServer{
		logger:  logger,
		usecase: usecase,
	}
}

func (s accountServer) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	err := s.usecase.Create(ctx, dto.CreateAccount{
		Name:                 req.GetName(),
		DisplayName:          req.GetDisplayName(),
		Email:                req.GetEmail(),
		Password:             req.GetPassword(),
		PasswordConfirmation: req.GetPasswordConfirmation(),
	})
	if err != nil {
		return nil, err
	}
	return &account.CreateAccountResponse{}, nil
}
