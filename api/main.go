// main.go
package main

import (
	"todo/delivery"
	"todo/repository"
	"todo/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// repositoryをインスタンス化
	tr := repository.NewSyncMapTodoRepository()
	// usecaseをインスタンス化
	tu := usecase.NewTodoUsecase(tr)

	// fiberをインスタンス化
	c := fiber.New()

	// CORSの設定
	c.Use(cors.New(cors.Config{
		// https://docs.gofiber.io/api/middleware/cors#config
		AllowCredentials: true,
	}))

	delivery.NewTodoAllGetHandler(c, tu)
	delivery.NewTodoDeleteHandler(c, tu)
	delivery.NewTodoStatusUpdateHandler(c, tu)
	delivery.NewTodoStoreHandler(c, tu)
	delivery.NewTodoSearchHandler(c, tu)

	c.Listen(":80")
}
