package service

import (
	"app/internal/dto"
	"context"
)

const (
	OrderStatusPending   = "pending"
	OrderStatusPaid      = "paid"
	OrderStatusShipped   = "shipped"
	OrderStatusCancelled = "cancelled"
)

func (uc *UseCase) OrderCreate(ctx context.Context, order *dto.Order) (*dto.Order, error) {
	newOrder, err := uc.orderService.Create(ctx, order)
	if err != nil {
		uc.logger.Error("failed to create order", map[string]any{
			"user_id":     order.UserID,
			"product_id":  order.ProductID,
			"total_price": order.TotalPrice,
			"quantity":    order.Quantity,
			"error":       err.Error(),
		})
		return nil, err
	}

	uc.logger.Info("successfully created order", map[string]any{
		"order_id":    newOrder.ID,
		"user_id":     newOrder.UserID,
		"product_id":  newOrder.ProductID,
		"status":      newOrder.Status,
		"total_price": newOrder.TotalPrice,
		"quantity":    newOrder.Quantity,
	})

	return newOrder, nil
}

func (uc *UseCase) OrderListByUserID(ctx context.Context, userID int64) ([]*dto.Order, error) {
	orders, err := uc.orderService.ListByUserID(ctx, userID)
	if err != nil {
		uc.logger.Error("failed to list orders", map[string]any{
			"user_id": userID,
			"error":   err.Error(),
		})
		return nil, err
	}

	orderIDs := make([]int64, 0, len(orders))
	for _, o := range orders {
		orderIDs = append(orderIDs, o.ID)
	}

	uc.logger.Info("successfully fetched orders", map[string]any{
		"user_id":      userID,
		"orders_count": len(orders),
		"order_ids":    orderIDs,
	})

	return orders, nil
}

func (uc *UseCase) OrderUpdateStatus(ctx context.Context, orderID int64, status string) error {
	err := uc.orderService.UpdateStatus(ctx, orderID, status)
	if err != nil {
		uc.logger.Error("failed to update order status", map[string]any{
			"order_id":   orderID,
			"new_status": status,
			"error":      err.Error(),
		})
		return err
	}

	uc.logger.Info("order status updated successfully", map[string]any{
		"order_id":   orderID,
		"new_status": status,
	})

	return nil
}

func (uc *UseCase) OrderAttachPaymentID(ctx context.Context, orderID int64, paymentID string) error {
	err := uc.orderService.AttachPaymentID(ctx, orderID, paymentID)
	if err != nil {
		uc.logger.Error("failed to attach payment id to order", map[string]any{
			"order_id":   orderID,
			"payment_id": paymentID,
			"error":      err.Error(),
		})
		return err
	}

	uc.logger.Info("payment id attached to order successfully", map[string]any{
		"order_id":   orderID,
		"payment_id": paymentID,
	})

	return nil
}

func (uc *UseCase) OrderGetByID(ctx context.Context, orderID int64) (*dto.Order, error) {
	order, err := uc.orderService.GetByID(ctx, orderID)
	if err != nil {
		uc.logger.Error("failed to get order by id", map[string]any{
			"order_id": orderID,
			"error":    err.Error(),
		})
		return nil, err
	}
	uc.logger.Info("successfully fetched order by id", map[string]any{
		"order_id":    orderID,
		"user_id":     order.UserID,
		"product_id":  order.ProductID,
		"status":      order.Status,
		"quantity":    order.Quantity,
		"total_price": order.TotalPrice,
	})
	return order, nil
}
