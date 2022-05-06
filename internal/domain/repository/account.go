package repository

import (
	"context"

	"github.com/taniko/rin/internal/domain/model"
)

type Account interface {
	Create(ctx context.Context, account model.Account) (*model.Account, error)
	FindByNameOrEmail(ctx context.Context, name, email string) (*model.Account, error)
}
