package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    // "gin-mongo-api/middleware"
    "net/http"
    // "time"
    // "fmt"

    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
)


func WatchGetAll() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIdString := c.MustGet("userId").(string)
        userId, err := primitive.ObjectIDFromHex(userIdString)
        if userIdString == "" || err != nil {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }

        watchs, err := models.FindWatch(bson.D{{"user", userId}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, watchs)
    }
}

func WatchGetSearch() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIdString := c.MustGet("userId").(string)
        userId, err := primitive.ObjectIDFromHex(userIdString)
        if userIdString == "" || err != nil {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }

        query := c.Param("query")

        watchs, err := models.FindWatch(bson.D{{"title", query}, {"user", userId}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, watchs)
    }
}

func WatchPost() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIdString := c.MustGet("userId").(string)
        userId, err := primitive.ObjectIDFromHex(userIdString)
        if userIdString == "" || err != nil {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }
        var watch models.Watch
        watch.User = userId

        if err := c.BindJSON(&watch); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        //use the validator library to validate required fields
        if err := validate.Struct(&watch); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        err = models.WriteWatch([]models.Watch{watch})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, "success")
    }
}


func WatchDelete() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIdString := c.MustGet("userId").(string)
        if userIdString == "" {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }
        query := c.Param("query")
        err := models.DeleteWatch([]string{query})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, "success")
    }
}

func WatchGetMovie() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIdString := c.MustGet("userId").(string)
        userId, err := primitive.ObjectIDFromHex(userIdString)
        if userIdString == "" || err != nil {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }
        query := c.Param("query")

        watchs, err := models.FindWatch(bson.D{{"movie", query}, {"user", userId}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, watchs)
    }
}

func WatchGetShow() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIdString := c.MustGet("userId").(string)
        userId, err := primitive.ObjectIDFromHex(userIdString)
        if userIdString == "" || err != nil {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }
        query := c.Param("query")

        watchs, err := models.FindWatch(bson.D{{"show", query}, {"user", userId}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, watchs)
    }
}

func WatchGetGame() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIdString := c.MustGet("userId").(string)
        userId, err := primitive.ObjectIDFromHex(userIdString)
        if userIdString == "" || err != nil {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }
        query := c.Param("query")

        watchs, err := models.FindWatch(bson.D{{"games", query}, {"user", userId}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, watchs)
    }
}

func WatchGetBook() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIdString := c.MustGet("userId").(string)
        userId, err := primitive.ObjectIDFromHex(userIdString)
        if userIdString == "" || err != nil {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }
        query := c.Param("query")

        watchs, err := models.FindWatch(bson.D{{"books", query}, {"user", userId}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, watchs)
    }
}
