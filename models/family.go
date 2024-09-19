package models

import (
	"time"

	"github.com/google/uuid"
)

type Family struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"size:255;not null"`
	IsActive  bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Family) TableName() string {
	return "family"
}
