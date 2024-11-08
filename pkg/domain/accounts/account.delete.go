package accounts

import (
	"go-api/pkg/shared/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteAccount(c *fiber.Ctx) error {
	id := c.Params("id")

	var account models.Account

	if result := h.DB.First(&account, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&account)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "The account has been successfully deleted.",
		"id":      id,
	})
}
