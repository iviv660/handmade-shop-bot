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
	result := db.gorm.WithContext(ctx).Where("product_id = ?", productID).Delete(&dto.Product{})
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
	var model dto.Product
	if err := db.gorm.WithContext(ctx).First(&model, productID).Error; err != nil {
		return nil, err
	}

	var photos []dto.ProductPhoto
	if err := db.gorm.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&photos).Error; err != nil {
		return nil, err
	}

	product := &dto.Product{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Stock:       model.Stock,
		CreatedAt:   model.CreatedAt,
		Photos:      make([]string, 0, len(photos)),
	}

	for _, p := range photos {
		product.Photos = append(product.Photos, p.FileID)
	}

	return product, nil
}

func (db *ProductDB) AddPhoto(ctx context.Context, productID int64, fileID string) error {
	photo := dto.ProductPhoto{
		ProductID: productID,
		FileID:    fileID,
	}

	result := db.gorm.WithContext(ctx).Create(&photo)
	return result.Error
}

func (db *ProductDB) RemovePhotos(ctx context.Context, productID int64) error {
	result := db.gorm.WithContext(ctx).
		Where("product_id = ?", productID).
		Delete(&dto.ProductPhoto{})
	return result.Error
}
