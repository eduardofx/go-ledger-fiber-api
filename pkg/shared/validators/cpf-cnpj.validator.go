package validators

import (
	"github.com/klassmann/cpfcnpj"
)

func CpfCnpjValidator(document string) bool {
	switch len(document) {
	case 11:
		return cpfcnpj.ValidateCPF(document)
	case 14:
		return cpfcnpj.ValidateCNPJ(document)
	default:
		return false
	}
}
