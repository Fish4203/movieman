package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    // "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "gin-mongo-api/middleware"
    // "fmt"
    // "strconv"
	"net/http"
    "net/url"
	"io"
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
        var torrent responses.IndexerResponse

        if err := c.BindJSON(&torrent); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        var downloadUrl string
        if torrent.Magnet != "" {
            downloadUrl = torrent.Magnet 
        } else {
            downloadUrl = torrent.Download
        }

        data := url.Values{
            "urls": {downloadUrl},
            "savepath": {torrent.Catagory +"s"},
            "category": {torrent.Catagory+"s"},
        }

        res, err := http.PostForm(os.Getenv("QBTURL") + "/api/v2/torrents/add", data)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        defer res.Body.Close()

        if res.StatusCode == 200 {
            c.JSON(http.StatusOK, "success")
        } else {
            body, _ := io.ReadAll(res.Body)
            c.JSON(http.StatusBadRequest, string(body))

        }
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

        var res responses.QbtResponse
        err = decoderJson(json, &res)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }
        

        c.JSON(http.StatusOK, res)
    }
}

func QbtEdit() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := c.Param("name")
        category := c.Param("category")

        data := url.Values{
            "hashes": {query},
            "location": {"/data/" + category +"s"},
            "category": {category+"s"},
        }

        res1, err := http.PostForm(os.Getenv("QBTURL") + "/api/v2/torrents/setCategory", data)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }
        defer res1.Body.Close()

        if res1.StatusCode != 200 {
            body, _ := io.ReadAll(res1.Body)
            c.JSON(http.StatusBadRequest, string(body))
            return
        }

        res2, err := http.PostForm(os.Getenv("QBTURL") + "/api/v2/torrents/setLocation", data)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        defer res2.Body.Close()
        if res2.StatusCode != 200 {
            body, _ := io.ReadAll(res2.Body)
            c.JSON(http.StatusBadRequest, string(body))
            return
        }

        c.JSON(http.StatusOK, "success")
    }
}

func QbtDelete() gin.HandlerFunc {
    return func(c *gin.Context) {
        var err error
        query := c.Param("name")

        req, err := http.NewRequest(http.MethodDelete, os.Getenv("QBTURL") + "/api/v2/torrents/delete?hashes=" + query + "&deleteFiles=true", nil)
        client := &http.Client{}
        res, err := client.Do(req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        defer res.Body.Close()

        if res.StatusCode == 200 {
            c.JSON(http.StatusOK, "success")
        } else {
            c.JSON(res.StatusCode, "badrequest")
        }
    }
}
