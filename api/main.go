// main.go
package main

import (
	"todo/delivery"
	"todo/repository"
	"todo/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// repositoryをインスタンス化
	tr := repository.NewSyncMapTodoRepository()
	// usecaseをインスタンス化
	tu := usecase.NewTodoUsecase(tr)

	// fiberをインスタンス化
	c := fiber.New()
	delivery.NewTodoAllGetHandler(c, tu)
	delivery.NewTodoDeleteHandler(c, tu)
	delivery.NewTodoStatusUpdateHandler(c, tu)
	delivery.NewTodoStoreHandler(c, tu)

	c.Listen(":80")
}
