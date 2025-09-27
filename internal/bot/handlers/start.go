package handlers

import (
	"app/internal/bot/keyboards"
	"app/internal/service"
	"context"
	"fmt"

	"gopkg.in/telebot.v4"
)

type Handlers struct {
	Bot     *telebot.Bot
	Uc      *service.UseCase
	AdminId int64
}

func (h *Handlers) Start(c telebot.Context) error {
	tgID := c.Sender().ID
	username := c.Sender().Username
	_, err := h.Uc.UserRegister(context.TODO(), tgID, username)
	if err != nil {
		return c.Send("❌ Ошибка при регистрации пользователя")
	}
	user, err := h.Uc.UserGetByTelegramID(context.TODO(), tgID)
	if err != nil {
		return c.Send("❌ Ошибка при получении роли пользователя")
	}
	if user.Role == "user" {
		return c.Send(fmt.Sprintf("Привет, %s 👋 Я магазин-бот.\n\nНажми «📦Каталог», чтобы посмотреть товары.", user.Username),
			keyboards.CatalogKeyboard(),
		)
	} else {
		return c.Send(fmt.Sprintln("Привет админ 👋"),
			keyboards.AdminKeyboard(),
		)
	}
}
