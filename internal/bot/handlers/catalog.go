package handlers

import (
	"context"
	"fmt"

	"gopkg.in/telebot.v4"
)

func (h *Handlers) Catalog(c telebot.Context) error {
	products, err := h.Uc.ProductList(context.TODO())
	if err != nil {
		return c.Send("❌ Ошибка при получении каталога")
	}
	if len(products) == 0 {
		return c.Send("Каталог пуст 📭")
	}

	markup := &telebot.ReplyMarkup{}
	var rows []telebot.Row
	for _, p := range products {
		btn := markup.Data(p.Name, "product", fmt.Sprint(p.ID))
		rows = append(rows, markup.Row(btn))
	}
	markup.Inline(rows...)

	return c.Send("📦 Выберите товар:", markup)
}
