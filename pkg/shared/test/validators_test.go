package test

import (
	"go-api/pkg/shared/validators"
	"testing"
)

func TestValidatingOperationType(t *testing.T) {
	result := validators.ValidateOperationType("teste")

	if result != false {
		t.Errorf("ValidateOperationType('g') = %t; want %t", result, false)
	}

	valid := validators.ValidateOperationType("asset")
	if valid != true {
		t.Errorf("ValidateOperationType('g') = %t; want %t", valid, true)
	}
}
