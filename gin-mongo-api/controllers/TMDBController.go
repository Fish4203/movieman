package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    "gin-mongo-api/middleware"
    "fmt"
    "strconv"
	"net/http"
	// "io"
    // "encoding/json"
    "os"
    "strings"
    // "compress/gzip"
    // "time"
    
    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    // "golang.org/x/crypto/bcrypt"
    "github.com/mitchellh/mapstructure"
)

var config = &mapstructure.DecoderConfig{
    TagName: "tmdb",
}


func genreHelper(json map[string]interface{}) []string {
    var out []string
    
    if json["genres"] != nil {
        genres := json["genres"].([]interface{})
        for _, v := range genres {
            re := v.(map[string]interface{})
            out = append(out, re["name"].(string))
        }
    } 
    
    return out
}

func extidHelper(json map[string]interface{}) map[string]string  {
    out := make(map[string]string) 
    
    out["tmdb"] = strconv.Itoa(int(json["id"].(float64)))

    if json["imdb_id"] != nil {
        out["imdb"] = json["imdb_id"].(string)
    } 
    
    return out
}

func decoder(json map[string]interface{}, obj interface{}) error {
    config.Result = &obj
    config.DecodeHook = mapstructure.StringToSliceHookFunc(",")
    decoder, _ := mapstructure.NewDecoder(config)
    err := decoder.Decode(json)

    return err
}


func TMDBTest() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("TMDB controler")

        json, err := middleware.JsonRequest("https://api.themoviedb.org/3/authentication", "Bearer " + os.Getenv("TMDB"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, json)
    }
}


func TMDBSearch() gin.HandlerFunc {
    return func(c *gin.Context) {

        var people []models.Person
        var movies []models.Movie
        var shows  []models.Show

        var err error
        query := strings.Replace(c.Query("q"), " ", "%20", -1)
        json, err := middleware.JsonRequest("https://api.themoviedb.org/3/search/multi?query=" + query + "&include_adult=false&language=en-US&page=1", "Bearer " + os.Getenv("TMDB"))
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
                var show models.Show
                decoder(result, &show) 
                show.Genre = genreHelper(result)
                show.ExternalIds = extidHelper(result)
                shows = append(shows, show) 
            } else if result["media_type"] == "movie" {
                var movie models.Movie
                decoder(result, &movie)
                movie.Genre = genreHelper(result)
                movie.ExternalIds = extidHelper(result)
                movies = append(movies, movie) 
            } else {
                var jsonPerson map[string]interface{}
                jsonPerson, _ = middleware.JsonRequest("https://api.themoviedb.org/3/person/" + strconv.Itoa(int(result["id"].(float64))) + "?language=en-US", "Bearer " + os.Getenv("TMDB"))
                var person models.Person
                decoder(jsonPerson, &person)
                person.ExternalIds = extidHelper(jsonPerson)
                people = append(people, person) 
            }
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }
        }

        err = models.WritePerson(people)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }
 
        err = models.WriteMovie(movies)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }
        
        err = models.WriteShow(shows)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }
        
        
        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBPopular() gin.HandlerFunc {
    return func(c *gin.Context) {
        var people []models.Person
        var movies []models.Movie
        var shows  []models.Show
        
        var err error
        json, err := middleware.JsonRequest("https://api.themoviedb.org/3/trending/all/week?language=en-US", "Bearer " + os.Getenv("TMDB"))
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
                var show models.Show
                decoder(result, &show) 
                show.Genre = genreHelper(result)
                show.ExternalIds = extidHelper(result)
                shows = append(shows, show) 
            } else if result["media_type"] == "movie" {
                var movie models.Movie
                decoder(result, &movie)
                movie.Genre = genreHelper(result)
                movie.ExternalIds = extidHelper(result)
                movies = append(movies, movie) 
            } else {
                var jsonPerson map[string]interface{}
                jsonPerson, _ = middleware.JsonRequest("https://api.themoviedb.org/3/person/" + strconv.Itoa(int(result["id"].(float64))) + "?language=en-US", "Bearer " + os.Getenv("TMDB"))
                var person models.Person
                decoder(jsonPerson, &person)
                person.ExternalIds = extidHelper(jsonPerson)
                people = append(people, person) 
            }
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }
        }

        err = models.WritePerson(people)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }
        err = models.WriteMovie(movies)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }
        err = models.WriteShow(shows)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }


        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}

func TMDBMovieDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := strings.Replace(c.Param("movieId"), " ", "%20", -1)
        result, err := middleware.JsonRequest("https://api.themoviedb.org/3/movie/" + query + "?language=en-US", "Bearer " + os.Getenv("TMDB"))
        
        var movie models.Movie
        decoder(result, &movie)
        movie.Genre = genreHelper(result)
        movie.ExternalIds = extidHelper(result)
        
        err = models.WriteMovie([]models.Movie{movie})
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
        json, err := middleware.JsonRequest("https://api.themoviedb.org/3/person/" + query + "?append_to_response=combined_credits&language=en-US", "Bearer " + os.Getenv("TMDB"))

        var person models.Person
        decoder(json, &person)
        person.ExternalIds = extidHelper(json)


        var movies []models.Movie
        var shows  []models.Show
        movieFilters := bson.A{}
        showFilters  := bson.A{}

        credits := json["combined_credits"].(map[string]interface{})["cast"].([]interface{})
        credits = append(credits, json["combined_credits"].(map[string]interface{})["crew"].([]interface{})...)

        for i := 0; i < len(credits); i++ {
            result := credits[i].(map[string]interface{})

            if result["media_type"] == "tv" {
                var show models.Show
                decoder(result, &show) 
                show.Genre = genreHelper(result)
                show.ExternalIds = extidHelper(result)
                showFilters = append(showFilters, bson.M{"title": show.Title, "date": show.Date})
                shows = append(shows, show) 
            } else {
                var movie models.Movie
                decoder(result, &movie)
                movie.Genre = genreHelper(result)
                movie.ExternalIds = extidHelper(result)
                movieFilters = append(movieFilters, bson.M{"title": movie.Title, "date": movie.Date})
                movies = append(movies, movie) 
            } 
        }
        
        err = models.WriteMovie(movies)
        movies, err = models.FindMovie(bson.D{{"$or", movieFilters}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }
        err = models.WriteShow(shows)
        shows, err = models.FindShow(bson.D{{"$or", showFilters}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }

        for i := 0; i < len(movies); i++ {
            person.Movies = append(person.Movies, movies[i].Id)
        }
        for i := 0; i < len(shows); i++ {
            person.Shows = append(person.Shows, shows[i].Id)
        }

        err = models.WritePerson([]models.Person{person})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        } else {
            c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
        }
    }
}

func TMDBShowDetails() gin.HandlerFunc {
    return func(c *gin.Context) {
        query := strings.Replace(c.Param("showId"), " ", "%20", -1)
        result, err := middleware.JsonRequest("https://api.themoviedb.org/3/tv/" + query + "?append_to_response=recommendations&language=en-US", "Bearer " + os.Getenv("TMDB"))

        var show models.Show
        decoder(result, &show) 
        show.Genre = genreHelper(result)
        show.ExternalIds = extidHelper(result)

        err = models.WriteShow([]models.Show{show})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }
        
        showsres, _ := models.FindShow(bson.D{{"title", show.Title}, {"date", show.Date}})
        show = showsres[0]

        var seasons  []models.ShowSeason
        var episodes []models.ShowEpisode

        for i := 1; i <= show.Seasons; i++ {
            seasonResult, err := middleware.JsonRequest("https://api.themoviedb.org/3/tv/" + query + "/season/" + strconv.Itoa(i) + "?language=en-US", "Bearer " + os.Getenv("TMDB"))
            if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }

            var season models.ShowSeason
            decoder(seasonResult, &season) 
            season.ShowId = show.Id

            seasons = append(seasons, season) 
            episodeResults := seasonResult["episodes"].([]interface{})

            for j := 0; j < len(episodeResults); j++ {
                var episode models.ShowEpisode
                decoder(episodeResults[j].(map[string]interface{}), &episode) 
                episode.ShowId = show.Id
                episode.SeasonID = i

                episodes = append(episodes, episode) 
            }
        }

        err = models.WriteShowSeason(seasons)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }
        err = models.WriteShowEpisode(episodes)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
        }

        c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
}
