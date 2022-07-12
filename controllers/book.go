package controllers

import (
	"net/http"

	"github.com/Duelana-Team/duelana-backend-v1/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &books)
}

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/books")

	routes.GET("/", h.GetBooks)
}
