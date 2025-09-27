package handlers

import (
	"gopkg.in/telebot.v4"
)

func (h *Handlers) AddProduct(c telebot.Context) error {
	tgID := c.Sender().ID

	addProductStates[tgID] = &AddProductState{Step: 1}

	return c.Send("Введите название товара:")
}
