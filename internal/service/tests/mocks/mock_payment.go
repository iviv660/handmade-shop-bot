package mocks

import (
	"app/internal/dto"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockPayment struct {
	mock.Mock
}

func NewMockPayment(t *testing.T) *MockPayment {
	m := &MockPayment{}
	t.Cleanup(func() {
		m.AssertExpectations(t)
	})
	return m
}

func (m *MockPayment) Create(ctx context.Context, order *dto.Order) (string, string, error) {
	args := m.Called(ctx, order)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockPayment) CheckStatus(ctx context.Context, paymentID string) (string, error) {
	args := m.Called(ctx, paymentID)
	return args.String(0), args.Error(1)
}
