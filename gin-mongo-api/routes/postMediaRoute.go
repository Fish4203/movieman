package routes

import (
	"gin-mongo-api/controllers"
	"gin-mongo-api/middleware"

	"github.com/gin-gonic/gin"
)

func PostMediaRoute(router *gin.Engine) {
	//post data
	router.POST("/media/movie", middleware.AuthMiddleware(), controllers.EditMovie())
	router.POST("/media/show", middleware.AuthMiddleware(), controllers.EditShow())
	router.POST("/media/showSeason", middleware.AuthMiddleware(), controllers.EditShowSeason())
	router.POST("/media/showEpisode", middleware.AuthMiddleware(), controllers.EditShowEpisode())
	router.POST("/media/person", middleware.AuthMiddleware(), controllers.EditPerson())
	router.POST("/media/book", middleware.AuthMiddleware(), controllers.EditBook())
	router.POST("/media/game", middleware.AuthMiddleware(), controllers.EditGame())
	router.POST("/media", controllers.BulkAdd())
}
