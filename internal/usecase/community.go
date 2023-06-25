package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/model/community"
	"github.com/taniko/rin/internal/domain/repository"
	"golang.org/x/xerrors"
)

var ErrAlreadyExistCommunity = errors.New("already exists")

type Community interface {
	Create(ctx context.Context, name community.Name, displayName community.DisplayName, access community.Access, user account.ID) error
}

type communityUsecase struct {
	community repository.Community
}

func NewCommunity(community repository.Community) Community {
	return communityUsecase{
		community: community,
	}
}

func (c communityUsecase) Create(ctx context.Context, name community.Name, displayName community.DisplayName, access community.Access, user account.ID) error {
	if v, err := c.community.FindByName(ctx, name); err != nil && !errors.Is(err, repository.ErrNotFound) {
		return xerrors.Errorf("find community by name: %w", err)
	} else if v != nil {
		return ErrAlreadyExistCommunity
	}
	err := c.community.Create(ctx, community.Community{
		ID:          community.ID(uuid.NewString()),
		Name:        name,
		DisplayName: displayName,
		Access:      access,
	}, user)
	if err != nil {
		return xerrors.Errorf("create community: %w", err)
	}
	return nil
}
