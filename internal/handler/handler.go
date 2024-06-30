package handler

import (
	"github.com/gin-gonic/gin"
	"main_car_service/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			brands := v1.Group("/brands")
			{
				brands.GET("/", h.getAllBrands)
				brands.POST("/", h.createBrand)
			}
		}
	}

	return router
}
