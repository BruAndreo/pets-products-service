package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       int    `json:"value"`
}

func NewProduct(c *fiber.Ctx) error {
	product := Product{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}
