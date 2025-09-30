package tests

import (
	"app/internal/dto"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProductCreate(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		input := &dto.Product{
			Name:        "Handmade mug",
			Description: "Ceramic mug",
			Price:       1200,
			Stock:       10,
		}

		created := &dto.Product{
			ID:          1,
			Name:        input.Name,
			Description: input.Description,
			Price:       input.Price,
			Stock:       input.Stock,
		}

		deps.ProdMock.On("Create", mock.Anything, input).Return(created, nil)
		deps.LogMock.On("Info", "successfully product created", mock.Anything).Return()

		got, err := deps.UC.ProductCreate(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Equal(t, created.ID, got.ID)
		assert.Equal(t, created.Name, got.Name)
		assert.Equal(t, created.Description, got.Description)
		assert.Equal(t, created.Price, got.Price)
		assert.Equal(t, created.Stock, got.Stock)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		input := &dto.Product{
			Name:        "Handmade mug",
			Description: "Ceramic mug",
			Price:       1200,
			Stock:       10,
		}

		testErr := errors.New("db error")

		deps.ProdMock.On("Create", mock.Anything, input).Return(((*dto.Product)(nil)), testErr)
		deps.LogMock.On("Error", "failed to create product", mock.Anything).Return()

		got, err := deps.UC.ProductCreate(ctx, input)
		require.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestProductDelete(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		productID := int64(1)

		deps.ProdMock.On("Delete", mock.Anything, productID).Return(nil)
		deps.LogMock.On("Info", "successfully delete product", mock.Anything).Return()

		err := deps.UC.ProductDelete(ctx, productID)
		require.NoError(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		productID := int64(2)
		testErr := errors.New("db error")

		deps.ProdMock.On("Delete", mock.Anything, productID).Return(testErr)
		deps.LogMock.On("Error", "failed to delete product", mock.Anything).Return()

		err := deps.UC.ProductDelete(ctx, productID)
		require.Error(t, err)
	})
}

func TestProductUpdate(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		input := &dto.Product{
			ID:          1,
			Name:        "Mug v1",
			Description: "Old desc",
			Price:       1000,
			Stock:       5,
		}

		updated := &dto.Product{
			ID:          1,
			Name:        "Mug v2",
			Description: "New desc",
			Price:       1200,
			Stock:       10,
		}

		deps.ProdMock.On("Update", mock.Anything, input).Return(updated, nil)
		deps.LogMock.On("Info", "successfully product updated", mock.Anything).Return()

		got, err := deps.UC.ProductUpdate(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Equal(t, updated.ID, got.ID)
		assert.Equal(t, updated.Name, got.Name)
		assert.Equal(t, updated.Description, got.Description)
		assert.Equal(t, updated.Price, got.Price)
		assert.Equal(t, updated.Stock, got.Stock)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		input := &dto.Product{
			ID:          2,
			Name:        "Mug",
			Description: "Desc",
			Price:       900,
			Stock:       3,
		}

		testErr := errors.New("update failed")
		deps.ProdMock.On("Update", mock.Anything, input).Return(((*dto.Product)(nil)), testErr)
		deps.LogMock.On("Error", "failed to update product", mock.Anything).Return()

		got, err := deps.UC.ProductUpdate(ctx, input)
		require.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestProductList(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		products := []*dto.Product{
			{ID: 1, Name: "Product A", Price: 100, Stock: 5, CreatedAt: time.Now()},
			{ID: 2, Name: "Product B", Price: 200, Stock: 3, CreatedAt: time.Now()},
		}

		deps.ProdMock.On("List", mock.Anything).Return(products, nil)
		deps.LogMock.On("Info", "successfully list products", mock.Anything).Return()

		got, err := deps.UC.ProductList(ctx)
		require.NoError(t, err)
		require.Len(t, got, 2)
		assert.Equal(t, products[0].ID, got[0].ID)
		assert.Equal(t, products[1].ID, got[1].ID)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()

		testErr := errors.New("db error")
		deps.ProdMock.On("List", mock.Anything).Return(([]*dto.Product)(nil), testErr)
		deps.LogMock.On("Error", "failed to list products", mock.Anything).Return()

		got, err := deps.UC.ProductList(ctx)
		require.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestProductGetByID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		productID := int64(1)

		expected := &dto.Product{
			ID:          productID,
			Name:        "Handmade mug",
			Description: "Ceramic mug",
			Price:       1200,
			Stock:       10,
		}

		deps.ProdMock.On("GetByID", mock.Anything, productID).Return(expected, nil)
		deps.LogMock.On("Info", "successfully get product", mock.Anything).Return()

		got, err := deps.UC.ProductGetByID(ctx, productID)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Equal(t, expected.ID, got.ID)
		assert.Equal(t, expected.Name, got.Name)
		assert.Equal(t, expected.Description, got.Description)
		assert.Equal(t, expected.Price, got.Price)
		assert.Equal(t, expected.Stock, got.Stock)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		productID := int64(2)

		testErr := errors.New("not found")
		deps.ProdMock.On("GetByID", mock.Anything, productID).Return(((*dto.Product)(nil)), testErr)
		deps.LogMock.On("Error", "failed to get product", mock.Anything).Return()

		got, err := deps.UC.ProductGetByID(ctx, productID)
		require.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestProductAddPhoto(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		productID := int64(1)
		fileID := "file_123"

		deps.ProdMock.On("AddPhoto", mock.Anything, productID, fileID).Return(nil)
		deps.LogMock.On("Info", "successfully add photo", mock.Anything).Return()

		err := deps.UC.ProductAddPhoto(ctx, productID, fileID)
		require.NoError(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		productID := int64(2)
		fileID := "file_404"
		testErr := errors.New("db error")

		deps.ProdMock.On("AddPhoto", mock.Anything, productID, fileID).Return(testErr)
		deps.LogMock.On("Error", "failed to add photo", mock.Anything).Return()

		err := deps.UC.ProductAddPhoto(ctx, productID, fileID)
		require.Error(t, err)
	})
}

func TestProductRemovePhoto(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		productID := int64(1)

		deps.ProdMock.On("RemovePhotos", mock.Anything, productID).Return(nil)
		deps.LogMock.On("Info", "successfully remove photo", mock.Anything).Return()

		err := deps.UC.ProductRemovePhoto(ctx, productID)
		require.NoError(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		deps := NewTestDeps(t)
		ctx := context.Background()
		productID := int64(2)
		testErr := errors.New("db error")

		deps.ProdMock.On("RemovePhotos", mock.Anything, productID).Return(testErr)
		deps.LogMock.On("Error", "failed to remove photo", mock.Anything).Return()

		err := deps.UC.ProductRemovePhoto(ctx, productID)
		require.Error(t, err)
	})
}
