package controllers

import (
	"net/http"
	"time"

	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PartyController struct {
	DB *gorm.DB
}

func NewPartyController(db *gorm.DB) *PartyController {
	return &PartyController{DB: db}
}

func (prc *PartyController) GetParties(c *gin.Context) {
	var parties []models.Party
	if err := prc.DB.Find(&parties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch parties"})
		return
	}
	c.JSON(http.StatusOK, parties)
}

func (prc *PartyController) GetPartyByID(c *gin.Context) {
	id := c.Param("id")
	var party models.Party
	if err := prc.DB.First(&party, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Party not found"})
		return
	}
	c.JSON(http.StatusOK, party)
}

func (prc *PartyController) CreateParty(c *gin.Context) {
	var newParty models.Party
	if err := c.BindJSON(&newParty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid party data"})
		return
	}
	newParty.CreatedAt = time.Now()
	newParty.UpdatedAt = time.Now()

	if err := prc.DB.Create(&newParty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create party"})
		return
	}
	c.JSON(http.StatusCreated, newParty)
}

func (prc *PartyController) UpdateParty(c *gin.Context) {
	id := c.Param("id")
	var party models.Party
	if err := prc.DB.First(&party, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Party not found"})
		return
	}
	if err := c.BindJSON(&party); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid party data"})
		return
	}
	party.UpdatedAt = time.Now()
	if err := prc.DB.Save(&party).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update party"})
		return
	}
	c.JSON(http.StatusOK, party)
}

func (prc *PartyController) DeleteParty(c *gin.Context) {
	id := c.Param("id")
	var party models.Party
	if err := prc.DB.First(&party, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Party not found"})
		return
	}
	if err := prc.DB.Delete(&party).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete party"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Party deleted successfully"})
}
