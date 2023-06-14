package handlers

import (
	"github.com/bruandreo/pets-products-service/domain"
	"github.com/bruandreo/pets-products-service/internal/database"
	"github.com/gofiber/fiber/v2"
)

func NewProduct(c *fiber.Ctx) error {
	product := domain.Product{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := database.Database.Create(product)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}
