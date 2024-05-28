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
	incomingRoutes.GET("/user/:id", uc.GetUserByID)
	incomingRoutes.PUT("/user/:id", uc.UpdateUser)
	incomingRoutes.DELETE("/user/:id", uc.DeleteUser)
}

func EncounterRoutes(incomingRoutes *gin.Engine, ec *controllers.EncounterController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/encounters", ec.GetEncounters)
	incomingRoutes.GET("/encounter/:id", ec.GetEncounterByID)
	incomingRoutes.POST("/encounter", ec.CreateEncounter)
	incomingRoutes.PUT("/encounter/:id", ec.UpdateEncounter)
	incomingRoutes.DELETE("/encounter/:id", ec.DeleteEncounter)
}

func MonsterRoutes(incomingRoutes *gin.Engine, mc *controllers.MonsterController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/monsters", mc.GetMonsters)
	incomingRoutes.GET("/monster/:id", mc.GetMonsterByID)
	incomingRoutes.POST("/monster", mc.CreateMonster)
	incomingRoutes.PUT("/monster/:id", mc.UpdateMonster)
	incomingRoutes.DELETE("/monster/:id", mc.DeleteMonster)
}

func PlayerRoutes(incomingRoutes *gin.Engine, pc *controllers.PlayerController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/players", pc.GetPlayers)
	incomingRoutes.GET("/player/:id", pc.GetPlayerByID)
	incomingRoutes.POST("/player", pc.CreatePlayer)
	incomingRoutes.PUT("/player/:id", pc.UpdatePlayer)
	incomingRoutes.DELETE("/player/:id", pc.DeletePlayer)
}

func PartyRoutes(incomingRoutes *gin.Engine, pc *controllers.PartyController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/parties", pc.GetParties)
	incomingRoutes.GET("/party/:id", pc.GetPartyByID)
	incomingRoutes.POST("/party", pc.CreateParty)
	incomingRoutes.PUT("/party/:id", pc.UpdateParty)
	incomingRoutes.DELETE("/party/:id", pc.DeleteParty)
}

func HistoryRoutes(incomingRoutes *gin.Engine, hc *controllers.HistoryController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/histories", hc.GetHistories)
	incomingRoutes.GET("/history/:id", hc.GetHistoryByID)
}

func CardRoutes(incomingRoutes *gin.Engine, cc *controllers.CardController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/cards", cc.GetCards)
	incomingRoutes.GET("/card/:id", cc.GetCardByID)
}
