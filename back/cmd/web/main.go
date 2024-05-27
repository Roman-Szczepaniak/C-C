package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Roman-Szczepaniak/C-C/back/internal/controllers"
	"github.com/Roman-Szczepaniak/C-C/back/internal/handlers"
	"github.com/Roman-Szczepaniak/C-C/back/internal/models"
	"github.com/Roman-Szczepaniak/C-C/back/internal/routes"
	"github.com/Roman-Szczepaniak/C-C/back/pkg/database"
	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	uc := controllers.NewUserController(database.DB)
	mc := controllers.NewMonsterController(database.DB)
	ec := controllers.NewEncounterController(database.DB)
	plc := controllers.NewPlayerController(database.DB)
	prc := controllers.NewPartyController(database.DB)
	cc := controllers.NewCardController(database.DB)

	routes.UserRoutes(router, uc)
	routes.AuthRoutes(router, uc)
	routes.EncounterRoutes(router, ec)
	routes.MonsterRoutes(router, mc)
	routes.PlayerRoutes(router, plc)
	routes.PartyRoutes(router, prc)
	routes.CardRoutes(router, cc)

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
