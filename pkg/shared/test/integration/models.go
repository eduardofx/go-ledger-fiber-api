package integration

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint   `gorm:"primaryKey"`
	Document  string `gorm:"unique"`
	Balance   float64
	CreatedAt time.Time
}

type Operation struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Type      Mode      `gorm:"type:operation_type;not null"`
	CreatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

type Transaction struct {
	Id             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Amount         float64   `json:"amount"`
	AccountId      uuid.UUID `json:"account_id"`
	OperationId    uuid.UUID `json:"operation_id"`
	IdempotencyKey string    `json:"idempotency_key"`
	CreatedAt      time.Time
}

type Mode string

const (
	Asset     Mode = "asset"
	Liability Mode = "liability"
)
