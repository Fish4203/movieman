package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    router.  POST("/user",                                controllers.CreateUser())
    router.   GET("/user",          middleware.UserAuth(),controllers.GetUser())
    router.   GET("/user/:userID",                        controllers.GetAUser())
    router.   PUT("/user",          middleware.UserAuth(),controllers.EditAUser())
    router.DELETE("/user",          middleware.UserAuth(),controllers.DeleteAUser())
    router.   GET("/users",                               controllers.GetAllUsers())
    router.  POST("/login",                               controllers.LoginUser())
}
