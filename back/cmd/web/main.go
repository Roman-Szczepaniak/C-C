package main

import (
	"fmt"
	"log"
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
	// database.DB.Migrator().DropTable(
	// 	&models.User{},
	// 	&models.Party{},
	// 	&models.History{},
	// 	&models.Card{},
	// 	&models.Action{},
	// 	&models.Monster{},
	// 	&models.Player{},
	// 	&models.Encounter{},
	// )

	// Auto-migration des modèles
	database.DB.AutoMigrate(
		&models.User{},
		&models.Party{},
		&models.History{},
		&models.Card{},
		&models.Action{},
		&models.Monster{},
		&models.Player{},
		&models.Encounter{},
	)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Initialisation du routeur Gin
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Request-Count"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	uc := controllers.NewUserController(database.DB)
	cc := controllers.NewCardController(database.DB)
	ac := controllers.NewActionController(database.DB)
	mc := controllers.NewMonsterController(database.DB)
	plc := controllers.NewPlayerController(database.DB)
	prc := controllers.NewPartyController(database.DB)
	ec := controllers.NewEncounterController(database.DB)
	hc := controllers.NewHistoryController(database.DB)

	routes.UserRoutes(router, uc)
	routes.AuthRoutes(router, uc)
	routes.CardRoutes(router, cc)
	routes.ActionRoutes(router, ac)
	routes.MonsterRoutes(router, mc)
	routes.PlayerRoutes(router, plc)
	routes.PartyRoutes(router, prc)
	routes.EncounterRoutes(router, ec)
	routes.HistoryRoutes(router, hc)

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
