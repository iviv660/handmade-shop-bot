package mocks

import (
	"app/internal/dto"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

func NewMockUser(t *testing.T) *MockUser {
	m := &MockUser{}
	t.Cleanup(func() {
		m.AssertExpectations(t)
	})
	return m
}

func (m *MockUser) Register(ctx context.Context, telegramID int64, username string) (*dto.User, error) {
	args := m.Called(ctx, telegramID, username)
	var u *dto.User
	if v := args.Get(0); v != nil {
		u = v.(*dto.User)
	}
	return u, args.Error(1)
}

func (m *MockUser) GetByID(ctx context.Context, userID int64) (*dto.User, error) {
	args := m.Called(ctx, userID)
	var u *dto.User
	if v := args.Get(0); v != nil {
		u = v.(*dto.User)
	}
	return u, args.Error(1)
}

func (m *MockUser) GetByTelegramID(ctx context.Context, telegramID int64) (*dto.User, error) {
	args := m.Called(ctx, telegramID)
	var u *dto.User
	if v := args.Get(0); v != nil {
		u = v.(*dto.User)
	}
	return u, args.Error(1)
}
