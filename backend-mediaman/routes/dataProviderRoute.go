package routes

import (
	"backend-mediaman/controllers"
	"backend-mediaman/middleware"

	"github.com/gin-gonic/gin"
)

func DataProviderRoute(router *gin.Engine)  {
  router.   GET("/provider/:providerID",                                controllers.GetProvidor())
  router.   GET("/provider",                                            controllers.GetProvidors())
  router.DELETE("/provider/:providerID",          middleware.UserAuth(),controllers.DeleteProvider())
  router.  POST("/provider/:providerID",          middleware.UserAuth(),controllers.SaveProvider())
}
