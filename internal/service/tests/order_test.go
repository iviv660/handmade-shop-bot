package tests

import (
	"app/internal/dto"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestOrderCreate(t *testing.T) {
	t.Parallel()

	deps := NewTestDeps(t)
	ctx := context.Background()

	input := &dto.Order{
		UserID:     1,
		ProductID:  1,
		TotalPrice: 500,
		Quantity:   1,
	}

	expected := &dto.Order{
		ID:         1,
		UserID:     1,
		ProductID:  1,
		Status:     "pending",
		TotalPrice: 500,
		Quantity:   1,
		PaymentID:  "2435",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	deps.OrdMock.On("Create", mock.Anything, mock.Anything).Return(expected, nil)
	deps.LogMock.On("Info", "successfully created order", mock.Anything).Return()

	got, err := deps.UC.OrderCreate(ctx, input)
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, expected.ID, got.ID)
	assert.Equal(t, expected.UserID, got.UserID)
	assert.Equal(t, expected.ProductID, got.ProductID)
	assert.Equal(t, expected.Status, got.Status)
	assert.Equal(t, expected.TotalPrice, got.TotalPrice)
	assert.Equal(t, expected.Quantity, got.Quantity)
	assert.Equal(t, expected.PaymentID, got.PaymentID)
	assert.WithinDuration(t, expected.CreatedAt, got.CreatedAt, time.Second)
	assert.WithinDuration(t, expected.UpdatedAt, got.UpdatedAt, time.Second)
}

func TestOrderListByUserID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		userID := int64(1)

		orders := []*dto.Order{
			{ID: 1, UserID: userID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, UserID: userID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		deps.OrdMock.On("ListByUserID", mock.Anything, userID).Return(orders, nil)
		deps.LogMock.On("Info", "successfully fetched orders", mock.Anything).Return()

		got, err := deps.UC.OrderListByUserID(ctx, userID)
		require.NoError(t, err)
		require.Len(t, got, 2)
		assert.Equal(t, orders[0].ID, got[0].ID)
		assert.Equal(t, orders[1].ID, got[1].ID)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		userID := int64(1)

		testErr := errors.New("db failure")
		deps.OrdMock.On("ListByUserID", mock.Anything, userID).Return(([]*dto.Order)(nil), testErr)
		deps.LogMock.On("Error", "failed to list orders", mock.Anything).Return()

		got, err := deps.UC.OrderListByUserID(ctx, userID)
		require.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestOrderUpdateStatus(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		orderID := int64(1)
		status := "shipped"

		deps.OrdMock.On("UpdateStatus", mock.Anything, orderID, status).Return(nil)
		deps.LogMock.On("Info", "order status updated successfully", mock.Anything).Return()

		err := deps.UC.OrderUpdateStatus(ctx, orderID, status)
		require.NoError(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		orderID := int64(1)
		status := "shipped"
		testErr := errors.New("db failure")

		deps.OrdMock.On("UpdateStatus", mock.Anything, orderID, status).Return(testErr)
		deps.LogMock.On("Error", "failed to update order status", mock.Anything).Return()

		err := deps.UC.OrderUpdateStatus(ctx, orderID, status)
		require.Error(t, err)
	})
}

func TestOrderAttachPaymentID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		orderID := int64(1)
		paymentID := "pay_123"

		deps.OrdMock.On("AttachPaymentID", mock.Anything, orderID, paymentID).Return(nil)
		deps.LogMock.On("Info", "payment id attached to order successfully", mock.Anything).Return()

		err := deps.UC.OrderAttachPaymentID(ctx, orderID, paymentID)
		require.NoError(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		orderID := int64(1)
		paymentID := "pay_123"
		testErr := errors.New("db failure")

		deps.OrdMock.On("AttachPaymentID", mock.Anything, orderID, paymentID).Return(testErr)
		deps.LogMock.On("Error", "failed to attach payment id to order", mock.Anything).Return()

		err := deps.UC.OrderAttachPaymentID(ctx, orderID, paymentID)
		require.Error(t, err)
	})
}

func TestOrderGetByID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		orderID := int64(1)

		expected := &dto.Order{
			ID:         orderID,
			UserID:     1,
			ProductID:  1,
			Status:     "pending",
			Quantity:   1,
			TotalPrice: 500,
		}

		deps.OrdMock.On("GetByID", mock.Anything, orderID).Return(expected, nil)
		deps.LogMock.On("Info", "successfully fetched order by id", mock.Anything).Return()

		got, err := deps.UC.OrderGetByID(ctx, orderID)
		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, expected.ID, got.ID)
		require.Equal(t, expected.UserID, got.UserID)
		require.Equal(t, expected.ProductID, got.ProductID)
		require.Equal(t, expected.Status, got.Status)
		require.Equal(t, expected.Quantity, got.Quantity)
		require.Equal(t, expected.TotalPrice, got.TotalPrice)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		orderID := int64(1)

		testErr := errors.New("not found")
		deps.OrdMock.On("GetByID", mock.Anything, orderID).Return(((*dto.Order)(nil)), testErr)
		deps.LogMock.On("Error", "failed to get order by id", mock.Anything).Return()

		got, err := deps.UC.OrderGetByID(ctx, orderID)
		require.Error(t, err)
		require.Nil(t, got)
	})
}
