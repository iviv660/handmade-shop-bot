package handlers

import (
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (h *Handlers) EditProduct(c telebot.Context) error {
	productID, _ := strconv.ParseInt(c.Data(), 10, 64)

	product, err := h.Uc.ProductGetByID(context.TODO(), productID)
	if err != nil {
		return c.Send("❌ Ошибка: товар не найден")
	}

	tgID := c.Sender().ID
	editProductStates[tgID] = &EditProductState{
		Step:        1,
		ProductID:   product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Photos:      product.Photos,
	}

	return c.Send(fmt.Sprintf("✏️ Редактирование товара #%d.\nВведите новое название (или оставьте старое: %s):",
		product.ID, product.Name))
}
