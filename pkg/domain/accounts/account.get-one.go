package accounts

import (
	"go-api/pkg/shared/models"

	"github.com/gofiber/fiber/v2"
	"github.com/klassmann/cpfcnpj"
)

func (h handler) GetAccountByDocument(c *fiber.Ctx) error {
	document := c.Params("document")
	document = cpfcnpj.Clean(document)

	var account models.Account

	if result := h.DB.First(&account, "document = ?", document); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&account)
}
