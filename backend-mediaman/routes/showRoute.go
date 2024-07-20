package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func ShowRoute(router *gin.Engine)  {
  router.   GET("/show",                                     controllers.GetShows())
  // router.   PUT("/show/merge",   middleware.AuthMiddleware(),controllers.MergeShows())
  router.DELETE("/show",         middleware.AuthMiddleware(),controllers.DeleteShow())
  router.  POST("/show",         middleware.AuthMiddleware(),controllers.CreateShow())

  router.  POST("/show/review",  middleware.AuthMiddleware(),controllers.CreateShowReview())
  router.   GET("/show/review/:userId",                      controllers.GetShowReview())
  router.DELETE("/show/review",  middleware.AuthMiddleware(),controllers.DeleteShowReview())
}
