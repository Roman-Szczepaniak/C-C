package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAPIRoutes(t *testing.T) {
	// Initialiser Gin en mode Test
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Ajouter les routes que vous souhaitez tester
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	// Créer une requête HTTP pour tester la route
	req, _ := http.NewRequest(http.MethodGet, "/api-1", nil)
	resp := httptest.NewRecorder()

	// Exécuter la requête via le routeur
	router.ServeHTTP(resp, req)

	// Vérifier la réponse
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Access granted for api-1")
}
