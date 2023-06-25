package handler

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/dto"
	v1 "github.com/taniko/rin/internal/gen/taniko/rin/account/v1"
	"github.com/taniko/rin/internal/gen/taniko/rin/account/v1/accountv1connect"
	"github.com/taniko/rin/internal/usecase"
)

type accountHandler struct {
	u usecase.Account
}

func NewAccount(u usecase.Account) accountv1connect.AccountServiceHandler {
	return &accountHandler{
		u: u,
	}
}

func (h accountHandler) CreateAccount(ctx context.Context, req *connect.Request[v1.CreateAccountRequest]) (*connect.Response[v1.CreateAccountResponse], error) {
	token, err := h.u.Create(ctx, dto.CreateAccount{
		Name:                 req.Msg.GetName(),
		DisplayName:          req.Msg.GetDisplayName(),
		Email:                account.Email(req.Msg.GetEmail()),
		Password:             account.RawPassword(req.Msg.GetPassword()),
		PasswordConfirmation: account.RawPassword(req.Msg.GetPasswordConfirmation()),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&v1.CreateAccountResponse{
		Token: string(token.Value()),
	}), nil
}

func (h accountHandler) UpdateAccount(ctx context.Context, req *connect.Request[v1.UpdateAccountRequest]) (*connect.Response[v1.UpdateAccountResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h accountHandler) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	token, err := h.u.Login(ctx, account.Email(req.Msg.GetEmail()), account.RawPassword(req.Msg.GetPassword()))
	if err != nil {
		return nil, err
	} else if token.IsNull() {
		return nil, connect.NewError(connect.CodeInternal, errors.New("issue token error"))
	}
	return connect.NewResponse(&v1.LoginResponse{
		Token: string(token.Value()),
	}), nil
}

func (h accountHandler) VerifyToken(ctx context.Context, req *connect.Request[v1.VerifyTokenRequest]) (*connect.Response[v1.VerifyTokenResponse], error) {
	//TODO implement me
	panic("implement me")
}
