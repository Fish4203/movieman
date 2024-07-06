package main

import (
    "gin-mongo-api/configs"
    "gin-mongo-api/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    configs.GetDB()

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"*"}
    corsConfig.AllowCredentials = true
    corsConfig.AllowHeaders = []string{"Authorization"}
    // corsConfig.Debug = true

    // router.Use(middleware.AuthMiddleware())
    router.Use(cors.New(corsConfig))

    routes.UserRoute(router)
    // routes.GetMediaRoute(router)
    // routes.PostMediaRoute(router)
    // routes.IndexerRoute(router)
    // routes.QbtRoute(router)
    // routes.WatchRoute(router)
    // routes.SearchProviderRoute(router)

    router.Run("localhost:4000")
}
