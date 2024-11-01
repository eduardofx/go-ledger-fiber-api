package operations

import (
	"go-api/pkg/shared/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetOperationAll(c *fiber.Ctx) error {

	var operation models.Operation

	if result := h.DB.First(&operation); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&operation)
}
