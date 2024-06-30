package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Brand Структура марки машины
type Brand struct {
	UID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"uid"`
	Title       string    `gorm:"type:varchar(30);uniqueIndex;index:idx_title,option:GIN" json:"title"`
	Description *string   `gorm:"type:text" json:"description"`
	PhotoURL    *string   `gorm:"type:text" json:"photo_url"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Brand{})
}
