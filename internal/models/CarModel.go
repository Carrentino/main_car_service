package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CarModel struct {
	UID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"uid"`
	Brand           *Brand     `gorm:"foreignKey:BrandID;references:UID" json:"brand"`
	BrandID         *uuid.UUID `gorm:"type:uuid" json:"brand_id"`
	Title           string     `gorm:"type:varchar(100);index:idx_title,option:GIN" json:"title"`
	EngineCapacity  *float64   `gorm:"type:decimal(3,1)" json:"engine_capacity"`
	Drive           *string    `gorm:"type:varchar(3)" json:"drive"`
	Gearbox         *string    `gorm:"type:varchar(2)" json:"gearbox"`
	BodyType        *string    `gorm:"type:varchar(2)" json:"body_type"`
	TypeOfFuel      *string    `gorm:"type:varchar(2)" json:"type_of_fuel"`
	FuelConsumption *float64   `gorm:"type:decimal(3,1)" json:"fuel_consumption"`
	HP              *int       `gorm:"type:int" json:"hp"`
}

func (c *CarModel) BeforeSave(tx *gorm.DB) (err error) {
	if c.Drive != nil && !contains(DRIVE_CHOICES, *c.Drive) {
		return fmt.Errorf("invalid value for Drive: %s", *c.Drive)
	}
	if c.Gearbox != nil && !contains(GEARBOX_CHOICES, *c.Gearbox) {
		return fmt.Errorf("invalid value for Gearbox: %s", *c.Gearbox)
	}
	if c.BodyType != nil && !contains(BODY_TYPE_CHOICES, *c.BodyType) {
		return fmt.Errorf("invalid value for BodyType: %s", *c.BodyType)
	}
	if c.TypeOfFuel != nil && !contains(FUEL_TYPE_CHOICES, *c.TypeOfFuel) {
		return fmt.Errorf("invalid value for TypeOfFuel: %s", *c.TypeOfFuel)
	}
	return
}

// Функция для проверки, содержит ли слайс строк определенное значение
func contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}
