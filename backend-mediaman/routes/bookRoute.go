package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine)  {
  router.   GET("/book",                                     controllers.GetBooks())
  // router.   PUT("/book/merge",   middleware.AuthMiddleware(),controllers.MergeBooks())
  router.DELETE("/book",         middleware.AuthMiddleware(),controllers.DeleteBook())
  router.  POST("/book",         middleware.AuthMiddleware(),controllers.CreateBook())

  router.  POST("/book/review",  middleware.AuthMiddleware(),controllers.CreateBookReview())
  router.   GET("/book/review/:userId",                      controllers.GetBookReview())
  router.DELETE("/book/review",  middleware.AuthMiddleware(),controllers.DeleteBookReview())
}
