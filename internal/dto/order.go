package dto

import "time"

type Order struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	UserID     int64     `json:"user_id" gorm:"index"`
	ProductID  int64     `json:"product_id" gorm:"index"`
	Status     string    `json:"status" gorm:"default:pending"`
	TotalPrice float64   `json:"total_price" gorm:"column:total_amount"`
	Quantity   int       `json:"quantity"`
	PaymentID  string    `json:"payment_id" gorm:"column:payment_id"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
