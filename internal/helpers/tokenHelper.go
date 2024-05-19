package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Roman-Szczepaniak/C-C/internal/models"
	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// SignedDetails contient les détails signés
type SignedDetails struct {
	Email              string
	FirstName          string
	LastName           string
	jwt.StandardClaims // Incorporation de la structure StandardClaims de jwt-go pour les informations de base sur le token
}

// Clé secrète utilisée pour signer les tokens JWT.
var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens génère un jeton d'accès et un jeton de rafraîchissement
func GenerateAllTokens(email string, firstName string, lastName string) (signedToken string, signedRefreshToken string, err error) {
	// Création des claims pour le token d'accès
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	// Création des claims pour le token de rafraîchissement
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	// Génération des tokens JWT signés
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

// ValidateToken valide le token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("le token est invalide")
		msg = err.Error()
		return
	}

	if claims.StandardClaims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("le token a expiré")
		msg = err.Error()
		return
	}
	return claims, msg
}

// UpdateAllTokens met à jour les tokens
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId uint, db *gorm.DB) {
	var user models.User
	if err := db.First(&user, userId).Error; err != nil {
		log.Panic(err)
		return
	}

	user.Token = &signedToken
	user.RefreshToken = &signedRefreshToken
	user.UpdatedAt = time.Now()

	if err := db.Save(&user).Error; err != nil {
		log.Panic(err)
		return
	}
}
