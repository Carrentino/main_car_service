package repository

import (
	"errors"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"main_car_service/internal/models"
)

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) Brand {
	return &BrandRepository{
		db: db,
	}
}

func (repo *BrandRepository) FindAll() ([]models.Brand, error) {
	var brands []models.Brand
	if err := repo.db.Find(&brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}

func (repo *BrandRepository) Create(brand *models.Brand) error {
	err := repo.db.Create(brand).Error
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return errors.New("unique constraint violated")
			}
		}
	}
	return err
}
