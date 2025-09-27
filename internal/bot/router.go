package bot

import (
	"app/internal/bot/handlers"
	"app/internal/service"

	"gopkg.in/telebot.v4"
)

func RegisterHandlers(b *telebot.Bot, uc *service.UseCase, adminID int64) {
	h := &handlers.Handlers{Bot: b, Uc: uc, AdminId: adminID}

	// команды
	b.Handle("/start", h.Start)
	b.Handle("📦Каталог", h.Catalog)
	b.Handle("➕Добавить товар", h.AddProduct)

	// inline кнопки
	b.Handle(&telebot.Btn{Unique: "product"}, h.Product)
	b.Handle(&telebot.Btn{Unique: "back"}, h.Catalog)
	b.Handle(&telebot.Btn{Unique: "buy"}, h.Buy)
	b.Handle(&telebot.Btn{Unique: "check"}, h.CheckPayment)
	b.Handle(&telebot.Btn{Unique: "edit"}, h.EditProduct)
	b.Handle(&telebot.Btn{Unique: "delete"}, h.DeleteProduct)

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		if err := h.HandleAdminInput(c); err != nil {
			return err
		}
		if err := h.HandleEditProductInput(c); err != nil {
			return err
		}
		return nil
	})
	b.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		if err := h.HandleAdminInput(c); err != nil {
			return err
		}
		if err := h.HandleEditProductInput(c); err != nil {
			return err
		}
		return nil
	})
}
