package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    // "gin-mongo-api/middleware"
    "fmt"
	"net/http"
	"io"
    "encoding/json"
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
)

func getTMDB(params string) (map[string]interface{}, error) {
    var jsonResponse map[string]interface{}

    url := "https://api.themoviedb.org/3/" + params
    // fmt.Println(url
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return jsonResponse, err
    }

    // req.Header.Add("Accept-Encoding", "gzip, deflate, br")
    req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + os.Getenv("TMDB"))


    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return jsonResponse, err
    }


    // reader, err := gzip.NewReader(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer reader.Close()

    defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    if err != nil {
        return jsonResponse, err
    }

    // fmt.Println(string(body))

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

        query := strings.Replace(c.Query("q"), " ", "%20", -1)
        json, err := getTMDB("search/multi?query=" + query + "&include_adult=false&language=en-US&page=1")
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        var result map[string]interface{}

        length := 20
        if int(json["total_results"].(float64)) < 20 {
            length = int(json["total_results"].(float64))
        }

        for i := 0; i < length; i++ {
            result = json["results"].([]interface{})[i].(map[string]interface{})
            // fmt.Println(results[i])

            if result["media_type"] == "tv" {
                fmt.Println("tv")
                show := models.Show{
                    Title: result["name"].(string),
                    Description: result["overview"].(string),
                    Date: result["first_air_date"].(string),
                    Popularity: result["popularity"].(float64),
                    VoteCount: int(result["vote_count"].(float64)),
                    VoteRating: result["vote_average"].(float64),
                    TMDB: int(result["id"].(float64)),
                }
                if result["poster_path"] != nil {
                    show.Image = result["poster_path"].(string)
                }

                err := show.Save()
                if err != nil {
                    c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
                }

            } else if result["media_type"] == "movie" {
                // fmt.Println("movie")

                movie := models.Movie{
                    Title: result["title"].(string),
                    Description: result["overview"].(string),
                    Date: result["release_date"].(string),
                    Popularity: result["popularity"].(float64),
                    VoteCount: int(result["vote_count"].(float64)),
                    VoteRating: result["vote_average"].(float64),
                    TMDB: int(result["id"].(float64)),
                }

                if result["poster_path"] != nil {
                    movie.Image = result["poster_path"].(string)
                }

                err := movie.Save()
                if err != nil {
                    c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
                }

            } else {
                fmt.Println("else")
            }
        }

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
