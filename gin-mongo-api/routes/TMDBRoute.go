package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func TMDBRoute(router *gin.Engine)  {
    router.GET("/TMDB", controllers.TMDBTest())
    //detailes
    router.GET("/TMDB/movie/:movieId", controllers.TMDBMovieDetails())
    router.GET("/TMDB/show/:showId", controllers.TMDBShowDetails())
    router.GET("/TMDB/person/:personId", controllers.TMDBPersonDetails())
    // serch
    router.GET("/TMDB/search", controllers.TMDBSearch())
    router.GET("/TMDB/popluar", controllers.TMDBPopular())
}
