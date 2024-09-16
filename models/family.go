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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Family) TableName() string {
	return "family"
}
