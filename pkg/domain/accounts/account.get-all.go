package accounts

import (
	"go-api/pkg/shared/models"
	"go-api/pkg/shared/validators"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAccountAll(c *fiber.Ctx) error {
	var account []models.Account

	if result := h.DB.Find(&account); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	for i := range account {
		formattedBalance, err := validators.TransformValueDecimal(account[i].Balance)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, `{"message": "Converting decimal error"}`)
		}
		account[i].Balance = formattedBalance
	}

	return c.Status(fiber.StatusOK).JSON(&account)
}
