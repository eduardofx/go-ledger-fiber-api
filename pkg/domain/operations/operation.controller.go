package operations

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{DB: db}

	routes := app.Group("/operations")
	routes.Get("/", h.GetOperationAll)
	routes.Post("/", h.CreateOperation)
	routes.Delete("/:id", h.DeleteOperation)
}
