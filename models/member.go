package models

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"type:uuid;not null"`
	FirstName string    `gorm:"size:255"`
	LastName  string    `gorm:"size:255"`
	Email     string    `gorm:"size:255;unique;not null"`
	FamilyID  uint      `gorm:"index"`
	Family    Family    `gorm:"foreignKey:FamilyID;constraint:OnDelete:SET NULL;"`
	IsActive  bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Member) TableName() string {
	return "member"
}
