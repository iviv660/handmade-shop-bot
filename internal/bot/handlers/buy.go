package handlers

import (
	"app/internal/dto"
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (h *Handlers) Buy(c telebot.Context) error {
	productID, _ := strconv.ParseInt(c.Data(), 10, 64)

	user, err := h.Uc.UserGetByTelegramID(context.TODO(), c.Sender().ID)
	if err != nil {
		return c.Send("❌ Ошибка: не удалось найти пользователя")
	}

	product, err := h.Uc.ProductGetByID(context.TODO(), productID)
	if err != nil {
		return c.Send("❌ Ошибка: товар не найден")
	}

	order := &dto.Order{
		UserID:     user.ID,
		ProductID:  product.ID,
		Quantity:   1,
		TotalPrice: product.Price,
	}

	created, err := h.Uc.OrderCreate(context.TODO(), order)
	if err != nil {
		return c.Send("❌ Ошибка при создании заказа")
	}

	url, payID, err := h.Uc.CreatePayment(context.TODO(), created)
	if err != nil {
		return c.Send("❌ Ошибка при создании платежа")
	}

	_ = h.Uc.OrderAttachPaymentID(context.TODO(), created.ID, payID)

	markup := &telebot.ReplyMarkup{}
	markup.Inline(
		markup.Row(markup.URL("💳 Оплатить", url)),
		markup.Row(markup.Data("🔄 Проверить оплату", "check", fmt.Sprint(created.ID))),
	)

	return c.Send(fmt.Sprintf("✅ Заказ №%d создан!\nСумма: %.2f ₽", created.ID, created.TotalPrice), markup)

}
