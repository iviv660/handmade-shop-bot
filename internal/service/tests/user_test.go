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

func TestUserRegister(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		telegramID := int64(12345)
		username := "vlad"

		expected := &dto.User{
			ID:         1,
			TelegramID: telegramID,
			Username:   username,
		}

		deps.UserMock.On("Register", mock.Anything, telegramID, username).Return(expected, nil)
		deps.LogMock.On("Info", "user registered successfully", mock.Anything).Return()

		got, err := deps.UC.UserRegister(ctx, telegramID, username)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Equal(t, expected.ID, got.ID)
		assert.Equal(t, expected.TelegramID, got.TelegramID)
		assert.Equal(t, expected.Username, got.Username)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		telegramID := int64(54321)
		username := "erruser"

		testErr := errors.New("register failed")
		deps.UserMock.On("Register", mock.Anything, telegramID, username).Return(((*dto.User)(nil)), testErr)
		deps.LogMock.On("Error", "failed to register user", mock.Anything).Return()

		got, err := deps.UC.UserRegister(ctx, telegramID, username)
		require.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestUserGetByID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		userID := int64(1)

		expected := &dto.User{
			ID:       userID,
			Username: "vlad",
			Role:     "student",
		}

		deps.UserMock.On("GetByID", mock.Anything, userID).Return(expected, nil)
		deps.LogMock.On("Info", "user get by id successfully", mock.Anything).Return()

		got, err := deps.UC.UserGetByID(ctx, userID)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Equal(t, expected.ID, got.ID)
		assert.Equal(t, expected.Username, got.Username)
		assert.Equal(t, expected.Role, got.Role)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		userID := int64(2)

		testErr := errors.New("not found")
		deps.UserMock.On("GetByID", mock.Anything, userID).Return(((*dto.User)(nil)), testErr)
		deps.LogMock.On("Error", "failed to get by id user", mock.Anything).Return()

		got, err := deps.UC.UserGetByID(ctx, userID)
		require.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestUserGetByTelegramID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		telegramID := int64(12345)

		expected := &dto.User{
			ID:         1,
			TelegramID: telegramID,
			Username:   "vlad",
			Role:       "student",
		}

		deps.UserMock.On("GetByTelegramID", mock.Anything, telegramID).Return(expected, nil)
		deps.LogMock.On("Info", "user get by telegram id successfully", mock.Anything).Return()

		got, err := deps.UC.UserGetByTelegramID(ctx, telegramID)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Equal(t, expected.ID, got.ID)
		assert.Equal(t, expected.TelegramID, got.TelegramID)
		assert.Equal(t, expected.Username, got.Username)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		telegramID := int64(54321)

		testErr := errors.New("not found")
		deps.UserMock.On("GetByTelegramID", mock.Anything, telegramID).Return(((*dto.User)(nil)), testErr)
		deps.LogMock.On("Error", "failed to get by telegram id user", mock.Anything).Return()

		got, err := deps.UC.UserGetByTelegramID(ctx, telegramID)
		require.Error(t, err)
		assert.Nil(t, got)
	})
}
