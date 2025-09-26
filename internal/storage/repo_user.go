package storage

import (
	"bot/internal/dto"
	"context"

	"gorm.io/gorm"
)

type UserDB struct {
	gorm *gorm.DB
}

func NewUserDB(gorm *gorm.DB) *UserDB {
	return &UserDB{gorm: gorm}
}

func (db *UserDB) Register(ctx context.Context, telegramID int64, username string) (*dto.User, error) {
	user := &dto.User{TelegramID: telegramID, Username: username}
	result := db.gorm.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (db *UserDB) GetByID(ctx context.Context, userID int64) (*dto.User, error) {
	user := &dto.User{}
	result := db.gorm.WithContext(ctx).Where("id = ?", userID).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (db *UserDB) GetByTelegramID(ctx context.Context, telegramID int64) (*dto.User, error) {
	user := &dto.User{}
	result := db.gorm.WithContext(ctx).Where("telegram_id = ?", telegramID).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
