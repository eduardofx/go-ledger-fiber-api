package integration

import (
	"bytes"
	"encoding/json"
	"go-api/pkg/domain/transactions"
	"go-api/pkg/shared/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func setupTransactionRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/transactions", func(c *fiber.Ctx) error {
		var transaction Transaction
		if err := c.BodyParser(&transaction); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		db.Create(&transaction)
		return c.Status(fiber.StatusCreated).JSON(transaction)
	})
}

func TestCreateTransactionIntegration(t *testing.T) {
	app := fiber.New()
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("Could not set up database: %v", err)
	}

	setupTransactionRoutes(app, db)

	var account models.Account
	if err := db.First(&account).Error; err == gorm.ErrRecordNotFound {
		account = models.Account{
			Id:       uuid.New(),
			Document: "26273128049",
			Balance:  500,
		}
		if err := db.Create(&account).Error; err != nil {
			t.Fatalf("Erro ao criar a conta: %v", err)
		}
	}
	db.Create(&account)

	var operation Operation
	if err := db.First(&operation).Error; err == gorm.ErrRecordNotFound {
		operation = Operation{
			Id:   uuid.New(),
			Type: Asset,
			Name: uuid.NewString(),
		}
		if err := db.Create(&operation).Error; err != nil {
			t.Fatalf("Error on creating operation: %v", err)
		}
	}
	db.Create(&operation)

	transactionRequest := transactions.TransactionRequest{
		Amount:         10,
		IdempotencyKey: uuid.NewString(),
		AccountId:      account.Id,
		OperationId:    operation.Id,
	}

	reqBody, _ := json.Marshal(transactionRequest)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)

	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Código de status esperado 201, mas recebeu %d", resp.StatusCode)
	}

	var transaction models.Transaction
	if err := json.NewDecoder(resp.Body).Decode(&transaction); err != nil {
		t.Fatalf("Erro ao decodificar a resposta: %v", err)
	}

	if transaction.Amount != transactionRequest.Amount {
		t.Errorf("Valor esperado %.2f, mas recebeu %.2f", transactionRequest.Amount, transaction.Amount)
	}

	if transaction.AccountId != account.Id {
		t.Errorf("AccountId esperado %v, mas recebeu %v", account.Id, transaction.AccountId)
	}
}
