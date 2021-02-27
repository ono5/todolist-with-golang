// main.go
package main

import (
	"fmt"
	"os"
	"todo/delivery"
	"todo/domain"
	"todo/repository"
	"todo/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// docker-composeファイルから環境変数を取得
var (
	schema         = "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	username       = os.Getenv("MYSQL_USER")
	password       = os.Getenv("MYSQL_PASSWORD")
	dbName         = os.Getenv("MYSQL_DATABASE")
	datasourceName = fmt.Sprintf(schema, username, password, dbName)
)

func connect() *gorm.DB {
	// mysqlへアクセス
	connection, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	// DBのスキーマーを作成
	connection.AutoMigrate(&domain.Todo{})

	return connection
}

func main() {
	// mysqlの接続
	db := connect()

	// mapのrepositoryをインスタンス化
	// tr := repository.NewSyncMapTodoRepository()

	// mysqlのrepositoryをインスタンス化
	tr := repository.NewTodoRepositoryMySQL(db)

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
