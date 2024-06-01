package controllers

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
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
	if err := hc.DB.Preload("User").Find(&histories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch histories"})
		return
	}
	c.JSON(http.StatusOK, histories)
}

func (hc *HistoryController) GetHistoryByID(c *gin.Context) {
	id := c.Param("id")
	var history models.History
	if err := hc.DB.Preload("User").First(&history, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "History not found"})
		return
	}
	c.JSON(http.StatusOK, history)
}

func (hc *HistoryController) CreateHistory(c *gin.Context) {
	var newHistory models.History
	if err := c.BindJSON(&newHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid history data"})
		return
	}
	var history models.History
	if err := hc.DB.First(&history, newHistory.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	if err := hc.DB.Create(&newHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create history"})
		return
	}
	if err := hc.DB.Preload("User").First(&newHistory, newHistory.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch history with user"})
		return
	}
	c.JSON(http.StatusCreated, newHistory)
}

func (hc *HistoryController) UpdateHistory(c *gin.Context) {
	id := c.Param("id")
	var history models.History
	if err := hc.DB.First(&history, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "History not found"})
		return
	}
	if err := c.BindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid history data"})
		return
	}
	if err := hc.DB.Save(&history).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update history"})
		return
	}
	if err := hc.DB.Preload("User").First(&history, history.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch history with user"})
		return
	}
	c.JSON(http.StatusOK, history)
}

func (hc *HistoryController) DeleteHistory(c *gin.Context) {
	id := c.Param("id")
	var history models.History
	if err := hc.DB.First(&history, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "History not found"})
		return
	}
	if err := hc.DB.Delete(&history).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete history"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "History deleted successfully"})
}
