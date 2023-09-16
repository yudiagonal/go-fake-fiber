package main

import (
	"go-fake-fiber/models"
	"math/rand"
	"net/http"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/go-fake-fiber"), &gorm.Config{})
	if err != nil {
		panic("could not connect to database")
	}
	db.AutoMigrate(&models.Item{})
	return db

}

func main() {
	var db = ConnectionDB()

	app := fiber.New()
	app.Use(cors.New())

	app.Post("/api/item/create", func(c *fiber.Ctx) error {
		for i := 0; i < 100; i++ {
			db.Create(&models.Item{
				Name:        faker.Word(),
				Descryption: faker.Paragraph(),
				Price:       rand.Intn(140) + 10,
			})
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "success",
		})
	})

	app.Get("/api/item/all", func(c *fiber.Ctx) error {
		var items []models.Item
		db.Find(&items)
		return c.Status(200).JSON(items)
	})

	app.Listen(":8000")
}
