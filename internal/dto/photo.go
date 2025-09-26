package dto

type ProductPhotoModel struct {
	ID        int64  `gorm:"primaryKey"`
	ProductID int64  `gorm:"index;not null"`
	FileID    string `gorm:"not null"`
}
