package operations

import (
	"go-api/pkg/shared/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteOperation(c *fiber.Ctx) error {
	id := c.Params("id")

	var operation models.Operation

	if result := h.DB.First(&operation, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&operation)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "The operation has been successfully deleted.",
		"id":      id,
	})
}
