package integration

import (
	"bytes"
	"encoding/json"
	"go-api/pkg/domain/accounts"
	"go-api/pkg/shared/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/accounts", func(c *fiber.Ctx) error {
		var account models.Account
		if err := c.BodyParser(&account); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		db.Create(&account)
		return c.Status(fiber.StatusCreated).JSON(account)
	})
}

func TestCreateAccountIntegration(t *testing.T) {
	app := fiber.New()
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("Could not set up database: %v", err)
	}

	setupRoutes(app, db)

	reqBody := []byte(`{"document": "26273128049"}`)
	req := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Error on app.Test: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", resp.StatusCode)
	}

	var account accounts.AddAccountRequest
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if account.Document != "26273128049" {
		t.Fatalf("Expected document to be '26273128049', got '%s'", account.Document)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
