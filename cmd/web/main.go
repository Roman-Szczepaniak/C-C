package main

import (
	"fmt"
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/internal/controllers"
	"github.com/Roman-Szczepaniak/C-C/internal/handlers"
	"github.com/Roman-Szczepaniak/C-C/internal/models"
	"github.com/Roman-Szczepaniak/C-C/internal/routes"
	"github.com/Roman-Szczepaniak/C-C/pkg/database"
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func main() {
	// Connection with the database in DBeaver
	database.ConnectDatabase()
	// Auto-migration des modèles
	database.DB.AutoMigrate(
		&models.User{},
		&models.Party{},
		&models.Card{},
		&models.History{},
		&models.Player{},
		&models.Monster{},
		&models.Encounter{},
		&models.Appear{},
	)

	// Initialisation du routeur Gin
	router := gin.Default()
	uc := controllers.NewUserController(database.DB)

	routes.UserRoutes(router, uc)
	routes.AuthRoutes(router, uc)
	// routes.EncounterRoutes(router)
	// routes.MonsterRoutes(router)
	// routes.PlayerRoutes(router)
	// routes.PartyRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})
	router.Run(port)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Listening at : <http://localhost:8080> - Server on port", port)
	http.ListenAndServe(port, nil)
}
