package keyboards

import "gopkg.in/telebot.v4"

func AdminKeyboard() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	btnCatalog := m.Text("📦Каталог")
	btnCreate := m.Text("➕Добавить товар")
	btnDelete := m.Text("🗑Удалить товар")
	btnUpdate := m.Text("✏️Обновить товар")

	m.Reply(m.Row(btnCatalog, btnCreate, btnDelete, btnUpdate))
	return m
}
