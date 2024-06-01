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

// GetMonsters récupère la liste des monstres en fonction des filtres de requête fournis
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
	limit := 20 // Limite par défaut
	if l := c.Query("limit"); l != "" {
		if lInt, err := strconv.Atoi(l); err == nil {
			limit = lInt
		}
	}

	offset := 0 // Décalage par défaut
	if o := c.Query("offset"); o != "" {
		if oInt, err := strconv.Atoi(o); err == nil {
			offset = oInt
		}
	}

	// Application de limit et offset à la requête
	query = query.Limit(limit).Offset(offset)

	// Exécution de la requête
	if err := query.Preload("Encounters").Preload("Card").Find(&monsters).Error; err != nil {
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
	// Recherche du monstre avec préchargement des relations (encounters et card)
	if err := mc.DB.Preload("Encounters").Preload("Card").First(&monster, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monster not found"})
		return
	}
	c.JSON(http.StatusOK, monster)
}

// CreateMonster crée un nouveau monstre
func (mc *MonsterController) CreateMonster(c *gin.Context) {
	var newMonster models.Monster
	if err := c.BindJSON(&newMonster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid monster data"})
		return
	}
	// Vérification de l'existence de la carte associée si besoin (un monstre n'a pas forcément de carte)
	if newMonster.CardID != nil {
		var card models.Card
		if err := mc.DB.First(&card, newMonster.CardID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Card not found"})
			return
		}
	}
	// Vérifier si les encounters existent avant de créer des liens
	for _, encounter := range newMonster.Encounters {
		var existingEncounter models.Encounter
		if err := mc.DB.First(&existingEncounter, encounter.ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Encounter not found"})
			return
		}
	}
	// Création du monstre
	if err := mc.DB.Create(&newMonster).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create monster"})
		return
	}
	// Recherche du monstre créé avec préchargement des relations
	if err := mc.DB.Preload("Encounters").Preload("Card").First(&newMonster, newMonster.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monster with card"})
		return
	}
	c.JSON(http.StatusCreated, newMonster)
}

// UpdateMonster met à jour un monstre existant
func (mc *MonsterController) UpdateMonster(c *gin.Context) {
	id := c.Param("id")
	var monster models.Monster
	// Recherche du monstre à mettre à jour
	if err := mc.DB.First(&monster, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monster not found"})
		return
	}
	// Vérification du body/JSON envoyé
	if err := c.BindJSON(&monster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid monster data"})
		return
	}
	// Vérifier si les encounters existent avant d'update des liens
	for _, encounter := range monster.Encounters {
		var existingEncounter models.Encounter
		if err := mc.DB.First(&existingEncounter, encounter.ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Encounter not found"})
			return
		}
	}
	// Mise à jour du monstre
	if err := mc.DB.Save(&monster).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update monster"})
		return
	}
	// Recherche du monstre mis à jour avec préchargement des relations
	if err := mc.DB.Preload("Encounters").Preload("Card").First(&monster, monster.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monster with card and encounters"})
		return
	}
	c.JSON(http.StatusOK, monster)
}

func (mc *MonsterController) DeleteMonster(c *gin.Context) {
	id := c.Param("id")
	var monster models.Monster
	// Recherche du monstre à supprimer
	if err := mc.DB.Preload("Card").First(&monster, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monster not found"})
		return
	}
	// Impossibilité de supprimer un monstre avec une card associé
	if monster.CardID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete monster with associated card"})
		return
	}
	// Supprimer les associations sans supprimer les monsters et les encounters
	if err := mc.DB.Model(&monster).Association("Encounters").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associations"})
		return
	}
	// Suppression du monstre
	if err := mc.DB.Delete(&monster).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete monster"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Monster deleted successfully"})
}
