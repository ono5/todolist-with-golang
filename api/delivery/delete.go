// delivery/delete.go
package delivery

import (
	"math/rand"
	"time"
	"todo/domain"

	"github.com/gofiber/fiber/v2"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Handlerを定義する
type todoDeleteHandler struct {
	todoUseCase domain.TodoUsecase
}

func NewTodoDeleteHandler(c *fiber.App, th domain.TodoUsecase) {
	handler := &todoDeleteHandler{
		todoUseCase: th,
	}

	c.Post("/todo/delete", handler.Delete)
}

func (h *todoDeleteHandler) Delete(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	err := c.BodyParser(todo)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Unexpected Request. To check your Todo data",
		})
	}

	// UseCaseのDeleteを呼びだす
	err = h.todoUseCase.Delete(todo.ID)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON("Delete OK")
}
