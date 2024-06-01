package controllers

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EncounterController struct {
	DB *gorm.DB
}

func NewEncounterController(db *gorm.DB) *EncounterController {
	return &EncounterController{DB: db}
}

func (ec *EncounterController) GetEncounters(c *gin.Context) {
	var encounters []models.Encounter
	if err := ec.DB.Preload("Monsters").Find(&encounters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch encounters"})
		return
	}
	c.JSON(http.StatusOK, encounters)
}

func (ec *EncounterController) GetEncounterByID(c *gin.Context) {
	id := c.Param("id")
	var encounter models.Encounter
	if err := ec.DB.Preload("Monsters").First(&encounter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Encounter not found"})
		return
	}
	c.JSON(http.StatusOK, encounter)
}

func (ec *EncounterController) CreateEncounter(c *gin.Context) {
	var newEncounter models.Encounter
	if err := c.BindJSON(&newEncounter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encounter data"})
		return
	}
	var party models.Party
	if err := ec.DB.First(&party, newEncounter.PartyID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Party not found"})
		return
	}
	for _, monster := range newEncounter.Monsters {
		var existingMonster models.Monster
		if err := ec.DB.First(&existingMonster, monster.ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Monster not found"})
			return
		}
	}
	if err := ec.DB.Create(&newEncounter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create encounter"})
		return
	}
	if err := ec.DB.Preload("Monsters").First(&newEncounter, newEncounter.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch encounter with monsters"})
		return
	}
	c.JSON(http.StatusCreated, newEncounter)
}

func (ec *EncounterController) UpdateEncounter(c *gin.Context) {
	id := c.Param("id")
	var encounter models.Encounter
	if err := ec.DB.First(&encounter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Encounter not found"})
		return
	}
	if err := c.BindJSON(&encounter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encounter data"})
		return
	}
	for _, monster := range encounter.Monsters {
		var existingMonster models.Monster
		if err := ec.DB.First(&existingMonster, monster.ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Monster not found"})
			return
		}
	}
	if err := ec.DB.Save(&encounter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update encounter"})
		return
	}
	if err := ec.DB.Preload("Monsters").First(&encounter, encounter.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch encounter with monsters"})
		return
	}
	c.JSON(http.StatusOK, encounter)
}

func (ec *EncounterController) DeleteEncounter(c *gin.Context) {
	id := c.Param("id")
	var encounter models.Encounter
	if err := ec.DB.First(&encounter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Encounter not found"})
		return
	}
	if err := ec.DB.Model(&encounter).Association("Monsters").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associations"})
		return
	}
	if err := ec.DB.Delete(&encounter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete encounter"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Encounter deleted successfully"})
}
