//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE
package repository

import (
	"context"

	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/model/community"
)

type Community interface {
	Create(ctx context.Context, community community.Community, user account.ID) error
	FindByName(ctx context.Context, name community.Name) (*community.Community, error)
}
