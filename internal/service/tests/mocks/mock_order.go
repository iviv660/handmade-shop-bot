package mocks

import (
	"app/internal/dto"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockOrder struct {
	mock.Mock
}

func NewMockOrder(t *testing.T) *MockOrder {
	m := &MockOrder{}
	t.Cleanup(func() {
		m.AssertExpectations(t)
	})
	return m
}

func (m *MockOrder) Create(ctx context.Context, order *dto.Order) (*dto.Order, error) {
	args := m.Called(ctx, order)
	var o *dto.Order
	if v := args.Get(0); v != nil {
		o = v.(*dto.Order)
	}
	return o, args.Error(1)
}

func (m *MockOrder) ListByUserID(ctx context.Context, userID int64) ([]*dto.Order, error) {
	args := m.Called(ctx, userID)
	var list []*dto.Order
	if v := args.Get(0); v != nil {
		list = v.([]*dto.Order)
	}
	return list, args.Error(1)
}

func (m *MockOrder) UpdateStatus(ctx context.Context, orderID int64, status string) error {
	args := m.Called(ctx, orderID, status)
	return args.Error(0)
}

func (m *MockOrder) AttachPaymentID(ctx context.Context, orderID int64, paymentID string) error {
	args := m.Called(ctx, orderID, paymentID)
	return args.Error(0)
}

func (m *MockOrder) GetByID(ctx context.Context, orderID int64) (*dto.Order, error) {
	args := m.Called(ctx, orderID)
	var o *dto.Order
	if v := args.Get(0); v != nil {
		o = v.(*dto.Order)
	}
	return o, args.Error(1)
}
