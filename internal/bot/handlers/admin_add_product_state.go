package handlers

import (
	"app/internal/dto"
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v4"
)

type AddProductState struct {
	Step        int
	Name        string
	Description string
	Price       float64
	Photo       string
}

var addProductStates = make(map[int64]*AddProductState)

func (h *Handlers) HandleAdminInput(c telebot.Context) error {
	tgID := c.Sender().ID
	state, ok := addProductStates[tgID]
	if !ok {
		return nil
	}

	switch state.Step {
	case 1:
		state.Name = c.Text()
		state.Step = 2
		return c.Send("Введите описание товара:")

	case 2:
		state.Description = c.Text()
		state.Step = 3
		return c.Send("Введите цену товара (например: 499.99):")

	case 3:
		price, err := strconv.ParseFloat(c.Text(), 64)
		if err != nil {
			return c.Send("❌ Неверная цена, попробуйте ещё раз:")
		}
		state.Price = price
		state.Step = 4
		return c.Send("Отправьте фото товара:")

	case 4:
		if c.Message().Photo != nil {
			fileID := c.Message().Photo.FileID
			state.Photo = fileID

			product := &dto.Product{
				Name:        state.Name,
				Description: state.Description,
				Price:       state.Price,
				PhotoID:     state.Photo,
			}

			_, err := h.Uc.ProductCreate(context.TODO(), product)
			if err != nil {
				return c.Send("❌ Ошибка при сохранении товара")
			}

			delete(addProductStates, tgID)
			return c.Send(fmt.Sprintf("✅ Товар \"%s\" успешно добавлен!", product.Name))
		}

		return c.Send("❌ Отправьте фото товара")

	}

	return nil
}
