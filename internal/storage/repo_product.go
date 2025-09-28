package storage

import (
	"app/internal/dto"
	"context"

	"gorm.io/gorm"
)

type ProductDB struct {
	gorm *gorm.DB
}

func NewProductDB(gorm *gorm.DB) *ProductDB {
	return &ProductDB{gorm: gorm}
}

func (db *ProductDB) Create(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	result := db.gorm.WithContext(ctx).Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (db *ProductDB) Delete(ctx context.Context, productID int64) error {
	result := db.gorm.WithContext(ctx).Where("id = ?", productID).Delete(&dto.Product{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *ProductDB) Update(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	result := db.gorm.WithContext(ctx).
		Model(&dto.Product{}).
		Where("id = ?", product.ID).
		Updates(map[string]any{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
		})

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (db *ProductDB) List(ctx context.Context) ([]*dto.Product, error) {
	rec := []*dto.Product{}
	result := db.gorm.WithContext(ctx).Find(&rec)
	if result.Error != nil {
		return nil, result.Error
	}
	return rec, nil
}

func (db *ProductDB) GetByID(ctx context.Context, productID int64) (*dto.Product, error) {
	var product dto.Product
	if err := db.gorm.WithContext(ctx).
		First(&product, productID).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (db *ProductDB) AddPhoto(ctx context.Context, productID int64, fileID string) error {
	result := db.gorm.WithContext(ctx).
		Model(&dto.Product{}).
		Where("id = ?", productID).
		Update("photo_id", fileID)
	return result.Error
}

func (db *ProductDB) RemovePhotos(ctx context.Context, productID int64) error {
	return db.gorm.WithContext(ctx).
		Model(&dto.Product{}).
		Where("id = ?", productID).
		Update("photo_id", nil).
		Error
}
