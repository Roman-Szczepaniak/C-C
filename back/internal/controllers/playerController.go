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

func (plc *PlayerController) GetPlayers(c *gin.Context) {
	var players []models.Player
	if err := plc.DB.Preload("Party").Preload("User").Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch players"})
		return
	}
	c.JSON(http.StatusOK, players)
}

func (plc *PlayerController) GetPlayerByID(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	if err := plc.DB.Preload("Party").Preload("User").First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (plc *PlayerController) CreatePlayer(c *gin.Context) {
	var newPlayer models.Player
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player data"})
		return
	}
	var party models.Party
	if err := plc.DB.First(&party, newPlayer.PartyID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Party not found"})
		return
	}
	var user models.User
	if err := plc.DB.First(&user, newPlayer.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	if err := plc.DB.Create(&newPlayer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
		return
	}
	if err := plc.DB.Preload("Party").Preload("User").First(&newPlayer, newPlayer.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch player with party and card"})
		return
	}
	c.JSON(http.StatusCreated, newPlayer)
}

func (plc *PlayerController) UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	if err := plc.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	if err := c.BindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player data"})
		return
	}
	if err := plc.DB.Save(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player"})
		return
	}
	if err := plc.DB.Preload("Party").Preload("User").First(&player, player.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch player with party and card"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (plc *PlayerController) DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	if err := plc.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	if err := plc.DB.Delete(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}
