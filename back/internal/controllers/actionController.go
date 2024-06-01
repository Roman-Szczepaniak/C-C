package controllers

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActionController struct {
	DB *gorm.DB
}

func NewActionController(db *gorm.DB) *ActionController {
	return &ActionController{DB: db}
}

func (ac *ActionController) GetActions(c *gin.Context) {
	var actions []models.Action
	if err := ac.DB.Preload("Cards").Find(&actions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch actions"})
		return
	}
	c.JSON(http.StatusOK, actions)
}

func (ac *ActionController) GetActionByID(c *gin.Context) {
	id := c.Param("id")
	var action models.Action
	if err := ac.DB.Preload("Cards").First(&action, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Action not found"})
		return
	}
	c.JSON(http.StatusOK, action)
}

func (ac *ActionController) CreateAction(c *gin.Context) {
	var newAction models.Action
	if err := c.BindJSON(&newAction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action data"})
		return
	}
	for _, card := range newAction.Cards {
		var existingCard models.Card
		if err := ac.DB.First(&existingCard, card.ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Card not found"})
			return
		}
	}
	if err := ac.DB.Create(&newAction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create action"})
		return
	}
	c.JSON(http.StatusCreated, newAction)
}

func (ac *ActionController) UpdateAction(c *gin.Context) {
	id := c.Param("id")
	var action models.Action
	if err := ac.DB.First(&action, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Action not found"})
		return
	}
	if err := c.BindJSON(&action); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action data"})
		return
	}
	for _, card := range action.Cards {
		var existingCard models.Card
		if err := ac.DB.First(&existingCard, card.ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Card not found"})
			return
		}
	}
	if err := ac.DB.Save(&action).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update action"})
		return
	}
	if err := ac.DB.Preload("Cards").First(&action, action.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch action with cards"})
		return
	}
	c.JSON(http.StatusOK, action)
}

func (ac *ActionController) DeleteAction(c *gin.Context) {
	id := c.Param("id")
	var action models.Action
	if err := ac.DB.First(&action, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Action not found"})
		return
	}
	if err := ac.DB.Model(&action).Association("Cards").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associations"})
		return
	}
	if err := ac.DB.Delete(&action).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete action"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Action deleted successfully"})
}
