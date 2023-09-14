package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func WatchRoute(router *gin.Engine)  {
    router.GET("/watch/movie/:query",   controllers.WatchGetMovie())
    router.GET("/watch/show/:query",    controllers.WatchGetShow())
    router.GET("/watch/game/:query",    controllers.WatchGetGame())
    router.GET("/watch/book/:query",    controllers.WatchGetBook())
    router.GET("/watch/:query", 		controllers.WatchGetSearch())
    router.GET("/watch", 				controllers.WatchGetAll())
    router.POST("/watch", 				controllers.WatchPost())
    // router.PUT("/watch",  				controllers.WatchEdit())
    router.DELETE("/watch",  			controllers.WatchDelete())
}
