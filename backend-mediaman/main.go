package main

import (
	"backend-mediaman/configs"
	"backend-mediaman/models"
	"backend-mediaman/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    configs.GetDB()

    configs.DB.AutoMigrate(
      &models.User{}, 
      &models.Movie{}, 
      &models.MovieExternal{}, 
      &models.MovieReview{},
      &models.Game{}, 
      &models.GameExternal{}, 
      &models.GameReview{},
      &models.Book{}, 
      &models.BookExternal{}, 
      &models.BookReview{},
      &models.Show{}, 
      &models.ShowExternal{}, 
      &models.ShowReview{},
      &models.ShowSeason{}, 
      &models.ShowSeasonExternal{}, 
      &models.ShowSeasonReview{},
      &models.ShowEpisode{}, 
      &models.ShowEpisodeExternal{}, 
      &models.ShowEpisodeReview{},
    )
    
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"*"}
    corsConfig.AllowCredentials = true
    corsConfig.AllowHeaders = []string{"Authorization"}
    // corsConfig.Debug = true

    // router.Use(middleware.AuthMiddleware())
    router.Use(cors.New(corsConfig))

    routes.UserRoute(router)
    routes.MovieRoute(router)
    // routes.GetMediaRoute(router)
    // routes.PostMediaRoute(router)
    // routes.IndexerRoute(router)
    // routes.QbtRoute(router)
    // routes.WatchRoute(router)
    // routes.SearchProviderRoute(router)

    router.Run("localhost:4000")
}
