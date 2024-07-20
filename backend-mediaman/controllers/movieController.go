package controllers

import (
	"backend-mediaman/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
