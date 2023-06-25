//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE
package repository

import (
	"context"

	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/model/community"
)

type Community interface {
	Create(ctx context.Context, community community.Community, user account.ID) error

	FindByID(ctx context.Context, id community.ID) (*community.Community, error)

	FindByName(ctx context.Context, name community.Name) (*community.Community, error)

	IssueInvitation(ctx context.Context, id community.ID, user account.ID) (*community.Invitation, error)

	GetUserRole(ctx context.Context, id community.ID, user account.ID) (community.Role, error)

	// FindInvitation Invitationを取得
	FindInvitation(ctx context.Context, invitation community.Invitation) (*community.Invitation, error)

	// Join communityにuserを追加
	Join(ctx context.Context, id community.ID, accountID account.ID, role community.Role) error
}
