package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func GameRoute(router *gin.Engine)  {
  router.   GET("/game",                                     controllers.GetGames())
  // router.   PUT("/game/merge",   middleware.AuthMiddleware(),controllers.MergeGames())
  router.DELETE("/game",         middleware.AuthMiddleware(),controllers.DeleteGame())
  router.  POST("/game",         middleware.AuthMiddleware(),controllers.CreateGame())

  router.  POST("/game/review",  middleware.AuthMiddleware(),controllers.CreateGameReview())
  router.   GET("/game/review/:userId",                      controllers.GetGameReview())
  router.DELETE("/game/review",  middleware.AuthMiddleware(),controllers.DeleteGameReview())
}
