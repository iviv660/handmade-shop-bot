package service

import (
	"app/internal/dto"
	"context"
)

type UseCase struct {
	userService    UserService
	productService ProductService
	orderService   OrderService
	paymentService PaymentService
	logger         LoggerService
}

func NewUseCase(
	userService UserService,
	productService ProductService,
	orderService OrderService,
	paymentService PaymentService,
	logger LoggerService,
) *UseCase {
	return &UseCase{
		userService:    userService,
		productService: productService,
		orderService:   orderService,
		paymentService: paymentService,
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
	Update(ctx context.Context, product *dto.Product) (*dto.Product, error)
	List(ctx context.Context) ([]*dto.Product, error)
	GetByID(ctx context.Context, productID int64) (*dto.Product, error)
	AddPhoto(ctx context.Context, productID int64, fileID string) error
	RemovePhotos(ctx context.Context, productID int64) error
}

type OrderService interface {
	Create(ctx context.Context, order *dto.Order) (*dto.Order, error)
	ListByUserID(ctx context.Context, userID int64) ([]*dto.Order, error)
	UpdateStatus(ctx context.Context, orderID int64, status string) error
	AttachPaymentID(ctx context.Context, orderID int64, paymentID string) error
	GetByID(ctx context.Context, orderID int64) (*dto.Order, error)
}

type LoggerService interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type PaymentService interface {
	Create(ctx context.Context, order *dto.Order) (string, string, error)
	CheckStatus(ctx context.Context, paymentID string) (string, error)
}
