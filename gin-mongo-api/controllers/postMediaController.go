package controllers

import (
	// "context"
	// "gin-mongo-api/configs"

	"gin-mongo-api/models"
	// "gin-mongo-api/responses"
	// "gin-mongo-api/middleware"
	// "fmt"
	"net/http"
	// "io"
	// "encoding/json"
	// "os"
	// "strings"
	// "compress/gzip"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	// "go.mongodb.org/mongo-driver/mongo"
	// "golang.org/x/crypto/bcrypt"
)

type bultRequest struct {
	Movies       []models.Movie       `json:"movies"`
	Shows        []models.Show        `json:"shows"`
	ShowSeasons  []models.ShowSeason  `json:"showSeasons"`
	ShowEpisodes []models.ShowEpisode `json:"showEpisodes"`
	People       []models.Person      `json:"people"`
	Books        []models.Book        `json:"books"`
	Games        []models.Game        `json:"games"`
}

func BulkAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request bultRequest

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
		if err := models.WriteShowSeason(request.ShowSeasons); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if err := models.WriteShowEpisode(request.ShowEpisodes); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if err := models.WritePerson(request.People); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if err := models.WriteBook(request.Books); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}
		if err := models.WriteGame(request.Games); err != nil {
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
