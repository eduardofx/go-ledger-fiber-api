package accounts

import (
	"fmt"
	"go-api/pkg/shared/models"
	"go-api/pkg/shared/validators"

	"github.com/gofiber/fiber/v2"
	"github.com/klassmann/cpfcnpj"
)

type AddAccountRequest struct {
	Document string  `json:"document"`
	Balance  float64 `json:"balance"`
}

func (h handler) CreateAccount(c *fiber.Ctx) error {

	body := AddAccountRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	document := cpfcnpj.Clean(body.Document)

	if !validators.CpfCnpjValidator(document) {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(`{"message": "Invalid document: %s"}`, document))
	}

	var account models.Account

	if result := h.DB.First(&account, "document = ?", document); result.Error == nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(`{"message": "The document already exists","document":"%s"}`, document))
	}

	account.Document = document
	account.Balance = 0

	if result := h.DB.Create(&account); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&account)

}
