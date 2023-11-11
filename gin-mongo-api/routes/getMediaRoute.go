package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func GetMediaRoute(router *gin.Engine) {
	//detailes
	router.GET("/media/movie/:movieId", controllers.MovieDetails())
	router.GET("/media/show/:showId", controllers.ShowDetails())
	router.GET("/media/episode/:showId/:seasonId/:episodeId", controllers.EpisodeDetails())
	router.GET("/media/book/:Id", controllers.BookDetails())
	router.GET("/media/game/:Id", controllers.GameDetails())
	router.GET("/media/person/:personId", controllers.PersonDetails())
	router.GET("/media/group/:Id", controllers.GroupDetails())
	router.GET("/media/company/:Id", controllers.CompanyDetails())
	// serch
	router.GET("/media/search", controllers.GetMediaSearch())
}
