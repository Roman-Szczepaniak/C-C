package routes

import (
	"github.com/Roman-Szczepaniak/C-C/internal/controllers"
	"github.com/Roman-Szczepaniak/C-C/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine, uc *controllers.UserController) {
	incomingRoutes.POST("users/signup", uc.Signup)
	incomingRoutes.POST("users/login", uc.Login)
}

func AuthRoutes(incomingRoutes *gin.Engine, uc *controllers.UserController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", uc.GetUsers)
	incomingRoutes.GET("/users/:id", uc.GetUserByID)
	incomingRoutes.PUT("/users/:id", uc.UpdateUser)
	incomingRoutes.DELETE("/users/:id", uc.DeleteUser)
}

func EncounterRoutes(incomingRoutes *gin.Engine, ec *controllers.EncounterController) {
	incomingRoutes.GET("/encounters", ec.GetEncounters)
	incomingRoutes.GET("/encounters/:id", ec.GetEncounterByID)
	incomingRoutes.POST("/encounters", ec.CreateEncounter)
	incomingRoutes.PUT("/encounters/:id", ec.UpdateEncounter)
	incomingRoutes.DELETE("/encounters/:id", ec.DeleteEncounter)
}

func MonsterRoutes(incomingRoutes *gin.Engine, mc *controllers.MonsterController) {
	incomingRoutes.GET("/monsters", mc.GetMonsters)
	incomingRoutes.GET("/monsters/:id", mc.GetMonsterByID)
	incomingRoutes.POST("/monsters", mc.CreateMonster)
	incomingRoutes.PUT("/monsters/:id", mc.UpdateMonster)
	incomingRoutes.DELETE("/monsters/:id", mc.DeleteMonster)
}

func PlayerRoutes(incomingRoutes *gin.Engine, pc *controllers.PlayerController) {
	incomingRoutes.GET("/players", pc.GetPlayers)
	incomingRoutes.GET("/players/:id", pc.GetPlayerByID)
	incomingRoutes.POST("/players", pc.CreatePlayer)
	incomingRoutes.PUT("/players/:id", pc.UpdatePlayer)
	incomingRoutes.DELETE("/players/:id", pc.DeletePlayer)
}

func PartyRoutes(incomingRoutes *gin.Engine, pc *controllers.PartyController) {
	incomingRoutes.GET("/party", pc.GetParties)
	incomingRoutes.GET("/party/:id", pc.GetPartyByID)
	incomingRoutes.POST("/party", pc.CreateParty)
	incomingRoutes.PUT("/party/:id", pc.UpdateParty)
	incomingRoutes.DELETE("/party/:id", pc.DeleteParty)
}

func HistoryRoutes(incomingRoutes *gin.Engine, hc *controllers.HistoryController) {
	incomingRoutes.GET("/history", hc.GetHistories)
	incomingRoutes.GET("/history/:id", hc.GetHistoryByID)
}

func CardRoutes(incomingRoutes *gin.Engine, cc *controllers.CardController) {
	incomingRoutes.GET("/card", cc.GetCards)
	incomingRoutes.GET("/card/:id", cc.GetCardByID)
}
