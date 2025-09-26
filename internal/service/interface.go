package service

import (
	"bot/internal/dto"
	"context"
)

type UseCase struct {
	userService    UserService
	productService ProductService
	orderService   OrderService
	logger         LoggerService
}

func NewUseCase(userService UserService, productService ProductService, orderService OrderService, logger LoggerService) *UseCase {
	return &UseCase{
		userService:    userService,
		productService: productService,
		orderService:   orderService,
		logger:         logger,
	}
}

type UserService interface {
	Register(ctx context.Context, telegramID int64, username string) (*dto.User, error)
	GetByID(ctx context.Context, userID int64) (*dto.User, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*dto.User, error)
}

type ProductService interface {
	Create(ctx context.Context, product *dto.Product) (*dto.Product, error)
	Delete(ctx context.Context, productID int64) error
	List(ctx context.Context) ([]*dto.Product, error)
	GetByID(ctx context.Context, productID int64) (*dto.Product, error)
	AddPhoto(ctx context.Context, productID int64, fileID string) error
}

type OrderService interface {
	Create(ctx context.Context, order *dto.Order) (*dto.Order, error)
	ListByUserID(ctx context.Context, userID int64) ([]*dto.Order, error)
	UpdateStatus(ctx context.Context, orderID int64, status string) error
}

type LoggerService interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}
