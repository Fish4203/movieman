package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func BulkRoute(router *gin.Engine)  {
  router.POST("/bulk",  middleware.AuthMiddleware(),controllers.BulkAdd())
}
