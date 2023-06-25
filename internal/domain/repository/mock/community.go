// Code generated by MockGen. DO NOT EDIT.
// Source: community.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	account "github.com/taniko/rin/internal/domain/model/account"
	community "github.com/taniko/rin/internal/domain/model/community"
)

// MockCommunity is a mock of Community interface.
type MockCommunity struct {
	ctrl     *gomock.Controller
	recorder *MockCommunityMockRecorder
}

// MockCommunityMockRecorder is the mock recorder for MockCommunity.
type MockCommunityMockRecorder struct {
	mock *MockCommunity
}

// NewMockCommunity creates a new mock instance.
func NewMockCommunity(ctrl *gomock.Controller) *MockCommunity {
	mock := &MockCommunity{ctrl: ctrl}
	mock.recorder = &MockCommunityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommunity) EXPECT() *MockCommunityMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCommunity) Create(ctx context.Context, community community.Community, user account.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, community, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCommunityMockRecorder) Create(ctx, community, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommunity)(nil).Create), ctx, community, user)
}

// FindByID mocks base method.
func (m *MockCommunity) FindByID(ctx context.Context, id community.ID) (*community.Community, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*community.Community)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCommunityMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCommunity)(nil).FindByID), ctx, id)
}

// FindByName mocks base method.
func (m *MockCommunity) FindByName(ctx context.Context, name community.Name) (*community.Community, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", ctx, name)
	ret0, _ := ret[0].(*community.Community)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockCommunityMockRecorder) FindByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockCommunity)(nil).FindByName), ctx, name)
}

// FindInvitation mocks base method.
func (m *MockCommunity) FindInvitation(ctx context.Context, invitation community.Invitation) (*community.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInvitation", ctx, invitation)
	ret0, _ := ret[0].(*community.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindInvitation indicates an expected call of FindInvitation.
func (mr *MockCommunityMockRecorder) FindInvitation(ctx, invitation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInvitation", reflect.TypeOf((*MockCommunity)(nil).FindInvitation), ctx, invitation)
}

// GetUserRole mocks base method.
func (m *MockCommunity) GetUserRole(ctx context.Context, id community.ID, user account.ID) (community.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", ctx, id, user)
	ret0, _ := ret[0].(community.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockCommunityMockRecorder) GetUserRole(ctx, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockCommunity)(nil).GetUserRole), ctx, id, user)
}

// IssueInvitation mocks base method.
func (m *MockCommunity) IssueInvitation(ctx context.Context, id community.ID, user account.ID) (*community.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueInvitation", ctx, id, user)
	ret0, _ := ret[0].(*community.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueInvitation indicates an expected call of IssueInvitation.
func (mr *MockCommunityMockRecorder) IssueInvitation(ctx, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueInvitation", reflect.TypeOf((*MockCommunity)(nil).IssueInvitation), ctx, id, user)
}

// Join mocks base method.
func (m *MockCommunity) Join(ctx context.Context, id community.ID, accountID account.ID, role community.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join", ctx, id, accountID, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// Join indicates an expected call of Join.
func (mr *MockCommunityMockRecorder) Join(ctx, id, accountID, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockCommunity)(nil).Join), ctx, id, accountID, role)
}
