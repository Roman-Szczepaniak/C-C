package controllers

import (
	"net/http"

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

func (pc *PartyController) GetParties(c *gin.Context) {
	var parties []models.Party
	if err := pc.DB.Find(&parties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch parties"})
		return
	}
	c.JSON(http.StatusOK, parties)
}

func (pc *PartyController) GetPartyByID(c *gin.Context) {
	id := c.Param("id")
	var party models.Party
	if err := pc.DB.First(&party, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Party not found"})
		return
	}
	c.JSON(http.StatusOK, party)
}

func (pc *PartyController) CreateParty(c *gin.Context) {
	var newParty models.Party
	if err := c.BindJSON(&newParty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid party data"})
		return
	}
	if err := pc.DB.Create(&newParty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create party"})
		return
	}
	c.JSON(http.StatusCreated, newParty)
}

func (pc *PartyController) UpdateParty(c *gin.Context) {
	id := c.Param("id")
	var party models.Party
	if err := pc.DB.First(&party, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Party not found"})
		return
	}
	if err := c.BindJSON(&party); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid party data"})
		return
	}
	if err := pc.DB.Save(&party).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update party"})
		return
	}
	c.JSON(http.StatusOK, party)
}

func (pc *PartyController) DeleteParty(c *gin.Context) {
	id := c.Param("id")
	var party models.Party
	if err := pc.DB.First(&party, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Party not found"})
		return
	}
	if err := pc.DB.Delete(&party).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete party"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Party deleted successfully"})
}
