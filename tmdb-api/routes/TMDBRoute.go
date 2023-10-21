package routes

import (
	"TMDB-api/controllers"

	"github.com/gin-gonic/gin"
)

func TMDBRoute(router *gin.Engine) {
	router.GET("/TMDB", controllers.TMDBTest())
	//detailes
	router.GET("/TMDB/movie/:movieId", controllers.TMDBMovieDetails())
	router.GET("/TMDB/show/:showId", controllers.TMDBShowDetails())
	router.GET("/TMDB/person/:personId", controllers.TMDBPersonDetails())
	// router.GET("/TMDB/company/:Id", controllers.TMDBCompanyDetails())
	router.GET("/TMDB/collection/:Id", controllers.TMDBCollectionDetails())
	// serch
	router.GET("/TMDB/search/company", controllers.TMDBSearchCompany())
	router.GET("/TMDB/search/collection", controllers.TMDBSearchCollection())
	router.GET("/TMDB/search", controllers.TMDBSearch())
	router.GET("/TMDB/popluar", controllers.TMDBPopular())
}
