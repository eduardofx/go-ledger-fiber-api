package accounts

import (
	"go-api/pkg/shared/models"
	"go-api/pkg/shared/validators"

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

	balance, _ := validators.TransformValueDecimal(account.Balance)
	account.Balance = balance

	return c.Status(fiber.StatusOK).JSON(&account)
}
