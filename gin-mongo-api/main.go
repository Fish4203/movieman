package main

import (
    "gin-mongo-api/configs"
    "gin-mongo-api/middleware"
    "gin-mongo-api/routes"
    "github.com/gin-gonic/gin"
)

func main() {
        router := gin.Default()

        configs.ConnectDB()

        router.Use(middleware.AuthMiddleware())

        routes.UserRoute(router)

        router.Run("localhost:3000")
}
