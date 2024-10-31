package accounts

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{DB: db}

	routes := app.Group("/accounts")
	routes.Get("", h.GetAccountAll)
	routes.Get("/:document", h.GetAccountByDocument)
	routes.Post("/", h.CreateAccount)
	routes.Delete("/:id", h.DeleteAccount)
}
