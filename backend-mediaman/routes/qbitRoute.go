package routes

import (
    "backend-mediaman/controllers"
    "backend-mediaman/middleware"
    "github.com/gin-gonic/gin"
)

func QbtRoute(router *gin.Engine)  {
    router.   GET("/qbt", 			     controllers.QbtGetAll())
    router.  POST("/qbt", 		         middleware.AuthMiddleware(), controllers.QbtAdd())
    router.   GET("/qbt/:name",	         controllers.QbtGet())
    router.   PUT("/qbt/:name/:category",middleware.AuthMiddleware(), controllers.QbtEdit())
    router.DELETE("/qbt/:name",	         middleware.AuthMiddleware(), controllers.QbtDelete())
}
