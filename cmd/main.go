package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Post("/users", CreateUser)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}

	//dsn := "host=localhost user=postgres password=200701 dbname=RicardinSalvaTudin port=5432 sslmode=disable"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//db.AutoMigrate(&domain.User{})
	//user := domain.User{Name: "João", Email: "joao@example.com"}
	//db.Create(&user)

}

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "success",
	})
}
