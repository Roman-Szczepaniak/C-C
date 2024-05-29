package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
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
	query := mc.DB

	// Ajout des filtres
	if alignment := c.Query("alignment"); alignment != "" {
		query = query.Where("LOWER(alignment) = LOWER(?)", alignment)
	}
	if size := c.Query("size"); size != "" {
		query = query.Where("LOWER(size) = LOWER(?)", size)
	}
	if monsterType := c.Query("type"); monsterType != "" {
		query = query.Where("LOWER(type) = LOWER(?)", monsterType)
	}
	if environment := c.Query("environment"); environment != "" {
		query = query.Where("LOWER(environment) = LOWER(?)", environment)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("LOWER(status) = LOWER(?)", status)
	}

	// Ajout de la pagination
	limit := 20 // Default limit
	if l := c.Query("limit"); l != "" {
		if lInt, err := strconv.Atoi(l); err == nil {
			limit = lInt
		}
	}

	offset := 0 // Default offset
	if o := c.Query("offset"); o != "" {
		if oInt, err := strconv.Atoi(o); err == nil {
			offset = oInt
		}
	}

	// Application de limit et offset à la requête
	query = query.Limit(limit).Offset(offset)

	// Exécution de la requête
	if err := query.Find(&monsters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monsters"})
		return
	}

	// Récupération du nombre total de monstres pour la pagination
	var total int64
	mc.DB.Model(&models.Monster{}).Count(&total)

	// Ajout des informations de pagination aux en-têtes de la réponse (optionnel)
	c.Header("X-Total-Count", fmt.Sprintf("%d", total))
	c.Header("X-Limit", fmt.Sprintf("%d", limit))
	c.Header("X-Offset", fmt.Sprintf("%d", offset))

	// Envoi de la réponse avec les monstres
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
