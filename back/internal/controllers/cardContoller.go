package controllers

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CardController struct {
	DB *gorm.DB
}

func NewCardController(db *gorm.DB) *CardController {
	return &CardController{DB: db}
}

func (cc *CardController) GetCards(c *gin.Context) {
	var cards []models.Card
	if err := cc.DB.Preload("Actions").Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cards"})
		return
	}
	c.JSON(http.StatusOK, cards)
}

func (cc *CardController) GetCardByID(c *gin.Context) {
	id := c.Param("id")
	var card models.Card
	if err := cc.DB.Preload("Actions").First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}
	c.JSON(http.StatusOK, card)
}

func (cc *CardController) CreateCard(c *gin.Context) {
	var newCard models.Card
	if err := c.BindJSON(&newCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card data"})
		return
	}
	if err := cc.DB.Create(&newCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create card"})
		return
	}
	c.JSON(http.StatusCreated, newCard)
}

func (cc *CardController) UpdateCard(c *gin.Context) {
	id := c.Param("id")
	var card models.Card
	if err := cc.DB.Preload("Actions").First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}
	if err := c.BindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card data"})
		return
	}
	if err := cc.DB.Save(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update card"})
		return
	}
	if err := cc.DB.Preload("Actions").First(&card, card.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch card with actions"})
		return
	}
	c.JSON(http.StatusOK, card)
}

func (cc *CardController) DeleteCard(c *gin.Context) {
	id := c.Param("id")
	var card models.Card
	if err := cc.DB.Preload("Actions").First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}
	// Supprimer les associations sans supprimer la carte
	if err := cc.DB.Model(&card).Association("Actions").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associations"})
		return
	}
	if err := cc.DB.Delete(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete card"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Card deleted successfully"})
}
