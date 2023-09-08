package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func IndexerRoute(router *gin.Engine)  {
    router.GET("/indexer", 		controllers.IndexerTest())
    router.GET("/indexer/search", controllers.IndexerSearch())
    // router.GET("/indexrs", 		controllers.IndexList())
}
