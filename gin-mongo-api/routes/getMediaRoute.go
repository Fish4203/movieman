package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func GetMediaRoute(router *gin.Engine)  {
    //detailes
    router.GET("/media/movie/:movieId", controllers.MovieDetails())
    router.GET("/media/show/:showId/:seasonId/:episodeId", controllers.ShowEpisodeDetails())
    router.GET("/media/show/:showId/:seasonId", controllers.ShowSeasonDetails())
    router.GET("/media/show/:showId", controllers.ShowDetails())
    router.GET("/media/person/:personId", controllers.PersonDetails())
    // serch
    router.GET("/media/search", controllers.GetMediaSearch())
}
