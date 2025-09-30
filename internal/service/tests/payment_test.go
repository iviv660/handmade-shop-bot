package tests

import (
	"app/internal/dto"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreatePayment(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		order := &dto.Order{
			ID:         1,
			UserID:     10,
			ProductID:  20,
			TotalPrice: 500,
			Quantity:   1,
		}

		expectedURL := "https://pay.example/checkout/1"
		expectedPID := "pay_123"

		deps.PayMock.On("Create", mock.Anything, mock.Anything).Return(expectedURL, expectedPID, nil)
		deps.OrdMock.On("AttachPaymentID", mock.Anything, order.ID, expectedPID).Return(nil)
		deps.LogMock.On("Info", "payment created", mock.Anything).Return()

		url, pid, err := deps.UC.CreatePayment(ctx, order)
		require.NoError(t, err)
		assert.Equal(t, expectedURL, url)
		assert.Equal(t, expectedPID, pid)
	})

	t.Run("payment creation fails", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		order := &dto.Order{
			ID:         2,
			UserID:     11,
			ProductID:  21,
			TotalPrice: 700,
			Quantity:   2,
		}

		testErr := errors.New("payment gateway down")

		deps.PayMock.On("Create", mock.Anything, mock.Anything).Return("", "", testErr)
		deps.LogMock.On("Error", "failed to create payment", mock.Anything).Return()

		url, pid, err := deps.UC.CreatePayment(ctx, order)
		require.Error(t, err)
		assert.Equal(t, "", url)
		assert.Equal(t, "", pid)
	})

	t.Run("attach payment id fails but returns url/pid", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		order := &dto.Order{
			ID:         3,
			UserID:     12,
			ProductID:  22,
			TotalPrice: 900,
			Quantity:   1,
		}

		expectedURL := "https://pay.example/checkout/3"
		expectedPID := "pay_789"
		attachErr := errors.New("db write failed")

		deps.PayMock.On("Create", mock.Anything, mock.Anything).Return(expectedURL, expectedPID, nil)
		deps.OrdMock.On("AttachPaymentID", mock.Anything, order.ID, expectedPID).Return(attachErr)
		deps.LogMock.On("Error", "failed to attach payment id", mock.Anything).Return()
		deps.LogMock.On("Info", "payment created", mock.Anything).Return()

		url, pid, err := deps.UC.CreatePayment(ctx, order)
		require.NoError(t, err)
		assert.Equal(t, expectedURL, url)
		assert.Equal(t, expectedPID, pid)
	})
}

func TestCheckPayment(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		paymentID := "pay_123"
		status := "paid"

		deps.PayMock.On("CheckStatus", mock.Anything, paymentID).Return(status, nil)
		deps.LogMock.On("Info", "payment status checked", mock.Anything).Return()

		got, err := deps.UC.CheckPayment(ctx, paymentID)
		require.NoError(t, err)
		assert.Equal(t, status, got)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		paymentID := "pay_404"
		testErr := errors.New("gateway error")

		deps.PayMock.On("CheckStatus", mock.Anything, paymentID).Return("", testErr)
		deps.LogMock.On("Error", "failed to check payment", mock.Anything).Return()

		got, err := deps.UC.CheckPayment(ctx, paymentID)
		require.Error(t, err)
		assert.Equal(t, "", got)
	})
}
