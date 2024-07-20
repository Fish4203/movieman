package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func ShowEpisodeRoute(router *gin.Engine)  {
  router.   GET("/episode",                                     controllers.GetShowEpisodes())
  // router.   PUT("/episode/merge",   middleware.AuthMiddleware(),controllers.MergeShowEpisodes())
  router.DELETE("/episode",         middleware.AuthMiddleware(),controllers.DeleteShowEpisode())
  router.  POST("/episode",         middleware.AuthMiddleware(),controllers.CreateShowEpisode())

  router.  POST("/episode/review",  middleware.AuthMiddleware(),controllers.CreateShowEpisodeReview())
  router.   GET("/episode/review/:userId",                      controllers.GetShowEpisodeReview())
  router.DELETE("/episode/review",  middleware.AuthMiddleware(),controllers.DeleteShowEpisodeReview())
}
