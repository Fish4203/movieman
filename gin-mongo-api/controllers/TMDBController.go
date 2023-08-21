package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    // "gin-mongo-api/middleware"
    "fmt"
    "strconv"
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
    "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
)

func genreHelper(genres []interface{}) []string {
    var out []string

    for _, v := range genres {
        re := v.(map[string]interface{})
        out = append(out, re["name"].(string))
    }

    return out
}

func movieDecoderTMDB(json map[string]interface{}) models.Movie {
    movie := models.Movie{
        Title:          json["title"].(string),
        Description:    json["overview"].(string),
        Date:           json["release_date"].(string),
        Popularity:     json["popularity"].(float64),
        VoteCount:      int(json["vote_count"].(float64)),
        VoteRating:     json["vote_average"].(float64),
        TMDB:           int(json["id"].(float64)),
    }

    if json["genres"] != nil {
        movie.Genre = genreHelper(json["genres"].([]interface{}))
    }
    if json["homepage"] != nil {
        movie.Info = json["homepage"].(string)
    }
    if json["imdb_id"] != nil {
        movie.IMDB = json["imdb_id"].(string)
    }
    if json["runtime"] != nil {
        movie.Length = int(json["runtime"].(float64))
    }
    if json["budget"] != nil {
        movie.Budget = int(json["budget"].(float64))
    }
    if json["poster_path"] != nil {
        movie.Image = json["poster_path"].(string)
    }

    return movie
}

func personDecoderTMDB(json map[string]interface{}) models.Person {
    person := models.Person{
        Name: json["name"].(string),
        Role: json["known_for_department"].(string),
        Description: json["biography"].(string),
        Popularity: json["popularity"].(float64),
        TMDB: int(json["id"].(float64)),
        IMDB: json["imdb_id"].(string),
    }

    if json["profile_path"] != nil {
        person.Image = json["profile_path"].(string)
    }
    if json["birthday"] != nil {
        person.Date = json["birthday"].(string)
    }

    return person
}

func showDecoderTMDB(json map[string]interface{}) models.Show {
    show := models.Show{
        Title:          json["name"].(string),
        Description:    json["overview"].(string),
        Date:           json["first_air_date"].(string),
        Popularity:     json["popularity"].(float64),
        VoteCount:      int(json["vote_count"].(float64)),
        VoteRating:     json["vote_average"].(float64),
        TMDB:           int(json["id"].(float64)),
    }

    if json["genres"] != nil {
        show.Genre = genreHelper(json["genres"].([]interface{}))
    }
    if json["number_of_seasons"] != nil {
        show.Seasons = int(json["number_of_seasons"].(float64))
    }
    if json["homepage"] != nil {
        show.Info = json["homepage"].(string)
    }
    if json["poster_path"] != nil {
        show.Image = json["poster_path"].(string)
    }

    return show
}

func showSeasonDecoderTMDB(json map[string]interface{}) models.ShowSeason {
    show := models.ShowSeason{
        SeasonID:       int(json["season_number"].(float64)),
        Episodes:       len(json["episodes"].([]interface{})),
        Description:    json["overview"].(string),
        Date:           json["air_date"].(string),
    }

    if json["poster_path"] != nil {
        show.Image = json["poster_path"].(string)
    }

    return show
}

func showEpisodeDecoderTMDB(json map[string]interface{}) models.ShowEpisode {
    show := models.ShowEpisode{
        EpisodeID:      int(json["episode_number"].(float64)),
        Title:          json["name"].(string),
        Description:    json["overview"].(string),
        Date:           json["air_date"].(string),
        VoteCount:      int(json["vote_count"].(float64)),
        VoteRating:     json["vote_average"].(float64),
    }

    if json["poster_path"] != nil {
        show.Image = json["poster_path"].(string)
    }

    return show
}


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

    err = json.Unmarshal(body, &jsonResponse)
    if err != nil {
        return jsonResponse, err
    }

    return jsonResponse, nil
}

