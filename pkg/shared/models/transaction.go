package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Amount         float64   `json:"amount"`
	AccountId      uint
	Account        Account `gorm:"constraint:OnUpdate:CASCADE"`
	OperationId    uint
	Operation      Operation `gorm:"constraint:OnUpdate:CASCADE"`
	IdempotencyKey string    `gorm:"uniqueIndex;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
