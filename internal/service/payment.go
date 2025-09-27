package service

import (
	"app/internal/dto"
	"context"
)

func (uc *UseCase) CreatePayment(ctx context.Context, order *dto.Order) (string, string, error) {
	paymentURL, paymentID, err := uc.paymentService.Create(ctx, order)
	if err != nil {
		uc.logger.Error("failed to create payment", map[string]any{
			"order_id":    order.ID,
			"user_id":     order.UserID,
			"product_id":  order.ProductID,
			"total_price": order.TotalPrice,
			"quantity":    order.Quantity,
			"error":       err.Error(),
		})
		return "", "", err
	}

	if err := uc.orderService.AttachPaymentID(ctx, order.ID, paymentID); err != nil {
		uc.logger.Error("failed to attach payment id", map[string]any{
			"order_id":   order.ID,
			"payment_id": paymentID,
			"error":      err.Error(),
		})
		// даже если ошибка сохранения, ссылку можно вернуть
	}

	uc.logger.Info("payment created", map[string]any{
		"order_id":    order.ID,
		"user_id":     order.UserID,
		"product_id":  order.ProductID,
		"payment_id":  paymentID,
		"total_price": order.TotalPrice,
		"quantity":    order.Quantity,
	})

	return paymentURL, paymentID, nil
}

func (uc *UseCase) CheckPayment(ctx context.Context, paymentID string) (string, error) {
	status, err := uc.paymentService.CheckStatus(ctx, paymentID)
	if err != nil {
		uc.logger.Error("failed to check payment", map[string]any{
			"payment_id": paymentID,
			"error":      err.Error(),
		})
		return "", err
	}

	uc.logger.Info("payment status checked", map[string]any{
		"payment_id": paymentID,
		"status":     status,
	})

	return status, nil
}
