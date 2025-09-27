package keyboards

import tb "gopkg.in/telebot.v4"

func CatalogKeyboard() *tb.ReplyMarkup {
	m := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnCatalog := m.Text("📦Каталог")
	m.Reply(m.Row(btnCatalog))
	return m
}
