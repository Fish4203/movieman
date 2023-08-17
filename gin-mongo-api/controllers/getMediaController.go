package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    // "gin-mongo-api/middleware"
    "fmt"
	"net/http"
	// "io"
    // "encoding/json"
    // "os"
    // "strings"
    // "compress/gzip"
    // "time"

    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
)


func GetMediaSearch() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := c.Query("q")
        fmt.Println(query)

        movies, err := models.FindMovie(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"movies": movies})
    }
}

func MovieDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func PersonDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func ShowDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func ShowSeasonDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}


func ShowEpisodeDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}
