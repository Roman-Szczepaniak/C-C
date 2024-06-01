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

func CardRoutes(incomingRoutes *gin.Engine, cc *controllers.CardController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/cards", cc.GetCards)
	incomingRoutes.GET("/cards/:id", cc.GetCardByID)
	incomingRoutes.POST("/cards", cc.CreateCard)
	incomingRoutes.PUT("/cards/:id", cc.UpdateCard)
	incomingRoutes.DELETE("/cards/:id", cc.DeleteCard)
}

func ActionRoutes(incomingRoutes *gin.Engine, ac *controllers.ActionController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/actions", ac.GetActions)
	incomingRoutes.GET("/actions/:id", ac.GetActionByID)
	incomingRoutes.POST("/actions", ac.CreateAction)
	incomingRoutes.PUT("/actions/:id", ac.UpdateAction)
	incomingRoutes.DELETE("/actions/:id", ac.DeleteAction)
}

func MonsterRoutes(incomingRoutes *gin.Engine, mc *controllers.MonsterController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/monsters", mc.GetMonsters)
	incomingRoutes.GET("/monsters/:id", mc.GetMonsterByID)
	incomingRoutes.POST("/monsters", mc.CreateMonster)
	incomingRoutes.PUT("/monsters/:id", mc.UpdateMonster)
	incomingRoutes.DELETE("/monsters/:id", mc.DeleteMonster)
}

func PlayerRoutes(incomingRoutes *gin.Engine, plc *controllers.PlayerController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/players", plc.GetPlayers)
	incomingRoutes.GET("/players/:id", plc.GetPlayerByID)
	incomingRoutes.POST("/players", plc.CreatePlayer)
	incomingRoutes.PUT("/players/:id", plc.UpdatePlayer)
	incomingRoutes.DELETE("/players/:id", plc.DeletePlayer)
}

func PartyRoutes(incomingRoutes *gin.Engine, ptc *controllers.PartyController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/parties", ptc.GetParties)
	incomingRoutes.GET("/parties/:id", ptc.GetPartyByID)
	incomingRoutes.POST("/parties", ptc.CreateParty)
	incomingRoutes.PUT("/parties/:id", ptc.UpdateParty)
	incomingRoutes.DELETE("/parties/:id", ptc.DeleteParty)
}

func EncounterRoutes(incomingRoutes *gin.Engine, ec *controllers.EncounterController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/encounters", ec.GetEncounters)
	incomingRoutes.GET("/encounters/:id", ec.GetEncounterByID)
	incomingRoutes.POST("/encounters", ec.CreateEncounter)
	incomingRoutes.PUT("/encounters/:id", ec.UpdateEncounter)
	incomingRoutes.DELETE("/encounters/:id", ec.DeleteEncounter)
}

func HistoryRoutes(incomingRoutes *gin.Engine, hc *controllers.HistoryController) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/histories", hc.GetHistories)
	incomingRoutes.GET("/histories/:id", hc.GetHistoryByID)
}
