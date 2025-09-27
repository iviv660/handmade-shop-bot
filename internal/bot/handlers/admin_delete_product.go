package handlers

import (
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (h *Handlers) DeleteProduct(c telebot.Context) error {
	productID, _ := strconv.ParseInt(c.Data(), 10, 64)
	if err := h.Uc.ProductDelete(context.TODO(), productID); err != nil {
		return c.Send("❌ Ошибка при удалении товара")
	}
	return c.Send(fmt.Sprintf("🗑 Товар #%d удалён!", productID))
}
