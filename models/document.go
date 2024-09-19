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
	MemberID         uint      `gorm:"index"`
	Member           Member    `gorm:"foreignKey:MemberID;constraint:OnDelete:SET NULL;"`
	Category         string    `gorm:"size:100;not null"`
	IsThumbnailReady bool      `gorm:"not null;default:false"`
	IsProcessed      bool      `gorm:"not null;default:false"`
	IsActive         bool      `gorm:"not null;default:true"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

func (Document) TableName() string {
	return "document"
}
