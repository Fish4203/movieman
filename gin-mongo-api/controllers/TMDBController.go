package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    // "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    // "gin-mongo-api/middleware"
    // "fmt"
	"net/http"
	"io"
    "encoding/json"
    "os"
    // "time"

    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    // "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
)

func getTMDB(params string) (map[string]interface{}, error) {
    var jsonResponse map[string]interface{}

    req, err := http.NewRequest(http.MethodGet, "https://api.themoviedb.org/3/" + params, nil)
    if err != nil {
        return jsonResponse, err
    }

    req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + os.Getenv("TMDB"))


    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return jsonResponse, err
    }

    defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    if err != nil {
        return jsonResponse, err
    }

    err = json.Unmarshal(body, &jsonResponse)
    if err != nil {
        return jsonResponse, err
    }

    return jsonResponse, nil
}

func TMDBTest() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        // defer cancel()

        json, err := getTMDB("authentication")
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }


        c.JSON(http.StatusOK, json)
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
