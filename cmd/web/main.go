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
	database.ConnectDatabase()

	database.DB.AutoMigrate(
		&models.Account{},
		&models.Party{},
		&models.Card{},
		&models.History{},
		&models.Player{},
		&models.Monster{},
		&models.Encounter{},
		&models.Appear{},
	)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Listening at : <http://localhost:8080> - Server on port", port)
	http.ListenAndServe(port, nil)
}
