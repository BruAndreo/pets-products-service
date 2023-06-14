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

func GetProducts(c *fiber.Ctx) error {
	var products []domain.Product

	database.Database.Find(&products)
	return c.Status(fiber.StatusOK).JSON(products)
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id param",
		})
	}

	var product domain.Product

	result := database.Database.Find(&product, id)
	if result.RowsAffected < 1 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id param",
		})
	}

	product := domain.Product{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product",
		})
	}

	result := database.Database.Where("id = ?", id).Updates(&product)

	if result.RowsAffected <= 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Product not updated",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product Updated",
	})
}

func RemoveProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id param",
		})
	}

	product := domain.Product{}

	result := database.Database.Delete(&product, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusOK)
}
