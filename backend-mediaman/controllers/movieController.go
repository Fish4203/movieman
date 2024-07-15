package controllers

import (
	"backend-mediaman/configs"
	"backend-mediaman/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreateMovie() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movie models.Movie

    if err := movie.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"movie": movie})
    }
  }
}

func GetMovies() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movies []*models.Movie
    var movie models.Movie

    if err := models.SearchMedia(c, &movie, &movies); err != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": err})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"movies": movies})
    }
  }
}

func DeleteMovie() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movie models.Movie

    if err := movie.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"movie": movie})
    }
  }
}

type bulkMovies struct {
  Movies        []models.Movie
  ExternalInfo  []models.MovieExternal
}

func BulkMovie() gin.HandlerFunc {
  return func(c *gin.Context) {
    var bulkMovies bulkMovies

    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    err = json.Unmarshal(body, &bulkMovies)
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    resultMovie := configs.DB.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&bulkMovies.Movies)
    if resultMovie.Error != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": resultMovie.Error})
       return
    }

    resultExternal := configs.DB.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&bulkMovies.ExternalInfo)
    if resultExternal.Error != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": resultExternal.Error})
       return
    }

    c.JSON(http.StatusCreated, map[string]interface{}{"moviesInserted": resultMovie.RowsAffected, "externalInfoInserted": resultExternal.RowsAffected})
  }
}

func CreateMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieReview models.MovieReview

    if err := movieReview.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"review": movieReview})
    } 
  }
}

func GetMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieReview models.MovieReview

    if err := movieReview.Get(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"review": movieReview})
    } 
  }
}

func DeleteMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieReview models.MovieReview

    if err := movieReview.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"status": "deleted"})
    } 
  }
}
