package controllers

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/internal/models"
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
	if err := cc.DB.Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cards"})
		return
	}
	c.JSON(http.StatusOK, cards)
}

func (cc *CardController) GetCardByID(c *gin.Context) {
	id := c.Param("id")
	var card models.Card
	if err := cc.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Card not found"})
	}
	c.JSON(http.StatusOK, card)
}
