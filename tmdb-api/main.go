package main

import (
	// "TMDB-api/configs"
	// "TMDB-api/middleware"
	"TMDB-api/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Authorization"}
	// corsConfig.Debug = true

	// router.Use(middleware.AuthMiddleware())
	router.Use(cors.New(corsConfig))

	routes.TMDBRoute(router)

	router.Run("localhost:4001")
}
