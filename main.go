package main

import (
	"github.com/bruandreo/pets-products-service/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"message": "Hello World!",
		})
	})

	app.Post("/products", handlers.NewProduct)

	app.Listen(":3002")
}
