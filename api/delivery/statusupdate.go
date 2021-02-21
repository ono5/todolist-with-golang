// delivery/statusupdate.go
package delivery

import (
	"todo/domain"

	"github.com/gofiber/fiber/v2"
)

// Handlerを定義する
type todoStatusUpdateHandler struct {
	todoUseCase domain.TodoUsecase
}

func NewTodoStatusUpdateHandler(c *fiber.App, th domain.TodoUsecase) {
	handler := &todoStatusUpdateHandler{
		todoUseCase: th,
	}

	c.Post("/todo/statusupdate", handler.StatusUpdate)
}

func (h *todoStatusUpdateHandler) StatusUpdate(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	err := c.BodyParser(todo)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Unexpected Request. To check your ID",
		})
	}
	// UseCaseのStatusUpdateを呼びだす
	err = h.todoUseCase.StatusUpdate(todo.ID)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON("StatusUpdate OK")
}
