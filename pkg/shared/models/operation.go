package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mode string

const (
	Asset     Mode = "asset"
	Liability Mode = "liability"
)

type Operation struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Type      Mode      `gorm:"type:operation_type;not null"`
	CreatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}
