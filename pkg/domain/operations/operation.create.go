package operations

import (
	"fmt"
	"go-api/pkg/shared/models"
	"go-api/pkg/shared/validators"

	"github.com/gofiber/fiber/v2"
)

type OperationRequest struct {
	Name string      `json:"name"`
	Type models.Mode `json:"type"`
}

func (h handler) CreateOperation(c *fiber.Ctx) error {

	body := OperationRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if !validators.ValidateOperationType(string(body.Type)) {
		return fiber.NewError(fiber.StatusBadRequest, `{"message": "This operation does not exist, choose 'asset' or 'liability"}`)
	}

	operation := models.Operation{
		Name: body.Name,
		Type: body.Type,
	}

	if result := h.DB.First(&operation, "name = ?", body.Name); result.Error == nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(`{"message": "The name already exists: %s"}`, body.Name))
	}

	if result := h.DB.Create(&operation); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&operation)

}
