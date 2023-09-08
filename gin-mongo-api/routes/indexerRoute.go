package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func IndexRoute(router *gin.Engine)  {
    router.GET("/index", 		controllers.IndexTest())
    router.GET("/index/search", controllers.IndexSearch())
    // router.GET("/indexrs", 		controllers.IndexList())
}
