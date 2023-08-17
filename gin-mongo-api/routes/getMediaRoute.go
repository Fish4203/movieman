package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func GetMediaRoute(router *gin.Engine)  {
    //detailes
    router.GET("/media/details/movie/:movieId", controllers.MovieDetails())
    router.GET("/media/details/show/:showId/:seasonId/:epesodeId", controllers.ShowEpisodeDetails())
    router.GET("/media/details/show/:showId/:seasonId", controllers.ShowSeasonDetails())
    router.GET("/media/details/show/:showId", controllers.ShowDetails())
    router.GET("/media/details/person/:personId", controllers.PersonDetails())
    // serch
    router.GET("/media/search", controllers.GetMediaSearch())
}
