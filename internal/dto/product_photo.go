package dto

type ProductPhoto struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	ProductID int64  `json:"product_id" gorm:"index;not null"`
	FileID    string `json:"file_id" gorm:"not null"`
}
