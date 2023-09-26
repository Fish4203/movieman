package main

import (
    "gin-mongo-api/configs"
    // "gin-mongo-api/middleware"
    "gin-mongo-api/routes"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
        router := gin.Default()

        configs.InitClient()

        corsConfig := cors.DefaultConfig()
        corsConfig.AllowOrigins = []string{"*"}
        corsConfig.AllowCredentials = true
        corsConfig.AllowHeaders = []string{"Authorization"}
        // corsConfig.Debug = true

        // router.Use(middleware.AuthMiddleware())
        router.Use(cors.New(corsConfig))

        routes.UserRoute(router)
        routes.TMDBRoute(router)
        routes.GetMediaRoute(router)
        routes.IndexerRoute(router)
        routes.QbtRoute(router)
        routes.WatchRoute(router)

        router.Run("localhost:4000")
}
