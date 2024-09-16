package models

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"type:uuid;not null"`
	FirstName string    `gorm:"size:255;not null"`
	LastName  string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	FamilyID  uint      `gorm:"index"`
	Family    Family    `gorm:"foreignKey:FamilyID;constraint:OnDelete:SET NULL;"`
	IsActive  bool      `gorm:"not null;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Member) TableName() string {
	return "member"
}
