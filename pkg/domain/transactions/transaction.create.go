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

	tx := h.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var account models.Account
	result := tx.Raw("SELECT * FROM accounts WHERE id = ? FOR UPDATE", body.AccountId).Scan(&account)
	if result.Error != nil {
		tx.Rollback()
		if result.RowsAffected == 0 {
			return fiber.NewError(fiber.StatusNotFound, `{"message": "Account not found"}`)
		}
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	var operation models.Operation
	if err := tx.First(&operation, "id = ?", body.OperationId).Error; err != nil {
		tx.Rollback()
		return fiber.NewError(fiber.StatusNotFound, `{"message": "Operation not found"}`)
	}

	if operation.Type == models.Liability {
		if validators.RoundToTwoDecimals(account.Balance) < body.Amount {
			tx.Rollback()
			return fiber.NewError(fiber.StatusBadRequest, `{"message": "Insufficient balance"}`)
		}

		balance := account.Balance - body.Amount

		if err := tx.Model(&account).Where("id = ?", account.Id).Update("balance", balance).Error; err != nil {
			tx.Rollback()
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		body.Amount = body.Amount * -1
	}

	if operation.Type == models.Asset {
		balance := account.Balance + body.Amount

		if err := tx.Model(&account).Where("id = ?", account.Id).Update("balance", balance).Error; err != nil {
			tx.Rollback()
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
		tx.Rollback()
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&transaction)

}
