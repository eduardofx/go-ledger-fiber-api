package models

import "github.com/google/uuid"

type Account struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Document string    `json:"document"`
	Balance  float64   `json:"balance"`
}
