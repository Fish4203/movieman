package routes

import (
	"TMDB-api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/TMDB", controllers.TMDBTest())
	//detailes
	router.GET("/TMDB/movie/:movieId", controllers.TMDBMovieDetails())
	router.GET("/TMDB/show/:showId", controllers.TMDBShowDetails())
	router.GET("/TMDB/person/:personId", controllers.TMDBPersonDetails())
	// router.GET("/TMDB/company/:Id", controllers.TMDBCompanyDetails())
	router.GET("/TMDB/group/:Id", controllers.TMDBCollectionDetails())
	// search
	router.GET("/TMDB/search", controllers.TMDBSearch())
	router.GET("/TMDB/popluar", controllers.TMDBPopular())
}
