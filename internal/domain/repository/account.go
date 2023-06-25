//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE
package repository

import (
	"context"

	"github.com/taniko/rin/internal/domain/model/account"
)

type Account interface {
	Create(ctx context.Context, in account.Account) (*account.Account, error)
	FindByID(ctx context.Context, ID account.ID) (*account.Account, error)
	FindByNameOrEmail(ctx context.Context, name string, email account.Email) (*account.Account, error)
	FindByEmail(ctx context.Context, email account.Email) (*account.Account, error)
}
