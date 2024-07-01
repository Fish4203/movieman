package routes

import (
	"gin-mongo-api/controllers"
	"gin-mongo-api/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.Engine)  {
    router.   GET("/movie",                                     controllers.GetMovies())
    router.   GET("/movie/:movieID",                            controllers.GetMovie())
    router.   PUT("/movie/edit",    middleware.AuthMiddleware(),controllers.EditMovie())
    router.   PUT("/movie/merge",   middleware.AuthMiddleware(),controllers.MergeMovies())
    router.DELETE("/movie/:movieID",middleware.AuthMiddleware(),controllers.DeleteMovie())
    router.  POST("/movie",         middleware.AuthMiddleware(),controllers.CreateMovie())
    router.  POST("/movie/bulk",    middleware.AuthMiddleware(),controllers.CreateMovies())
}
