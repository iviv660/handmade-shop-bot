package handlers

import (
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

func (h *Handlers) Product(c telebot.Context) error {
	productID, _ := strconv.ParseInt(c.Data(), 10, 64)
	product, err := h.Uc.ProductGetByID(context.TODO(), productID)
	if err != nil {
		return err
	}

	// получаем пользователя (чтобы проверить роль)
	user, err := h.Uc.UserGetByTelegramID(context.TODO(), c.Sender().ID)
	if err != nil {
		return c.Send("❌ Ошибка: не удалось получить пользователя")
	}

	markup := &telebot.ReplyMarkup{}

	// общие кнопки
	btnBack := markup.Data("⬅️ Назад", "back")
	btnBuy := markup.Data("🛒 Купить", "buy", fmt.Sprint(product.ID))

	row := []telebot.Btn{btnBack, btnBuy}

	// если админ → добавляем доп. кнопки
	if user.Role == "admin" {
		btnEdit := markup.Data("✏️ Изменить", "edit", fmt.Sprint(product.ID))
		btnDelete := markup.Data("🗑 Удалить", "delete", fmt.Sprint(product.ID))
		row = append(row, btnEdit, btnDelete)
	}

	markup.Inline(markup.Row(row...))

	text := fmt.Sprintf("*%s*\n\n%s\n\n💰 Цена: %.2f ₽",
		product.Name, product.Description, product.Price)

	if len(product.Photos) > 0 {
		_, err := h.Bot.Send(c.Sender(), &telebot.Photo{
			File:    telebot.File{FileID: product.Photos[0]},
			Caption: text,
		}, markup, telebot.ModeMarkdown)
		return err
	}

	return c.Send(text, markup, telebot.ModeMarkdown)
}
