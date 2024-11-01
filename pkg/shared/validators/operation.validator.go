package validators

import (
	"fmt"
	"go-api/pkg/shared/models"
	"math"
	"strconv"
)

func ValidateOperationType(value string) bool {

	if value != string(models.Asset) && value != string(models.Liability) {
		return false
	}
	return true
}

func TransformValueDecimal(balance float64) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf("%.2f", balance), 64)
}

func RoundToTwoDecimals(value float64) float64 {
	return math.Floor(value*100+0.000001) / 100
}
