package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Document  string    `json:"document"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}
