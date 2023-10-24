package routes

import (
	"gin-mongo-api/controllers"
	"gin-mongo-api/middleware"

	"github.com/gin-gonic/gin"
)

func SearchProviderRoute(router *gin.Engine) {
	router.GET("/prov", controllers.GetProviders())
	router.POST("/prov", middleware.AuthMiddleware(), controllers.CreateProvider())
	router.DELETE("/prov/:Id", middleware.AuthMiddleware(), controllers.DeleteProvider())
}
