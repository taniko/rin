package community

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/model/community"
	"github.com/taniko/rin/internal/domain/repository"
)

type dao struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) repository.Community {
	return &dao{
		db: db,
	}
}

func (d dao) Create(ctx context.Context, community community.Community, user account.ID) error {
	//TODO implement me
	panic("implement me")
}

func (d dao) FindByID(ctx context.Context, id community.ID) (*community.Community, error) {
	//TODO implement me
	panic("implement me")
}

func (d dao) FindByName(ctx context.Context, name community.Name) (*community.Community, error) {
	//TODO implement me
	panic("implement me")
}

func (d dao) IssueInvitation(ctx context.Context, id community.ID, user account.ID) (*community.Invitation, error) {
	//TODO implement me
	panic("implement me")
}

func (d dao) GetUserRole(ctx context.Context, id community.ID, user account.ID) (community.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (d dao) FindInvitation(ctx context.Context, invitation community.Invitation) (*community.Invitation, error) {
	//TODO implement me
	panic("implement me")
}

func (d dao) Join(ctx context.Context, id community.ID, accountID account.ID, role community.Role) error {
	//TODO implement me
	panic("implement me")
}
