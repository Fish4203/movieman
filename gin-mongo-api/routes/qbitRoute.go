package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func QbtRoute(router *gin.Engine)  {
    router.GET("/qbt", 			controllers.QbtGetAll())
    router.POST("/qbt", 		controllers.QbtAdd())
    router.GET("/qbt/:name",	controllers.QbtGet())
    router.PUT("/qbt/:name",	controllers.QbtEdit())
    router.DELETE("/qbt/:name",	controllers.QbtDelete())
}
