package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.Engine)  {
  router.   GET("/movie",                                     controllers.GetMovies())
  // router.   PUT("/movie/merge",   middleware.AuthMiddleware(),controllers.MergeMovies())
  router.DELETE("/movie",         middleware.AuthMiddleware(),controllers.DeleteMovie())
  router.  POST("/movie",         middleware.AuthMiddleware(),controllers.CreateMovie())

  router.  POST("/movie/review",  middleware.AuthMiddleware(),controllers.CreateMovieReview())
  router.   GET("/movie/review/:userId",                      controllers.GetMovieReview())
  router.DELETE("/movie/review",  middleware.AuthMiddleware(),controllers.DeleteMovieReview())
}
