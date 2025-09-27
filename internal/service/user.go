package service

import (
	"app/internal/dto"
	"context"
)

func (uc *UseCase) UserRegister(ctx context.Context, telegramID int64, username string) (*dto.User, error) {
	user, err := uc.userService.Register(ctx, telegramID, username)
	if err != nil {
		uc.logger.Error("failed to register user", map[string]any{
			"telegram_id": telegramID,
			"username":    username,
			"error":       err.Error(),
		})
		return nil, err
	}

	uc.logger.Info("user registered successfully", map[string]any{
		"user_id":     user.ID,
		"telegram_id": telegramID,
	})

	return user, nil
}

func (uc *UseCase) UserGetByID(ctx context.Context, userID int64) (*dto.User, error) {
	user, err := uc.userService.GetByID(ctx, userID)
	if err != nil {
		uc.logger.Error("failed to get by id user", map[string]any{
			"user_id": userID,
			"error":   err.Error(),
		})
		return nil, err
	}
	uc.logger.Info("user get by id successfully", map[string]any{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
	return user, nil
}

func (uc *UseCase) UserGetByTelegramID(ctx context.Context, telegramID int64) (*dto.User, error) {
	user, err := uc.userService.GetByTelegramID(ctx, telegramID)
	if err != nil {
		uc.logger.Error("failed to get by telegram id user", map[string]any{
			"telegram_id": telegramID,
			"error":       err.Error(),
		})
		return nil, err
	}
	uc.logger.Info("user get by telegram id successfully", map[string]any{
		"user_id":     user.ID,
		"telegram_id": user.TelegramID,
		"username":    user.Username,
	})
	return user, nil
}
