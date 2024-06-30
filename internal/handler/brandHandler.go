package handler

import (
	"github.com/gin-gonic/gin"
	"main_car_service/internal/models"
	"net/http"
)

func (h *Handler) getAllBrands(c *gin.Context) {
	brands, err := h.services.Brand.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (h *Handler) createBrand(c *gin.Context) {
	var brand models.Brand
	if err := c.BindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Brand.Create(&brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, brand)
}
