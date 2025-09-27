package storage

import (
	"app/internal/dto"
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
	tx := db.gorm.WithContext(ctx).Begin()
	if err := tx.Model(&dto.Product{}).
		Where("id = ? AND stock >= ?", order.ProductID, order.Quantity).
		Update("stock", gorm.Expr("stock - ?", order.Quantity)).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return order, nil
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

func (db *OrderDB) AttachPaymentID(ctx context.Context, orderID int64, paymentID string) error {
	result := db.gorm.WithContext(ctx).
		Model(&dto.Order{}).
		Where("id = ?", orderID).
		Update("payment_id", paymentID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *OrderDB) GetByID(ctx context.Context, orderID int64) (*dto.Order, error) {
	var order dto.Order
	result := db.gorm.WithContext(ctx).First(&order, orderID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}
