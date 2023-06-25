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

var (
	ErrAlreadyExistCommunity = errors.New("already exists")
	ErrPermissionDeny        = errors.New("permission deny")
)

type Community interface {
	Create(ctx context.Context, name community.Name, displayName community.DisplayName, access community.Access, user account.ID) error
	IssueInvitation(ctx context.Context, user account.ID, community community.ID) (*community.Invitation, error)
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

func (c communityUsecase) IssueInvitation(ctx context.Context, userID account.ID, communityID community.ID) (*community.Invitation, error) {
	role, err := c.community.GetUserRole(ctx, communityID, userID)
	if err != nil {
		return nil, xerrors.Errorf("get user role: %w", err)
	} else if !role.CanIssueInvitation() {
		return nil, ErrPermissionDeny
	}
	invitation, err := c.community.IssueInvitation(ctx, communityID, userID)
	if err != nil {
		return nil, xerrors.Errorf("issue invitation: %w", err)
	}
	return invitation, nil
}
