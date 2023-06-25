package handler

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/repository"
	mock "github.com/taniko/rin/internal/domain/repository/mock"
	"github.com/taniko/rin/internal/dto"
	accountv1 "github.com/taniko/rin/internal/gen/taniko/rin/account/v1"
	"github.com/taniko/rin/internal/usecase"
)

func TestAccountServer_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockAccount(ctrl)
	s := NewAccount(usecase.NewAccountUsecase(repo, []byte("key")))

	repo.EXPECT().FindByNameOrEmail(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, repository.ErrNotFound).AnyTimes()
	repo.EXPECT().Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, account account.Account) (*account.Account, error) {
			return &account, nil
		}).AnyTimes()

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
			_, err := s.CreateAccount(ctx, connect.NewRequest(&accountv1.CreateAccountRequest{
				Name:                 tt.account.Name,
				DisplayName:          tt.account.DisplayName,
				Email:                string(tt.account.Email),
				Password:             string(tt.account.Password),
				PasswordConfirmation: string(tt.account.PasswordConfirmation),
			}))
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
			_, err := s.CreateAccount(ctx, connect.NewRequest(&accountv1.CreateAccountRequest{
				Name:                 tt.account.Name,
				DisplayName:          tt.account.DisplayName,
				Email:                string(tt.account.Email),
				Password:             string(tt.account.Password),
				PasswordConfirmation: string(tt.account.PasswordConfirmation),
			}))
			assert.Error(t, err)
		})
	}
}
