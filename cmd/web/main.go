package main

import (
	"fmt"
	"net/http"

	"github.com/Roman-Szczepaniak/C-C/internal/handlers"
	"github.com/Roman-Szczepaniak/C-C/internal/models"
	"github.com/Roman-Szczepaniak/C-C/pkg/database"
)

const port = ":8080"

func main() {
	// Connection with the database in DBeaver
	database.ConnectDatabase()
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

	/*
		router := gin.New()
		routes.UserRoutes(router)
		routes.AuthRoutes(router)
		routes.EncounterRoutes(router)
		routes.MonsterRoutes(router)
		routes.PlayerRoutes(router)
		routes.AccountRoutes(router)
		routes.PartyRoutes(router)

		router.GET("/api-1", func(c *gin.Context) {
			c.JSON(200, gin.H{"success": "Access granted for api-1"})
		})

		router.GET("/api-2", func(c *gin.Context) {
			c.JSON(200, gin.H{"success": "Access granted for api-2"})
		})
		router.Run(":" + port)
	*/

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Listening at : <http://localhost:8080> - Server on port", port)
	http.ListenAndServe(port, nil)
}
