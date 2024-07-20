package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func ShowSeasonRoute(router *gin.Engine)  {
  router.   GET("/season",                                     controllers.GetShowSeasons())
  // router.   PUT("/season/merge",   middleware.AuthMiddleware(),controllers.MergeShowSeasons())
  router.DELETE("/season",         middleware.AuthMiddleware(),controllers.DeleteShowSeason())
  router.  POST("/season",         middleware.AuthMiddleware(),controllers.CreateShowSeason())

  router.  POST("/season/review",  middleware.AuthMiddleware(),controllers.CreateShowSeasonReview())
  router.   GET("/season/review/:userId",                      controllers.GetShowSeasonReview())
  router.DELETE("/season/review",  middleware.AuthMiddleware(),controllers.DeleteShowSeasonReview())
}
