package transactions

import (
	"go-api/pkg/shared/models"
	"go-api/pkg/shared/validators"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TransactionRequest struct {
	Amount         float64   `json:"amount"`
	IdempotencyKey string    `json:"idempotency_key"`
	AccountId      uuid.UUID `json:"account_id"`
	OperationId    uuid.UUID `json:"operation_id"`
}

func (h handler) CreateTransaction(c *fiber.Ctx) error {
	body := TransactionRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var checkIdempotencyKey models.Transaction
	if err := h.DB.First(&checkIdempotencyKey, "idempotency_key = ?", body.IdempotencyKey).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(checkIdempotencyKey)
	}

	if err := h.validateTransactionRequest(body); err != nil {
		return err
	}

	body.Amount = validators.RoundToTwoDecimals(body.Amount)

	var account models.Account
	if err := h.DB.First(&account, "id = ?", body.AccountId).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, `{"message": "Account not found"}`)
	}

	var operation models.Operation
	if err := h.DB.First(&operation, "id = ?", body.OperationId).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, `{"message": "Operation not found"}`)
	}

	if operation.Type == models.Liability {
		if validators.RoundToTwoDecimals(account.Balance) < body.Amount {
			return fiber.NewError(fiber.StatusNotFound, `{"message": "Insufficient balance"}`)
		}

		balance := account.Balance - body.Amount

		if err := h.DB.Model(&account).Where("id = ?", account.Id).Update("balance", balance).Error; err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		body.Amount = body.Amount * -1
	}

	if operation.Type == models.Asset {
		balance := account.Balance + body.Amount

		if err := h.DB.Model(&account).Where("id = ?", account.Id).Update("balance", balance).Error; err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	transaction := models.Transaction{
		Amount:         body.Amount,
		IdempotencyKey: body.IdempotencyKey,
		OperationId:    body.OperationId,
		AccountId:      body.AccountId,
	}

	if result := h.DB.Create(&transaction); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(&transaction)

}
