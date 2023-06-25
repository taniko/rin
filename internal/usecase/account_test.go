package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/repository"
	mock "github.com/taniko/rin/internal/domain/repository/mock"
	"github.com/taniko/rin/internal/dto"
)

func TestAccountUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	password := uuid.NewString()
	tests := []struct {
		name    string
		in      dto.CreateAccount
		wantErr bool
	}{
		{
			in: dto.CreateAccount{
				Name:                 "foo",
				DisplayName:          "foo",
				Email:                "example@example.com",
				Password:             account.RawPassword(password),
				PasswordConfirmation: account.RawPassword(password),
			},
			wantErr: false,
		}, {
			name: "duplicate name",
			in: dto.CreateAccount{
				Name:                 "foo",
				DisplayName:          uuid.NewString(),
				Email:                account.Email(fmt.Sprintf("example-{%s}@example.com", uuid.NewString())),
				Password:             account.RawPassword(password),
				PasswordConfirmation: account.RawPassword(password),
			},
			wantErr: true,
		}, {
			name: "duplicate email",
			in: dto.CreateAccount{
				Name:                 fmt.Sprintf("foo-%s", uuid.NewString()),
				DisplayName:          uuid.NewString(),
				Email:                "example@example.com",
				Password:             account.RawPassword(password),
				PasswordConfirmation: account.RawPassword(password),
			},
			wantErr: true,
		},
	}
	var accounts []*account.Account
	repo := mock.NewMockAccount(ctrl)
	repo.EXPECT().Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, in account.Account) (*account.Account, error) {
			for i := range accounts {
				if accounts[i].ID == in.ID {
					return nil, errors.New("duplicate id")
				}
			}
			accounts = append(accounts, &in)
			return &in, nil
		}).MinTimes(1)

	repo.EXPECT().FindByNameOrEmail(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, name string, email account.Email) (*account.Account, error) {
			for i := range accounts {
				if accounts[i].Name == name || accounts[i].Email == email {
					return accounts[i], nil
				}
			}
			return nil, repository.ErrNotFound
		}).MinTimes(1)

	u := NewAccountUsecase(repo, []byte("key"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := u.Create(context.Background(), tt.in)
			if tt.wantErr {
				assert.Error(t, err)
				assert.True(t, token.IsNull())
			} else {
				assert.NoError(t, err)
				assert.False(t, token.IsNull())
			}
		})
	}
}

func TestAccountUsecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	password := account.RawPassword(uuid.NewString())
	user := struct {
		email    account.Email
		password account.RawPassword
	}{
		email:    "example@example.com",
		password: password,
	}
	tests := []struct {
		name    string
		in      dto.LoginRequest
		wantErr bool
	}{
		{
			name: "success",
			in: dto.LoginRequest{
				Email:    user.email,
				Password: user.password,
			},
			wantErr: false,
		}, {
			name: "incorrect email",
			in: dto.LoginRequest{
				Email:    account.Email(fmt.Sprintf("none_%s", user.email)),
				Password: user.password,
			},
			wantErr: true,
		}, {
			name: "incorrect password",
			in: dto.LoginRequest{
				Email:    user.email,
				Password: account.RawPassword(uuid.NewString()),
			},
			wantErr: true,
		},
	}
	var accounts []*account.Account
	repo := mock.NewMockAccount(ctrl)
	repo.EXPECT().Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, in account.Account) (*account.Account, error) {
			for i := range accounts {
				if accounts[i].ID == in.ID {
					return nil, errors.New("duplicate id")
				}
			}
			accounts = append(accounts, &in)
			return &in, nil
		}).Times(1)
	repo.EXPECT().FindByNameOrEmail(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, name string, email account.Email) (*account.Account, error) {
			for i := range accounts {
				if accounts[i].Name == name || accounts[i].Email == email {
					return accounts[i], nil
				}
			}
			return nil, repository.ErrNotFound
		}).MinTimes(1)
	repo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, email account.Email) (*account.Account, error) {
			for i := range accounts {
				if accounts[i].Email == email {
					return accounts[i], nil
				}
			}
			return nil, repository.ErrNotFound
		}).MinTimes(1)
	u := NewAccountUsecase(repo, []byte("key"))
	_, err := u.Create(context.Background(), dto.CreateAccount{
		Name:                 "login_check",
		DisplayName:          "login check",
		Email:                user.email,
		Password:             user.password,
		PasswordConfirmation: user.password,
	})
	assert.NoError(t, err, "failed to create account: %w", err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := u.Login(context.Background(), tt.in.Email, tt.in.Password)
			if tt.wantErr {
				assert.Error(t, err)
				assert.True(t, token.IsNull())
				return
			}
			assert.NoError(t, err)
			assert.False(t, token.IsNull())
		})
	}
}
