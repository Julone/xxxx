package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm-mysql/database"
	"gorm-mysql/models"
	"gorm-mysql/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allbooks", routes.AllBooks)
	app.Get("/book/:id", routes.GetBook)
	app.Post("/book", routes.AddBook)
	app.Put("/book/:id", routes.Update)
	app.Delete("/book/:id", routes.Delete)
	app.Get("/comments", routes.GetCommentsList)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)
	go func() {
		database.DBConn.AutoMigrate(&models.Book{}, &models.Page{}, &models.User{})
		database.DBConn.AutoMigrate(&models.Comments{})
	}()

	app.Use(cors.New())
	c := make(chan int, 1)
	go DoAdd(c, 12, 34, 2, 4)
	go DoAdd(c, 12, 34, 2, 40)
	fmt.Printf("%v, %v", <-c, <-c)
	safeRegisterApolloPortal()
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"asdfsadfas saf ads
	})

	log.Fatal(app.Listen(":3000"))
}
func safeRegisterApolloPortal() {
	defer func() {
		//fasdfasf
		// 处理 registerApolloPortal() 出现的 panic 错误
		r := recover()
		if (interface{})(r) != nil {
			log.Error("panic when register apollo portal to platform. err: %v", r)
		}
	}()

	// 处理 registerApolloPortal() 返回的 error 类型
	panic("asfas")
}

func DoAdd(c chan int, args ...int) {
	a := 0
	for _, arg := range args {
		a = a + arg
	}
	c <- a
}
