package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func TMDBRoute(router *gin.Engine)  {
    router.GET("/TMDB", controllers.TMDBTest())
    //detailes
    router.GET("/TMDB/details/movie/:movieId", controllers.TMDBMovieDetails())
    router.GET("/TMDB/details/show/:showId", controllers.TMDBShowDetails())
    router.GET("/TMDB/details/person/:personId", controllers.TMDBPersonDetails())
    // serch
    router.GET("/TMDB/search", controllers.TMDBSearch())
}
