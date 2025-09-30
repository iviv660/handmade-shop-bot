package mocks

import (
	"app/internal/dto"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockProduct struct {
	mock.Mock
}

func NewMockProduct(t *testing.T) *MockProduct {
	m := &MockProduct{}
	t.Cleanup(func() {
		m.AssertExpectations(t)
	})
	return m
}

func (m *MockProduct) Create(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	args := m.Called(ctx, product)
	var p *dto.Product
	if v := args.Get(0); v != nil {
		p = v.(*dto.Product)
	}
	return p, args.Error(1)
}

func (m *MockProduct) Delete(ctx context.Context, productID int64) error {
	args := m.Called(ctx, productID)
	return args.Error(0)
}

func (m *MockProduct) Update(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	args := m.Called(ctx, product)
	var p *dto.Product
	if v := args.Get(0); v != nil {
		p = v.(*dto.Product)
	}
	return p, args.Error(1)
}

func (m *MockProduct) List(ctx context.Context) ([]*dto.Product, error) {
	args := m.Called(ctx)
	var list []*dto.Product
	if v := args.Get(0); v != nil {
		list = v.([]*dto.Product)
	}
	return list, args.Error(1)
}

func (m *MockProduct) GetByID(ctx context.Context, productID int64) (*dto.Product, error) {
	args := m.Called(ctx, productID)
	var p *dto.Product
	if v := args.Get(0); v != nil {
		p = v.(*dto.Product)
	}
	return p, args.Error(1)
}

func (m *MockProduct) AddPhoto(ctx context.Context, productID int64, fileID string) error {
	args := m.Called(ctx, productID, fileID)
	return args.Error(0)
}

func (m *MockProduct) RemovePhotos(ctx context.Context, productID int64) error {
	args := m.Called(ctx, productID)
	return args.Error(0)
}
