package routes

import (
    "backend-mediaman/controllers"
    "backend-mediaman/middleware"
    "github.com/gin-gonic/gin"
)

func IndexerRoute(router *gin.Engine)  {
    router.GET("/indexer", 		    controllers.IndexerTest())
    router.GET("/indexer/search",   middleware.AuthMiddleware(),controllers.IndexerSearch())
    // router.GET("/indexrs", 		controllers.IndexList())
}
