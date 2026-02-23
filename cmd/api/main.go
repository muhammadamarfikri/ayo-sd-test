package main

import (
	"ayo/config"
	"ayo/delivery/handler"
	"ayo/impl/postgres"
	mm "ayo/match_management"
	mrm "ayo/match_result_management"
	pm "ayo/player_management"
	tm "ayo/team_management"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := config.AppServer()
	db := config.NewPostgres()

	repo := postgres.NewPostgres(db)

	teamManagement := tm.NewTeamManagement(repo)

	playerManagement := pm.NewPlayerManagement(repo)

	matchManagement := mm.NewMatchManagement(repo)

	matchResultManagement := mrm.NewMatchResultManagement(repo)

	h := handler.NewHandler(teamManagement, playerManagement, matchManagement, matchResultManagement)

	server := gin.New()

	api := server.Group("/api")
	{
		api.POST("/team-management", h.CreateTeam)
		api.GET("/:id/team-management", h.GetTeam)
		api.PATCH("/:id/team-management", h.UpdateTeam)
		api.DELETE("/:id/team-management", h.DeleteTeam)

		api.POST("/player-management", h.CreatePlayer)
		api.GET("/:id/player-management", h.GetPlayer)
		api.PATCH("/:id/player-management", h.UpdatePlayer)
		api.DELETE("/:id/player-management", h.DeletePlayer)

		api.POST("/match-management", h.CreateMatch)
		api.GET("/:id/match-management", h.GetMatch)
		api.PATCH("/:id/match-management", h.UpdateMatch)
		api.DELETE("/:id/match-management", h.DeleteMatch)

		api.POST("/match-result-management", h.CreateMatchResult)
		api.GET("/:id/match-result-management", h.GetMatchResult)
		api.PATCH("/:id/match-result-management", h.UpdateMatchResult)
		api.DELETE("/:id/match-result-management", h.DeleteMatch)
	}

	err := server.Run(":" + appConfig.Port)
	if err != nil {
		fmt.Println("error while starting server: ", err)
	}

}
