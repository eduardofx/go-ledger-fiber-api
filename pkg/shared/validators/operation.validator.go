package validators

import "go-api/pkg/shared/models"

func ValidateOperationType(value string) bool {

	if value != string(models.Asset) && value != string(models.Liability) {
		return false
	}
	return true
}
