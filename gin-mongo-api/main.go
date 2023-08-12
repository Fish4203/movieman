package main

import (
    "gin-mongo-api/configs"
    "gin-mongo-api/middleware"
    "gin-mongo-api/routes"
    "github.com/gin-gonic/gin"
)

func main() {
        router := gin.Default()

        configs.InitClient()

        router.Use(middleware.AuthMiddleware())

        routes.UserRoute(router)
        routes.TMDBRoute(router)

        router.Run("localhost:3000")
}
