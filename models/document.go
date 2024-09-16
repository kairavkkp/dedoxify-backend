package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID               uint      `gorm:"primaryKey"`
	UUID             uuid.UUID `gorm:"type:uuid;not null"`
	FamilyID         uint      `gorm:"index"`
	Family           Family    `gorm:"foreignKey:FamilyID;constraint:OnDelete:SET NULL;"`
	UserID           uint      `gorm:"index"`
	User             Member    `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;"`
	Type             string    `gorm:"size:100;not null"`
	IsThumbnailReady bool      `gorm:"not null;default:false"`
	IsProcessed      bool      `gorm:"not null;default:false"`
	IsActive         bool      `gorm:"not null;default:true"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (Document) TableName() string {
	return "document"
}
