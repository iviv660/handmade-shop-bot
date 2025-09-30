package bot

import (
	"app/internal/bot/handlers"
	"app/internal/bot/middleware"
	"app/internal/service"

	"gopkg.in/telebot.v4"
)

func RegisterHandlers(b *telebot.Bot, uc *service.UseCase, adminID int64) {
	h := &handlers.Handlers{Bot: b, Uc: uc, AdminId: adminID}

	b.Handle("/start", middleware.InstrumentHandler("start_cmd", h.Start))
	b.Handle("📦Каталог", middleware.InstrumentHandler("catalog_cmd", h.Catalog))
	b.Handle("➕Добавить товар", middleware.InstrumentHandler("add_product_cmd", h.AddProduct))

	b.Handle(&telebot.Btn{Unique: "product"}, middleware.InstrumentHandler("btn_product", h.Product))
	b.Handle(&telebot.Btn{Unique: "back"}, middleware.InstrumentHandler("btn_back", h.Catalog))
	b.Handle(&telebot.Btn{Unique: "buy"}, middleware.InstrumentHandler("btn_buy", h.Buy))
	b.Handle(&telebot.Btn{Unique: "check"}, middleware.InstrumentHandler("btn_check", h.CheckPayment))
	b.Handle(&telebot.Btn{Unique: "edit"}, middleware.InstrumentHandler("btn_edit", h.EditProduct))
	b.Handle(&telebot.Btn{Unique: "delete"}, middleware.InstrumentHandler("btn_delete", h.DeleteProduct))

	b.Handle(telebot.OnText, middleware.InstrumentHandler("on_text", func(c telebot.Context) error {
		if err := h.HandleAdminInput(c); err != nil {
			return err
		}
		if err := h.HandleEditProductInput(c); err != nil {
			return err
		}
		return nil
	}))

	b.Handle(telebot.OnPhoto, middleware.InstrumentHandler("on_photo", func(c telebot.Context) error {
		if err := h.HandleAdminInput(c); err != nil {
			return err
		}
		if err := h.HandleEditProductInput(c); err != nil {
			return err
		}
		return nil
	}))
}
