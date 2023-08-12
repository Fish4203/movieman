package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    // "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    // "gin-mongo-api/middleware"
    "net/http"
    // "time"

    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    // "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
)


func TMDBTest() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}


func TMDBSearch() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBMovieDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBPersonDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBShowDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBShowSeasonDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}
