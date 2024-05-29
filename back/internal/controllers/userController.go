package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Roman-Szczepaniak/C-C/back/internal/helpers"
	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

var validate = validator.New()

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// GetUsers récupère tous les utilisateurs avec pagination
func (uc *UserController) GetUsers(c *gin.Context) {
	// Convertit les paramètres de requête "recordPerPage" et "page" en int
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// Récupère les utilisateurs avec pagination
	var users []models.User
	if err := uc.DB.Limit(recordPerPage).Offset((page - 1) * recordPerPage).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetProfile(c *gin.Context) {
	// Récupère l'email à partir du contexte (ajouté par le middleware d'authentification)
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := uc.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Signup crée un nouvel utilisateur (inscription)
func (uc *UserController) Signup(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	// Valide les données utilisateur
	validateErr := validate.Struct(newUser)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}

	// Vérifie si l'email existe déjà
	var count int64
	uc.DB.Model(&models.User{}).Where("email = ?", newUser.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Hash le mot de passe
	hashedPassword := HashPassword(newUser.HashPassword)
	newUser.HashPassword = hashedPassword
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	// Crée l'utilisateur dans la base de données
	if err := uc.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Génère les tokens
	token, refreshToken, _ := helpers.GenerateAllTokens(newUser.Email, newUser.FirstName, newUser.LastName)
	newUser.Token = &token
	newUser.RefreshToken = &refreshToken

	// Met à jour l'utilisateur avec les tokens
	if err := uc.DB.Save(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser})
}

// UpdateUser met à jour un utilisateur existant
func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := uc.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}
	if err := uc.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser supprime un utilisateur existant
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := uc.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := uc.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// HashPassword hashe un mot de passe
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// VerifyPassword vérifie un mot de passe
func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providePassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprintf("E-mail or Password is incorrect")
		check = false
	}
	return check, msg
}

// Login gère la connexion d'un utilisateur
func (uc *UserController) Login(c *gin.Context) {
	var user models.User
	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifie si l'email existe
	if err := uc.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email or password is incorrect"})
		return
	}

	// Vérifie si le mot de passe est valide
	passwordIsValid, msg := VerifyPassword(foundUser.HashPassword, user.HashPassword)
	if !passwordIsValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}

	// Génère les tokens
	token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.Email, foundUser.FirstName, foundUser.LastName)
	foundUser.Token = &token
	foundUser.RefreshToken = &refreshToken

	// Met à jour l'utilisateur avec les tokens
	if err := uc.DB.Save(&foundUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user tokens"})
		return
	}

	c.JSON(http.StatusOK, foundUser)
}
