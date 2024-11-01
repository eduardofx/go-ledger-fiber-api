package transactions

import (
	"go-api/pkg/shared/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetTransactionById(c *fiber.Ctx) error {
	var account models.Account

	if result := h.DB.First(&account, "id = ?", c.Params("id")); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&account)
}