func TMDBTest() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("TMDB controler")

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
        var err error
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

            if result["media_type"] == "tv" {
                show := showDecoderTMDB(result)
                err = show.Save()
            } else if result["media_type"] == "movie" {
                movie := movieDecoderTMDB(result)
                err = movie.Save()
            } else {
                var jsonPerson map[string]interface{}
                jsonPerson, err = getTMDB("person/" + strconv.Itoa(int(result["id"].(float64))) + "?language=en-US")
                person := personDecoderTMDB(jsonPerson)
                err = person.Save()
            }
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }
        }

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBPopular() gin.HandlerFunc {
    return func(c *gin.Context) {
        var err error
        json, err := getTMDB("trending/all/week?language=en-US")
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

            if result["media_type"] == "tv" {
                show := showDecoderTMDB(result)
                err = show.Save()
            } else if result["media_type"] == "movie" {
                movie := movieDecoderTMDB(result)
                err = movie.Save()
            } else {
                var jsonPerson map[string]interface{}
                jsonPerson, err = getTMDB("person/" + strconv.Itoa(int(result["id"].(float64))) + "?language=en-US")
                person := personDecoderTMDB(jsonPerson)
                err = person.Save()
            }
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }
        }

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBMovieDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := strings.Replace(c.Param("movieId"), " ", "%20", -1)
        result, err := getTMDB("movie/" + query + "?append_to_response=recommendations&language=en-US")

        movie := movieDecoderTMDB(result)


        recommendations := result["recommendations"].(map[string]interface{})["results"].([]interface{})
        var movieRecs []primitive.ObjectID
        var showRecs  []primitive.ObjectID

        for i := 1; i < len(recommendations); i++ {
            recommendation := recommendations[i].(map[string]interface{})
            // fmt.Println(recommendation)
            if recommendation["media_type"] == "tv" {
                showRec := showDecoderTMDB(recommendation)
                errRec := showRec.Save()

                if errRec != nil {
                    c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
                } else {
                    showRecs = append(showRecs, showRec.Id)
                }
            } else {
                movieRec := movieDecoderTMDB(recommendation)
                errRec := movieRec.Save()

                if errRec != nil {
                    c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
                } else {
                    movieRecs = append(movieRecs, movieRec.Id)
                }
            }
        }

        movie.AdjMovies = movieRecs
        movie.AdjShows  = showRecs

        err = movie.Save()
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        } else {
            c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
        }
    }
}

func TMDBPersonDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := strings.Replace(c.Param("personId"), " ", "%20", -1)
        json, err := getTMDB("person/" + query + "?language=en-US")

        person := personDecoderTMDB(json)

        err = person.Save()
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBShowDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := strings.Replace(c.Param("showId"), " ", "%20", -1)
        result, err := getTMDB("tv/" + query + "?append_to_response=recommendations&language=en-US")

        show := showDecoderTMDB(result)

        recommendations := result["recommendations"].(map[string]interface{})["results"].([]interface{})
        var movieRecs []primitive.ObjectID
        var showRecs  []primitive.ObjectID

        for i := 1; i < len(recommendations); i++ {
            recommendation := recommendations[i].(map[string]interface{})
            // fmt.Println(recommendation)
            if recommendation["media_type"] == "tv" {
                showRec := showDecoderTMDB(recommendation)
                errRec := showRec.Save()

                if errRec != nil {
                    c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
                } else {
                    showRecs = append(showRecs, showRec.Id)
                }
            } else {
                movieRec := movieDecoderTMDB(recommendation)
                errRec := movieRec.Save()

                if errRec != nil {
                    c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
                } else {
                    movieRecs = append(movieRecs, movieRec.Id)
                }
            }
        }

        show.AdjMovies = movieRecs
        show.AdjShows  = showRecs


        err = show.Save()
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        for i := 1; i <= show.Seasons; i++ {
            seasonResult, err := getTMDB("tv/" + query + "/season/" + strconv.Itoa(i) + "?language=en-US")

            season := showSeasonDecoderTMDB(seasonResult)
            season.ShowId = show.Id

            err = season.Save()
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }

            episodes := seasonResult["episodes"].([]interface{})

            for j := 0; j < len(episodes); j++ {
                episode := showEpisodeDecoderTMDB(episodes[j].(map[string]interface{}))
                episode.ShowId = show.Id
                episode.SeasonID = i

                err = episode.Save()
                if err != nil {
                    c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
                }
            }
        }

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}
