package main

import (
    "gin-mongo-api/configs"
    "gin-mongo-api/middleware"
    "gin-mongo-api/routes"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
        router := gin.Default()

        configs.InitClient()

        router.Use(middleware.AuthMiddleware())
        router.Use(cors.Default())

        routes.UserRoute(router)
        routes.TMDBRoute(router)
        routes.GetMediaRoute(router)
        routes.IndexerRoute(router)

        router.Run("localhost:4000")
}
