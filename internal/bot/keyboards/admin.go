package keyboards

import "gopkg.in/telebot.v4"

func AdminKeyboard() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnCatalog := m.Text("📦Каталог")
	btnCreate := m.Text("➕Добавить товар")

	m.Reply(m.Row(btnCatalog, btnCreate))
	return m
}
