package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    router.POST("/user", controllers.CreateUser())
    router.GET("/user/:userId", controllers.GetAUser())
    router.PUT("/user", controllers.EditAUser())
    router.DELETE("/user", controllers.DeleteAUser())
    router.GET("/users", controllers.GetAllUsers())
    router.POST("/login", controllers.Login())
}
