package routes

import (
    "gin-mongo-api/controllers"
    "gin-mongo-api/middleware"
    "github.com/gin-gonic/gin"
)

func IndexerRoute(router *gin.Engine)  {
    router.GET("/indexer", 		    controllers.IndexerTest())
    router.GET("/indexer/search",   middleware.AuthMiddleware(),controllers.IndexerSearch())
    // router.GET("/indexrs", 		controllers.IndexList())
}
