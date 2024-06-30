package service

import (
	"main_car_service/internal/models"
	"main_car_service/internal/repository"
)

// BrandService Сервис для марок
type BrandService struct {
	repo repository.Brand
}

func (s *BrandService) FindAll() ([]models.Brand, error) {
	return s.repo.FindAll()
}

func (s *BrandService) Create(brand *models.Brand) error {
	return s.repo.Create(brand)
}
