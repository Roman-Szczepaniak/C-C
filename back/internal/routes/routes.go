package routes

import (
	"github.com/Roman-Szczepaniak/C-C/back/internal/controllers"
	"github.com/Roman-Szczepaniak/C-C/back/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine, uc *controllers.UserController) {
	incomingRoutes.POST("users/signup", uc.Signup)
	incomingRoutes.POST("users/login", uc.Login)
}

func AuthRoutes(incomingRoutes *gin.Engine, uc *controllers.UserController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", uc.GetUsers)
	incomingRoutes.GET("/profile", uc.GetProfile)
	incomingRoutes.PUT("/users/:id", uc.UpdateUser)
	incomingRoutes.DELETE("/users/:id", uc.DeleteUser)
}

func EncounterRoutes(incomingRoutes *gin.Engine, ec *controllers.EncounterController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/encounters", ec.GetEncounters)
	incomingRoutes.GET("/encounters/:id", ec.GetEncounterByID)
	incomingRoutes.POST("/encounters", ec.CreateEncounter)
	incomingRoutes.PUT("/encounters/:id", ec.UpdateEncounter)
	incomingRoutes.DELETE("/encounters/:id", ec.DeleteEncounter)
}

func MonsterRoutes(incomingRoutes *gin.Engine, mc *controllers.MonsterController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/monsters", mc.GetMonsters)
	incomingRoutes.GET("/monsters/:id", mc.GetMonsterByID)
	incomingRoutes.POST("/monsters", mc.CreateMonster)
	incomingRoutes.PUT("/monsters/:id", mc.UpdateMonster)
	incomingRoutes.DELETE("/monsters/:id", mc.DeleteMonster)
}

func PlayerRoutes(incomingRoutes *gin.Engine, pc *controllers.PlayerController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/players", pc.GetPlayers)
	incomingRoutes.GET("/players/:id", pc.GetPlayerByID)
	incomingRoutes.POST("/players", pc.CreatePlayer)
	incomingRoutes.PUT("/players/:id", pc.UpdatePlayer)
	incomingRoutes.DELETE("/players/:id", pc.DeletePlayer)
}

func PartyRoutes(incomingRoutes *gin.Engine, pc *controllers.PartyController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/parties", pc.GetParties)
	incomingRoutes.GET("/parties/:id", pc.GetPartyByID)
	incomingRoutes.POST("/parties", pc.CreateParty)
	incomingRoutes.PUT("/parties/:id", pc.UpdateParty)
	incomingRoutes.DELETE("/parties/:id", pc.DeleteParty)
}

func HistoryRoutes(incomingRoutes *gin.Engine, hc *controllers.HistoryController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/histories", hc.GetHistories)
	incomingRoutes.GET("/histories/:id", hc.GetHistoryByID)
}

func CardRoutes(incomingRoutes *gin.Engine, cc *controllers.CardController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/cards", cc.GetCards)
	incomingRoutes.GET("/cards/:id", cc.GetCardByID)
}
