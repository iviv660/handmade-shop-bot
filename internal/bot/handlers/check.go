package handlers

import (
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (h *Handlers) CheckPayment(c telebot.Context) error {
	orderID, _ := strconv.ParseInt(c.Data(), 10, 64)

	order, err := h.Uc.OrderGetByID(context.TODO(), orderID)
	if err != nil {
		return c.Send("❌ Ошибка: заказ не найден")
	}

	if order.PaymentID == "" {
		return c.Send("⚠️ Для этого заказа нет платежа")
	}

	status, err := h.Uc.CheckPayment(context.TODO(), order.PaymentID)
	if err != nil {
		return c.Send("❌ Ошибка при проверке платежа")
	}

	if status == "succeeded" {
		_ = h.Uc.OrderUpdateStatus(context.TODO(), order.ID, "paid")

		// уведомляем покупателя
		_ = c.Send(fmt.Sprintf("✅ Заказ №%d оплачен! Мы скоро свяжемся для доставки 🚚", order.ID))

		// ищем покупателя в БД
		user, err := h.Uc.UserGetByID(context.TODO(), order.UserID)
		if err == nil {
			var chatLink string
			if user.Username != "" {
				chatLink = fmt.Sprintf("https://t.me/%s", user.Username)
			} else {
				chatLink = fmt.Sprintf("tg://user?id=%d", user.TelegramID)
			}

			adminID := int64(h.AdminId)
			_, _ = h.Bot.Send(&telebot.User{ID: adminID},
				fmt.Sprintf("💰 Заказ №%d оплачен.\nПокупатель: %s\nСумма: %.2f ₽\nСсылка: %s",
					order.ID, user.Username, order.TotalPrice, chatLink))
		}

		return nil
	}

	return c.Send(fmt.Sprintf("⚠️ Заказ №%d пока не оплачен. Статус: %s", order.ID, status))
}
