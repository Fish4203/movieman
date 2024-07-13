package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

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
