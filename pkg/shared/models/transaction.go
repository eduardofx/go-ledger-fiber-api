package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Amount         float64   `json:"amount"`
	AccountId      uuid.UUID `json:"account_id"`
	OperationId    uuid.UUID `json:"operation_id"`
	IdempotencyKey string    `gorm:"uniqueIndex;not null;index:idx_idempotency_key"`
	CreatedAt      time.Time
}
