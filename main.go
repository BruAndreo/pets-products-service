package main

import (
	"log"

	"github.com/bruandreo/pets-products-service/handlers"
	"github.com/bruandreo/pets-products-service/internal/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"message": "Hello World!",
		})
	})

	app.Post("/products", handlers.NewProduct)
	app.Get("/products", handlers.GetProducts)
	app.Get("/products/:id", handlers.GetProductById)
	app.Put("/products/:id", handlers.UpdateProduct)

	app.Listen(":3002")
}
