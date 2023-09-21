package routes

import (
    "gin-mongo-api/controllers"
    "gin-mongo-api/middleware"
    "github.com/gin-gonic/gin"
)

func WatchRoute(router *gin.Engine)  {
    router.GET("/watch/movie/:query",   middleware.AuthMiddleware(), controllers.WatchGetMovie())
    router.GET("/watch/show/:query",    middleware.AuthMiddleware(), controllers.WatchGetShow())
    router.GET("/watch/game/:query",    middleware.AuthMiddleware(), controllers.WatchGetGame())
    router.GET("/watch/book/:query",    middleware.AuthMiddleware(), controllers.WatchGetBook())
    router.GET("/watch/:query", 		middleware.AuthMiddleware(), controllers.WatchGetSearch())
    router.GET("/watch", 				middleware.AuthMiddleware(), controllers.WatchGetAll())
    router.POST("/watch", 				middleware.AuthMiddleware(), controllers.WatchPost())
    // router.PUT("/watch",  				controllers.WatchEdit())
    router.DELETE("/watch",  			middleware.AuthMiddleware(), controllers.WatchDelete())
}
