package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    // "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "gin-mongo-api/middleware"
    "fmt"
    // "strconv"
	"net/http"
	// "io"
    // "encoding/json"
    "os"
    // "strings"
    // "regexp"
    // "compress/gzip"
    // "time"
    
    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    // "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
    // "github.com/mitchellh/mapstructure"
)



func QbtGetAll() gin.HandlerFunc {
    return func(c *gin.Context) {
        var torrents []responses.QbtResponse

        var err error
        json, err := middleware.JsonRequest(os.Getenv("QBTURL") + "/api/v2/torrents/info", "")
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        
        results := json["results"].([]interface{})
        var result map[string]interface{}


        for i := 0; i < len(results); i++ {
            result = results[i].(map[string]interface{})

            var res responses.QbtResponse
            err := decoderJson(result, &res)
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }

            torrents = append(torrents, res)
        }
        
        c.JSON(http.StatusOK, torrents)
    }
}

func QbtAdd() gin.HandlerFunc {
    return func(c *gin.Context) {
        json := "tree"

        c.JSON(http.StatusOK, json)
    }
}


func QbtGet() gin.HandlerFunc {
    return func(c *gin.Context) {
        var err error
        query := c.Param("name")

        json, err := middleware.JsonRequest(os.Getenv("QBTURL") + "/api/v2/torrents/properties?hash=" + query, "")
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }
        fmt.Println(json)

        var res responses.QbtResponse
        err = decoderJson(json, &res)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }
        

        c.JSON(http.StatusOK, res)
    }
}

func QbtEdit() gin.HandlerFunc {
    return func(c *gin.Context) {
        json := "tree"

        c.JSON(http.StatusOK, json)
    }
}

func QbtDelete() gin.HandlerFunc {
    return func(c *gin.Context) {
        json := "tree"

        c.JSON(http.StatusOK, json)
    }
}
