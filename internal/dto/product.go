package dto

import "time"

type Product struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	PhotoID     string    `json:"photo_id" gorm:"photo_id"`
}
