package controllers

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HistoryController struct {
	DB *gorm.DB
}

func NewHistoryController(db *gorm.DB) *HistoryController {
	return &HistoryController{DB: db}
}

func (hc *HistoryController) GetHistories(c *gin.Context) {
	var histories []models.History
	if err := hc.DB.Find(&histories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch histories"})
		return
	}
	c.JSON(http.StatusInternalServerError, histories)
}

func (hc *HistoryController) GetHistoryByID(c *gin.Context) {
	id := c.Param("id")
	var history models.History
	if err := hc.DB.First(&history, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "History not found"})
	}
	c.JSON(http.StatusOK, history)
}
