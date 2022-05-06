package server

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taniko/rin/internal/domain/model"
	"github.com/taniko/rin/internal/domain/repository"
	"github.com/taniko/rin/internal/dto"
	account "github.com/taniko/rin/internal/pb/taniko/rin/account/v1"
	"github.com/taniko/rin/internal/usecase"
	"github.com/taniko/sumire"
)

var _ repository.Account = (*inmemoryAccountDatabase)(nil)

type inmemoryAccountDatabase struct {
	data map[string]model.Account
}

func (i inmemoryAccountDatabase) FindByNameOrEmail(_ context.Context, _, _ string) (*model.Account, error) {
	return nil, nil
}

func newInmemoryAccountDatabase() inmemoryAccountDatabase {
	return inmemoryAccountDatabase{
		data: map[string]model.Account{},
	}
}

func (i inmemoryAccountDatabase) Create(_ context.Context, account model.Account) (*model.Account, error) {
	if _, ok := i.data[account.ID]; ok {
		return nil, errors.New("exists id")
	}
	i.data[account.ID] = account
	return &account, nil
}

func TestAccountServer_CreateAccount(t *testing.T) {
	logger := newTestLogger()
	s := NewAccountServer(logger, usecase.NewAccountUsecase(logger, newInmemoryAccountDatabase()))

	type successTest struct {
		account dto.CreateAccount
	}
	successTests := []successTest{
		{
			account: dto.CreateAccount{
				Name:                 "name_1",
				DisplayName:          "name 1",
				Email:                "foo@example.com",
				Password:             "password123456789",
				PasswordConfirmation: "password123456789",
			},
		}, {
			account: dto.CreateAccount{
				Name:                 "name_2",
				DisplayName:          "name 2",
				Email:                "bar@example.com",
				Password:             "password314159265",
				PasswordConfirmation: "password314159265",
			},
		},
	}
	for _, tt := range successTests {
		t.Run("success test", func(t *testing.T) {
			ctx := context.Background()
			_, err := s.CreateAccount(ctx, &account.CreateAccountRequest{
				Name:                 tt.account.Name,
				DisplayName:          tt.account.DisplayName,
				Email:                tt.account.Email,
				Password:             tt.account.Password,
				PasswordConfirmation: tt.account.PasswordConfirmation,
			})
			assert.NoError(t, err)
		})
	}

	type failTest struct {
		name    string
		account dto.CreateAccount
	}
	failTests := []failTest{
		{
			name: "name is short",
			account: dto.CreateAccount{
				Name:                 "a",
				DisplayName:          "name",
				Email:                "a@example.com",
				Password:             "password123456789",
				PasswordConfirmation: "password123456789",
			},
		}, {
			name: "display name is empty",
			account: dto.CreateAccount{
				Name:                 "abc",
				DisplayName:          "",
				Email:                "abc@example.com",
				Password:             "password123456789",
				PasswordConfirmation: "password123456789",
			},
		}, {
			name: "invalid email",
			account: dto.CreateAccount{
				Name:                 "abc",
				DisplayName:          "name",
				Email:                "example.com",
				Password:             "password123456789",
				PasswordConfirmation: "password123456789",
			},
		}, {
			name: "password is short",
			account: dto.CreateAccount{
				Name:                 "abc",
				DisplayName:          "name",
				Email:                "abc@example.com",
				Password:             "password",
				PasswordConfirmation: "password",
			},
		}, {
			name: "password is not match",
			account: dto.CreateAccount{
				Name:                 "abc",
				DisplayName:          "name",
				Email:                "abc@example.com",
				Password:             "password123456789",
				PasswordConfirmation: "password12345678",
			},
		},
	}
	for _, tt := range failTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := s.CreateAccount(ctx, &account.CreateAccountRequest{
				Name:                 tt.account.Name,
				DisplayName:          tt.account.DisplayName,
				Email:                tt.account.Email,
				Password:             tt.account.Password,
				PasswordConfirmation: tt.account.PasswordConfirmation,
			})
			assert.Error(t, err)
		})
	}
}

func newTestLogger() *sumire.Sumire {
	return sumire.NewLogger("test")
}
