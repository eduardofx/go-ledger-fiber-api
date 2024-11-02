package transactions

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{DB: db}

	routes := app.Group("/transactions")
	routes.Post("/", h.CreateTransaction)
}
