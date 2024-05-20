package controllers

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerController struct {
	DB *gorm.DB
}

func NewPlayerController(db *gorm.DB) *PlayerController {
	return &PlayerController{DB: db}
}

func (pc *PlayerController) GetPlayers(c *gin.Context) {
	var players []models.Player
	if err := pc.DB.Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch players"})
		return
	}
	c.JSON(http.StatusOK, players)
}

func (pc *PlayerController) GetPlayerByID(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	if err := pc.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (pc *PlayerController) CreatePlayer(c *gin.Context) {
	var newPlayer models.Player
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player data"})
		return
	}
	if err := pc.DB.Create(&newPlayer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
		return
	}
	c.JSON(http.StatusCreated, newPlayer)
}

func (pc *PlayerController) UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	if err := pc.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	if err := c.BindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player data"})
		return
	}
	if err := pc.DB.Save(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (pc *PlayerController) DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	if err := pc.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	if err := pc.DB.Delete(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}
