package controllers

import (
	"TMDB-api/middleware"
	"TMDB-api/models"
	"fmt"
	"net/http"
	"strconv"

	"os"
	"strings"

	"github.com/gin-gonic/gin"
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

func extidHelper(json map[string]interface{}) map[string]string {
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

		json, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/authentication", "Bearer "+os.Getenv("TMDB"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, json)
	}
}

func TMDBSearch() gin.HandlerFunc {
	return func(c *gin.Context) {

		var companies []models.Company
		var groups []models.Group
		var people []models.Person
		var movies []models.Movie
		var shows []models.Show

		types := c.Query("types")
		query := strings.Replace(c.Query("q"), " ", "%20", -1)
		var err error

		if strings.ContainsAny(types, "m") || strings.ContainsAny(types, "s") || strings.ContainsAny(types, "p") {
			json, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/search/multi?query="+query+"&language=en-US&page=1", "Bearer "+os.Getenv("TMDB"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}

			var result map[string]interface{}

			for i := 0; i < len(json["results"].([]interface{})); i++ {
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
					jsonPerson, _ = middleware.JsonRequestGet("https://api.themoviedb.org/3/person/"+strconv.Itoa(int(result["id"].(float64)))+"?language=en-US", "Bearer "+os.Getenv("TMDB"))
					var person models.Person
					decoder(jsonPerson, &person)
					person.ExternalIds = extidHelper(jsonPerson)
					people = append(people, person)
				}
				if err != nil {
					c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				}
			}
		}

		if strings.ContainsAny(types, "g") {
			json, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/search/company?query="+query+"&language=en-US&page=1", "Bearer "+os.Getenv("TMDB"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}

			var result map[string]interface{}

			for i := 0; i < len(json["results"].([]interface{})); i++ {
				result = json["results"].([]interface{})[i].(map[string]interface{})

				var company models.Company
				decoder(result, &company)
				company.ExternalIds = extidHelper(result)
				companies = append(companies, company)

			}
		}

		if strings.ContainsAny(types, "c") {
			json, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/search/collection?query="+query+"&include_adult=true&language=en-US&page=1", "Bearer "+os.Getenv("TMDB"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}

			var result map[string]interface{}

			for i := 0; i < len(json["results"].([]interface{})); i++ {
				result = json["results"].([]interface{})[i].(map[string]interface{})

				var group models.Group
				decoder(result, &group)
				group.ExternalIds = extidHelper(result)
				groups = append(groups, group)

			}
		}

		var out = make(map[string]interface{})
		out["linked"] = false
		out["movies"] = movies
		out["shows"] = shows
		out["people"] = people
		out["groups"] = groups
		out["companies"] = companies

		err = middleware.JsonRequestPost(out)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
	}
}

func TMDBPopular() gin.HandlerFunc {
	return func(c *gin.Context) {
		var people []models.Person
		var movies []models.Movie
		var shows []models.Show

		var err error
		json, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/trending/all/week?language=en-US", "Bearer "+os.Getenv("TMDB"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		var result map[string]interface{}

		for i := 0; i < len(json["results"].([]interface{})); i++ {
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
				jsonPerson, _ = middleware.JsonRequestGet("https://api.themoviedb.org/3/person/"+strconv.Itoa(int(result["id"].(float64)))+"?language=en-US", "Bearer "+os.Getenv("TMDB"))
				var person models.Person
				decoder(jsonPerson, &person)
				person.ExternalIds = extidHelper(jsonPerson)
				people = append(people, person)
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}
		}

		var out = make(map[string]interface{})
		out["linked"] = false
		out["movies"] = movies
		out["shows"] = shows
		out["people"] = people

		err = middleware.JsonRequestPost(out)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
	}
}

func TMDBMovieDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := strings.Replace(c.Param("movieId"), " ", "%20", -1)
		result, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/movie/"+query+"?language=en-US", "Bearer "+os.Getenv("TMDB"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		var movie models.Movie
		decoder(result, &movie)
		movie.Genre = genreHelper(result)
		movie.ExternalIds = extidHelper(result)

		var out = make(map[string]interface{})
		out["linked"] = false
		out["movies"] = []models.Movie{movie}

		err = middleware.JsonRequestPost(out)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
	}
}

func TMDBPersonDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := strings.Replace(c.Param("personId"), " ", "%20", -1)
		json, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/person/"+query+"?append_to_response=combined_credits&language=en-US", "Bearer "+os.Getenv("TMDB"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		var person models.Person
		decoder(json, &person)
		person.ExternalIds = extidHelper(json)

		var movies []models.Movie
		var shows []models.Show

		credits := json["combined_credits"].(map[string]interface{})["cast"].([]interface{})
		credits = append(credits, json["combined_credits"].(map[string]interface{})["crew"].([]interface{})...)

		for i := 0; i < len(credits); i++ {
			result := credits[i].(map[string]interface{})

			if result["media_type"] == "tv" {
				var show models.Show
				decoder(result, &show)
				show.Genre = genreHelper(result)
				show.ExternalIds = extidHelper(result)
				shows = append(shows, show)
			} else {
				var movie models.Movie
				decoder(result, &movie)
				movie.Genre = genreHelper(result)
				movie.ExternalIds = extidHelper(result)
				movies = append(movies, movie)
			}
		}

		var out = make(map[string]interface{})
		out["linked"] = true
		out["movies"] = movies
		out["shows"] = shows
		out["people"] = []models.Person{person}

		err = middleware.JsonRequestPost(out)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
	}
}

func TMDBCollectionDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := strings.Replace(c.Param("Id"), " ", "%20", -1)
		json, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/collection/"+query+"&language=en-US", "Bearer "+os.Getenv("TMDB"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		var group models.Group
		decoder(json, &group)
		group.ExternalIds = extidHelper(json)

		var movies []models.Movie
		var shows []models.Show

		credits := json["parts"].([]interface{})

		for i := 0; i < len(credits); i++ {
			result := credits[i].(map[string]interface{})

			if result["media_type"] == "tv" {
				var show models.Show
				decoder(result, &show)
				show.Genre = genreHelper(result)
				show.ExternalIds = extidHelper(result)
				shows = append(shows, show)
			} else {
				var movie models.Movie
				decoder(result, &movie)
				movie.Genre = genreHelper(result)
				movie.ExternalIds = extidHelper(result)
				movies = append(movies, movie)
			}
		}

		var out = make(map[string]interface{})
		out["linked"] = true
		out["movies"] = movies
		out["shows"] = shows
		out["groups"] = []models.Group{group}

		err = middleware.JsonRequestPost(out)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
	}
}

func TMDBShowDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := strings.Replace(c.Param("showId"), " ", "%20", -1)
		result, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/tv/"+query+"?append_to_response=recommendations&language=en-US", "Bearer "+os.Getenv("TMDB"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		var show models.Show
		decoder(result, &show)
		show.Genre = genreHelper(result)
		show.ExternalIds = extidHelper(result)

		var seasons []models.ShowSeason
		var episodes []models.ShowEpisode

		for i := 1; i <= show.Seasons; i++ {
			seasonResult, err := middleware.JsonRequestGet("https://api.themoviedb.org/3/tv/"+query+"/season/"+strconv.Itoa(i)+"?language=en-US", "Bearer "+os.Getenv("TMDB"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}

			var season models.ShowSeason
			decoder(seasonResult, &season)

			seasons = append(seasons, season)
			episodeResults := seasonResult["episodes"].([]interface{})

			for j := 0; j < len(episodeResults); j++ {
				var episode models.ShowEpisode
				decoder(episodeResults[j].(map[string]interface{}), &episode)
				episode.SeasonID = i

				episodes = append(episodes, episode)
			}
		}

		var out = make(map[string]interface{})
		out["linked"] = false
		out["shows"] = []models.Show{show}
		out["showSeasons"] = [][]models.ShowSeason{seasons}
		out["showEpisodes"] = [][]models.ShowEpisode{episodes}

		err = middleware.JsonRequestPost(out)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
	}
}
