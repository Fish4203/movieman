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
    "strings"
    // "compress/gzip"
    // "time"
    
    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    // "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
    "github.com/mitchellh/mapstructure"
)

var configIndex = &mapstructure.DecoderConfig{
    TagName: "json",
}


func decoderIndex(json map[string]interface{}, obj interface{}) error {
    configIndex.Result = &obj
    // configIndex.DecodeHook = mapstructure.StringToSliceHookFunc(",")
    decoder, _ := mapstructure.NewDecoder(configIndex)
    err := decoder.Decode(json)

    return err
}

func IndexerTest() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("Indexer controler")

        json, err := middleware.JsonRequest(os.Getenv("INDEXERIP") + "/api/v1/indexer", os.Getenv("INDEXERKEY"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, json)
    }
}



func IndexerSearch() gin.HandlerFunc {
    return func(c *gin.Context) {
        var torrents []responses.IndexerResponse

        var err error
        query := strings.Replace(c.Query("q"), " ", "%20", -1)
        json, err := middleware.JsonRequest(os.Getenv("INDEXERIP") + "/api/v1/search?q=" + query, os.Getenv("INDEXERKEY"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        
        results := json["results"].([]interface{})
        var result map[string]interface{}


        for i := 0; i < len(results); i++ {
            result = results[i].(map[string]interface{})

            var res responses.IndexerResponse
            err := decoderIndex(result, &res)
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }


            torrents = append(torrents, res)
        }
        
        c.JSON(http.StatusOK, torrents)
    }
}
