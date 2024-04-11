package controllers

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MonsterController struct {
	DB *gorm.DB
}

func NewMonsterController(db *gorm.DB) *MonsterController {
	return &MonsterController{DB: db}
}

func (mc *MonsterController) GetMonsters(c *gin.Context) {
	var monsters []models.Monster
	if err := mc.DB.Find(&monsters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monsters"})
		return
	}
	c.JSON(http.StatusOK, monsters)
}

func (mc *MonsterController) GetMonsterByID(c *gin.Context) {
	id := c.Param("id")
	var monster models.Monster
	if err := mc.DB.First(&monster, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monster not found"})
		return
	}
	c.JSON(http.StatusOK, monster)
}

func (mc *MonsterController) CreateMonster(c *gin.Context) {
	var newMonster models.Monster
	if err := c.BindJSON(&newMonster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid monster data"})
		return
	}
	if err := mc.DB.Create(&newMonster).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create monster"})
		return
	}
	c.JSON(http.StatusCreated, newMonster)
}

func (mc *MonsterController) UpdateMonster(c *gin.Context) {
	id := c.Param("id")
	var monster models.Monster
	if err := mc.DB.First(&monster, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monster not found"})
		return
	}
	if err := c.BindJSON(&monster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid monster data"})
		return
	}
	if err := mc.DB.Save(&monster).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update monster"})
		return
	}
	c.JSON(http.StatusOK, monster)
}

func (mc *MonsterController) DeleteMonster(c *gin.Context) {
	id := c.Param("id")
	var monster models.Monster
	if err := mc.DB.First(&monster, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monster not found"})
		return
	}
	if err := mc.DB.Delete(&monster).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete monster"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Monster deleted successfully"})
}
