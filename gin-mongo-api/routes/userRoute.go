package routes

import (
	"gin-mongo-api/controllers"
	"gin-mongo-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    router.  POST("/user",                                      controllers.CreateUser())
    router.   GET("/user",          middleware.AuthMiddleware(),controllers.GetUser())
    router.   GET("/user/:userId",                              controllers.GetAUser())
    router.   PUT("/user",          middleware.AuthMiddleware(),controllers.EditAUser())
    router.DELETE("/user",          middleware.AuthMiddleware(),controllers.DeleteAUser())
    router.   GET("/users",                                     controllers.GetAllUsers())
    router.  POST("/login",                                     controllers.Login())
}
