package repository

import (
	"gorm.io/gorm"
	"main_car_service/internal/models"
)

type Brand interface {
	FindAll() ([]models.Brand, error)
	Create(brand *models.Brand) error
}

type Repository struct {
	Brand
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Brand: NewBrandRepository(db),
	}
}
