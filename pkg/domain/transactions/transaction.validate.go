package transactions

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h handler) validateTransactionRequest(body TransactionRequest) error {

	if body.Amount < 0.01 {
		return fiber.NewError(fiber.StatusBadRequest, `{"message": "You must only pass values above 0.01"}`)
	}

	valueStr := strconv.FormatFloat(body.Amount, 'f', -1, 64)
	if strings.Contains(valueStr, ".") {
		parts := strings.Split(valueStr, ".")
		if len(parts) > 1 && len(parts[1]) > 2 {
			return fiber.NewError(fiber.StatusBadRequest, `{"message": "Amount must have at most two decimal places"}`)
		}
	}

	return nil
}
