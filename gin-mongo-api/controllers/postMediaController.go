package controllers

import (
	"gin-mongo-api/middleware"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BulkAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := middleware.ExtractToken(c)

		if tokenString != os.Getenv("APITOKEN") {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "invald api token"})
			return
		}

		var request responses.BulkRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := models.WriteMovie(request.Movies); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if err := models.WriteShow(request.Shows); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if len(request.ShowSeasons) > 0 && len(request.ShowSeasons) <= len(request.Shows) {
			var filter bson.A
			var seasons []models.ShowSeason
			var episodes []models.ShowEpisode

			for i := 0; i < len(request.Shows); i++ {
				filter = append(filter, bson.M{"title": request.Shows[i].Title, "date": request.Shows[i].Date})
			}

			showsres, err := models.FindShow(bson.D{{"$or", filter}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}

			for i := 0; i < len(request.ShowSeasons); i++ {
				for j := 0; j < len(request.ShowSeasons[i]); j++ {
					request.ShowSeasons[i][j].ShowId = showsres[i].Id
				}
				seasons = append(seasons, request.ShowSeasons[i]...)

				for j := 0; j < len(request.ShowEpisodes[i]); j++ {
					request.ShowEpisodes[i][j].ShowId = showsres[i].Id
				}
				episodes = append(episodes, request.ShowEpisodes[i]...)
			}

			if err := models.WriteShowSeason(seasons); err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}
			if err := models.WriteShowEpisode(episodes); err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}
		}
		if err := models.WriteBook(request.Books); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if err := models.WriteGame(request.Games); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}

		// adds the movies and shows to the people/ groups/ companies if relevent
		if request.Linked {
			// defs
			var movies []primitive.ObjectID
			var shows []primitive.ObjectID
			var books []primitive.ObjectID
			var games []primitive.ObjectID
			var filterMovies bson.A
			var filterShows bson.A
			var filterBooks bson.A
			var filterGames bson.A

			// adding all the media to the tilters
			for i := 0; i < len(request.Movies); i++ {
				filterMovies = append(filterMovies, bson.M{"title": request.Movies[i].Title, "date": request.Movies[i].Date})
			}
			for i := 0; i < len(request.Shows); i++ {
				filterShows = append(filterShows, bson.M{"title": request.Shows[i].Title, "date": request.Shows[i].Date})
			}
			for i := 0; i < len(request.Books); i++ {
				filterBooks = append(filterBooks, bson.M{"title": request.Books[i].Title, "date": request.Books[i].Date})
			}
			for i := 0; i < len(request.Games); i++ {
				filterGames = append(filterGames, bson.M{"title": request.Games[i].Title, "date": request.Games[i].Date})
			}

			// searching the database for the media
			movieres, err := models.FindMovie(bson.D{{"$or", filterMovies}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}
			showres, err := models.FindShow(bson.D{{"$or", filterShows}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}
			bookres, err := models.FindBook(bson.D{{"$or", filterBooks}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}
			gameres, err := models.FindGame(bson.D{{"$or", filterGames}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}

			// converting the search results into lists of ids
			for i := 0; i < len(movieres); i++ {
				movies = append(movies, movieres[i].Id)
			}
			for i := 0; i < len(showres); i++ {
				shows = append(shows, showres[i].Id)
			}
			for i := 0; i < len(bookres); i++ {
				books = append(books, bookres[i].Id)
			}
			for i := 0; i < len(gameres); i++ {
				games = append(games, gameres[i].Id)
			}

			// adding the media ids to the people, companies, groups
			for i := 0; i < len(request.People); i++ {
				request.People[i].Movies = movies
				request.People[i].Shows = shows
				request.People[i].Books = books
				request.People[i].Games = games
			}
			for i := 0; i < len(request.Companies); i++ {
				request.Companies[i].Movies = movies
				request.Companies[i].Shows = shows
				request.Companies[i].Books = books
				request.Companies[i].Games = games
			}
			for i := 0; i < len(request.Groups); i++ {
				request.Groups[i].Movies = movies
				request.Groups[i].Shows = shows
				request.Groups[i].Books = books
				request.Groups[i].Games = games
			}
		}
		if err := models.WritePerson(request.People); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if err := models.WriteCompany(request.Companies); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if err := models.WriteGroup(request.Groups); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, nil)
	}
}

func EditMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie models.Movie

		if err := c.BindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&movie); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := models.WriteMovie([]models.Movie{movie}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}

func EditShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		var show models.Show

		if err := c.BindJSON(&show); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&show); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := models.WriteShow([]models.Show{show}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}

func EditShowSeason() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj models.ShowSeason

		if err := c.BindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := models.WriteShowSeason([]models.ShowSeason{obj}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}

func EditShowEpisode() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj models.ShowEpisode

		if err := c.BindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := models.WriteShowEpisode([]models.ShowEpisode{obj}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}

func EditPerson() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj models.Person

		if err := c.BindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := models.WritePerson([]models.Person{obj}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}

func EditBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj models.Book

		if err := c.BindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := models.WriteBook([]models.Book{obj}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}

func EditGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj models.Game

		if err := c.BindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&obj); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		if err := models.WriteGame([]models.Game{obj}); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}
