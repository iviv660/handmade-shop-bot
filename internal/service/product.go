package service

import (
	"app/internal/dto"
	"context"
)

func (uc *UseCase) ProductCreate(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	newProduct, err := uc.productService.Create(ctx, product)
	if err != nil {
		uc.logger.Error("failed to create product", map[string]any{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"error":       err.Error(),
		})
		return nil, err
	}
	uc.logger.Info("successfully product created", map[string]any{
		"product_id":  newProduct.ID,
		"name":        newProduct.Name,
		"description": newProduct.Description,
		"price":       newProduct.Price,
		"stock":       newProduct.Stock,
	})

	return newProduct, nil
}

func (uc *UseCase) ProductDelete(ctx context.Context, productID int64) error {
	err := uc.productService.Delete(ctx, productID)
	if err != nil {
		uc.logger.Error("failed to delete product", map[string]any{
			"product_id": productID,
			"error":      err.Error(),
		})
		return err
	}
	uc.logger.Info("successfully delete product", map[string]any{
		"product_id": productID,
	})
	return nil
}

func (uc *UseCase) ProductUpdate(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	newProduct, err := uc.productService.Update(ctx, product)
	if err != nil {
		uc.logger.Error("failed to update product", map[string]any{
			"product_id":  product.ID,
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"error":       err.Error(),
		})
		return nil, err
	}
	uc.logger.Info("successfully product updated", map[string]any{
		"product_id":  newProduct.ID,
		"name":        newProduct.Name,
		"description": newProduct.Description,
		"price":       newProduct.Price,
		"stock":       newProduct.Stock,
	})

	return newProduct, nil
}

func (uc *UseCase) ProductList(ctx context.Context) ([]*dto.Product, error) {
	products, err := uc.productService.List(ctx)
	if err != nil {
		uc.logger.Error("failed to list products", map[string]any{
			"error": err.Error(),
		})
		return nil, err
	}
	uc.logger.Info("successfully list products", map[string]any{
		"count_products": len(products),
	})
	return products, nil
}

func (uc *UseCase) ProductGetByID(ctx context.Context, productID int64) (*dto.Product, error) {
	product, err := uc.productService.GetByID(ctx, productID)
	if err != nil {
		uc.logger.Error("failed to get product", map[string]any{
			"product_id": productID,
			"error":      err.Error(),
		})
		return nil, err
	}
	uc.logger.Info("successfully get product", map[string]any{
		"product_id":  productID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock,
	})
	return product, nil
}

func (uc *UseCase) ProductAddPhoto(ctx context.Context, productID int64, fileID string) error {
	err := uc.productService.AddPhoto(ctx, productID, fileID)
	if err != nil {
		uc.logger.Error("failed to add photo", map[string]any{
			"product_id": productID,
			"file_id":    fileID,
			"error":      err.Error(),
		})
		return err
	}
	uc.logger.Info("successfully add photo", map[string]any{
		"product_id": productID,
		"file_id":    fileID,
	})
	return nil
}

func (uc *UseCase) ProductRemovePhoto(ctx context.Context, productID int64) error {
	err := uc.productService.RemovePhotos(ctx, productID)
	if err != nil {
		uc.logger.Error("failed to remove photo", map[string]any{
			"product_id": productID,
		})
		return err
	}
	uc.logger.Info("successfully remove photo", map[string]any{
		"product_id": productID,
	})
	return nil
}
