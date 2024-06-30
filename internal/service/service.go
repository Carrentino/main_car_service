package service

import (
	"main_car_service/internal/models"
	"main_car_service/internal/repository"
)

type Brand interface {
	FindAll() ([]models.Brand, error)
	Create(brand *models.Brand) error
}

type Service struct {
	Brand
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Brand: &BrandService{repo: r.Brand},
	}

}
