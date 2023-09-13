package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    // "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "gin-mongo-api/middleware"
    "fmt"
    "strconv"
	"net/http"
	// "io"
    // "encoding/json"
    "os"
    "strings"
    "regexp"
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


func decoderJson(json map[string]interface{}, obj interface{}) error {
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

        yearRegex := regexp.MustCompile(`(19|20)[0-9][0-9]`)
        seasonRegex := regexp.MustCompile(`S[0-9][0-9]`)
        episodeRegex := regexp.MustCompile(`E[0-9][0-9]`)
        encodingRegex := regexp.MustCompile(`((H|h|x)26(4|5)|AV1)`)
        resolutionRegex := regexp.MustCompile(`[0-9]+p`)
        movieRegex := regexp.MustCompile(`20[0-9][0-9]`)
        showRegex := regexp.MustCompile(`50[0-9][0-9]`)
        gameRegex := regexp.MustCompile(`(40[0-9][0-9]|10[0-9][0-9])`)
        bookRegex := regexp.MustCompile(`70[0-9][0-9]`)


        var err error
        query := strings.Replace(c.Query("q"), " ", "%20", -1)
        json, err := middleware.JsonRequest(os.Getenv("INDEXERIP") + "/api/v1/search?query=" + query, os.Getenv("INDEXERKEY"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        
        results := json["results"].([]interface{})
        var result map[string]interface{}


        for i := 0; i < len(results); i++ {
            result = results[i].(map[string]interface{})

            var res responses.IndexerResponse
            err := decoderJson(result, &res)
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }


            res.ReleaseYear, _ = strconv.Atoi(yearRegex.FindString(res.Title))
            res.SeasonNum   = seasonRegex.FindString(res.Title)
            res.EpisodeNum  = episodeRegex.FindString(res.Title)
            res.Encoding    = encodingRegex.FindString(res.Title)
            res.Resolution  = resolutionRegex.FindString(res.Title)

            catagory := strconv.Itoa(int(result["categories"].([]interface{})[0].(map[string]interface{})["id"].(float64))) 

            if movieRegex.MatchString(catagory) {
                res.Catagory = "movie"
            } else if showRegex.MatchString(catagory) {
                res.Catagory = "show"
            } else if gameRegex.MatchString(catagory) {
                res.Catagory = "game"
            } else if bookRegex.MatchString(catagory) {
                res.Catagory = "book"
            } else {
                res.Catagory = "other"
            }

            torrents = append(torrents, res)
        }
        
        c.JSON(http.StatusOK, torrents)
    }
}
