package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/model/community"
	"github.com/taniko/rin/internal/domain/repository"
	mock "github.com/taniko/rin/internal/domain/repository/mock"
)

func TestCommunityUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userID := account.ID(uuid.NewString())
	communityID := community.ID(uuid.NewString())
	communityName := community.Name(uuid.NewString())
	communityRepo := mock.NewMockCommunity(ctrl)
	communityRepo.EXPECT().FindByName(gomock.Any(), communityName).Return(nil, repository.ErrNotFound).Times(1)
	communityRepo.EXPECT().FindByName(gomock.Any(), communityName).Return(&community.Community{
		ID:        communityID,
		Name:      communityName,
		CreatedAt: time.Now(),
	}, nil).AnyTimes()
	communityRepo.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	type input struct {
		name        community.Name
		displayName community.DisplayName
		access      community.Access
		user        account.ID
	}
	tests := []struct {
		name    string
		in      input
		wantErr bool
	}{
		{
			name: "ok",
			in: input{
				name:        communityName,
				displayName: "community1",
				access:      community.Private,
				user:        userID,
			},
			wantErr: false,
		}, {
			name: "duplicate",
			in: input{
				name:        communityName,
				displayName: "community2",
				access:      community.Private,
				user:        userID,
			},
			wantErr: true,
		},
	}
	u := NewCommunity(communityRepo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := u.Create(context.Background(), tt.in.name, tt.in.displayName, tt.in.access, tt.in.user)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
