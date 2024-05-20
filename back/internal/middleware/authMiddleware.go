package middleware

import (
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/back/internal/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}

		// Validez le token
		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		// Ajoutez les informations du token dans le contexte de la requête
		c.Set("email", claims.Email)
		c.Set("firstname", claims.FirstName)
		c.Set("lastname", claims.LastName)

		c.Next() // Continuez avec le prochain gestionnaire
	}
}
