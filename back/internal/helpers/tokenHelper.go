package helpers

import (
	"log"
	"time"

	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// SignedDetails contient les détails signés
type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	jwt.RegisteredClaims
}

// Clé secrète utilisée pour signer les tokens JWT.
var SECRET_KEY string = "f6Mk_RirJVfMd0SoC01hII2WlorCxUZlrEA-DuMPIxg"

// GenerateAllTokens génère un jeton d'accès et un jeton de rafraîchissement
func GenerateAllTokens(email string, firstName string, lastName string) (signedToken string, signedRefreshToken string, err error) {
	// Création des claims pour le token d'accès
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Création des claims pour le token de rafraîchissement
	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
		},
	}

	// Génération des tokens JWT signés
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, nil
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
	if !ok || !token.Valid {
		msg = "le token est invalide"
		return
	}

	if claims.ExpiresAt.Before(time.Now()) {
		msg = "le token a expiré"
		return
	}

	return claims, ""
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
