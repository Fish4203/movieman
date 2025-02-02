package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.Engine)  {
  router.   GET("/movie/:movieID",                                      controllers.GetMovie())
  router.DELETE("/movie/:movieID",                middleware.UserAuth(),controllers.DeleteMovie())

  router.   GET("/movie/:movieID/review",                               controllers.GetMovieReview())
  router.  POST("/movie/review",                  middleware.UserAuth(),controllers.SaveMovieReview())
  router.DELETE("/movie/:movieID/review/:userID", middleware.UserAuth(),controllers.DeleteMovieReview())
}
