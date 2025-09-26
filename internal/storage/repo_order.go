package storage

import (
	"bot/internal/dto"
	"context"

	"gorm.io/gorm"
)

type OrderDB struct {
	gorm *gorm.DB
}

func NewOrderDB(gorm *gorm.DB) *OrderDB {
	return &OrderDB{gorm: gorm}
}

func (db *OrderDB) Create(ctx context.Context, order *dto.Order) (*dto.Order, error) {
	result := db.gorm.WithContext(ctx).Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil // уже содержит ID
}

func (db *OrderDB) ListByUserID(ctx context.Context, userID int64) ([]*dto.Order, error) {
	orders := []*dto.Order{}
	result := db.gorm.WithContext(ctx).Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (db *OrderDB) UpdateStatus(ctx context.Context, orderID int64, status string) error {
	result := db.gorm.WithContext(ctx).
		Model(&dto.Order{}).
		Where("id = ?", orderID).
		Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
