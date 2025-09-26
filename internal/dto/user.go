package dto

import "time"

type User struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	TelegramID int64     `json:"telegram_id" gorm:"uniqueIndex"`
	Username   string    `json:"username"`
	Role       string    `json:"role" gorm:"default:user"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}
