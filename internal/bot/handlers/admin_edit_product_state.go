package handlers

import (
	"app/internal/dto"
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

type EditProductState struct {
	Step        int
	ProductID   int64
	Name        string
	Description string
	Price       float64
	Photos      []string
}

var editProductStates = make(map[int64]*EditProductState)

func (h *Handlers) HandleEditProductInput(c telebot.Context) error {
	tgID := c.Sender().ID
	state, ok := editProductStates[tgID]
	if !ok {
		return nil
	}

	switch state.Step {
	case 1:
		if c.Text() != "" {
			state.Name = c.Text()
		}
		state.Step = 2
		return c.Send(fmt.Sprintf("Введите новое описание (или оставьте старое: %s):", state.Description))

	case 2:
		if c.Text() != "" {
			state.Description = c.Text()
		}
		state.Step = 3
		return c.Send(fmt.Sprintf("Введите новую цену (или оставьте старую: %.2f):", state.Price))

	case 3:
		if c.Text() != "" {
			price, err := strconv.ParseFloat(c.Text(), 64)
			if err != nil {
				return c.Send("❌ Неверная цена, попробуйте ещё раз:")
			}
			state.Price = price
		}
		state.Step = 4
		return c.Send("Отправьте новое фото (или пропустите этот шаг):")

	case 4:
		if c.Message().Photo != nil {
			fileID := c.Message().Photo.FileID
			state.Photos = []string{fileID}
		}
		product := &dto.Product{
			ID:          state.ProductID,
			Name:        state.Name,
			Description: state.Description,
			Price:       state.Price,
			Photos:      state.Photos,
		}
		_, err := h.Uc.ProductUpdate(context.TODO(), product)
		if err != nil {
			return c.Send("❌ Ошибка при обновлении товара")
		}

		delete(editProductStates, tgID)
		return c.Send(fmt.Sprintf("✅ Товар #%d успешно обновлён!", product.ID))
	}

	return nil
}
