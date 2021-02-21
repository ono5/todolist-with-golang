// delivery/store.go
package delivery

import (
	"todo/domain"

	"github.com/gofiber/fiber/v2"
)

// Handlerを定義する
type todoStoreHandler struct {
	todoUseCase domain.TodoUsecase
}

func NewTodoStoreHandler(c *fiber.App, th domain.TodoUsecase) {
	handler := &todoStoreHandler{
		todoUseCase: th,
	}

	c.Post("/todo/store", handler.Store)
}

func (h *todoStoreHandler) Store(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	err := c.BodyParser(todo)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Unexpected Request. To check your Data",
		})
	}
	// UseCaseのStoreを呼びだす
	err = h.todoUseCase.Store(*todo)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON("Store OK")
}
