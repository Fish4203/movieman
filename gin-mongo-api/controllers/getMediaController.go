package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    // "gin-mongo-api/middleware"
    // "fmt"
	"net/http"
	// "io"
    // "encoding/json"
    // "os"
    // "strings"
    // "compress/gzip"
    // "time"
    "strconv"

    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
)


func GetMediaSearch() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := c.Query("q")

        movies, err := models.FindMovie(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }
        persons, err := models.FindPerson(bson.D{{"name", bson.D{{"$regex", query}, {"$options", "i"}}}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }
        shows, err := models.FindShow(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }
        episodes, err := models.FindShowEpisode(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"movies": movies, "shows": shows, "episodes": episodes, "people": persons})
    }
}

func MovieDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := c.Param("movieId")
        objId, _ := primitive.ObjectIDFromHex(query)

        movies, err := models.FindMovie(bson.D{{"_id", objId}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"movies": movies})
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
        query := c.Param("showId")
        objId, _ := primitive.ObjectIDFromHex(query)

        results, err := models.FindShow(bson.D{{"_id", objId}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"shows": results})
    }
}

func ShowSeasonDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := c.Param("showId")
        season, err := strconv.Atoi(c.Param("seasonId"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        objId, _ := primitive.ObjectIDFromHex(query)

        results, err := models.FindShowSeason(bson.D{{"showId", objId}, {"seasonId", season}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"season": results})
    }
}


func ShowEpisodeDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := c.Param("showId")
        season, err := strconv.Atoi(c.Param("seasonId"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }
        episode, err := strconv.Atoi(c.Param("episodeId"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }


        objId, _ := primitive.ObjectIDFromHex(query)

        results, err := models.FindShowEpisode(bson.D{{"showId", objId}, {"seasonId", season}, {"episodeId", episode}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"episode": results})
    }
}
