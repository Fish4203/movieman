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

func EditMovie() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movie models.Movie

    if err := c.BindJSON(&movie); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    if result := configs.DB.Save(&movie); result.Error != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error})
      return
    }
        
    c.JSON(http.StatusCreated, map[string]interface{}{"movie": movie})
  }
}

func GetMovies() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movies []models.Movie
    var movie models.Movie

    if err := c.BindJSON(&movie); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    title := movie.Title
    movie.Title = ""
    
    if result := configs.DB.Preload(clause.Associations).Where("title LIKE ?", "%" + title + "%").Find(&movies); result.Error != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
       return
    }
        
    c.JSON(http.StatusOK, map[string]interface{}{"movies": movies})
  }
}

func DeleteMovie() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movie models.Movie

    if err := c.BindJSON(&movie); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    if result := configs.DB.Delete(&movie); result.Error != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
       return
    }
    
    c.JSON(http.StatusOK, map[string]interface{}{"result": "Movie successfully deleted"})
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
    userID := c.MustGet("userId").(uint)
    
    if err := c.BindJSON(&movieReview); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    movieReview.UserID = userID

    if result := configs.DB.Create(&movieReview); result.Error != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error})
      return
    }
        
    c.JSON(http.StatusCreated, map[string]interface{}{"movieReview": movieReview})
  }
}

func EditMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieReview models.MovieReview
    userID := c.MustGet("userId").(uint)
    
    if err := c.BindJSON(&movieReview); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    movieReview.UserID = userID

    if result := configs.DB.Save(&movieReview); result.Error != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error})
      return
    }
        
    c.JSON(http.StatusCreated, map[string]interface{}{"movieReview": movieReview})
  }
}

func GetMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieReview models.MovieReview

    if err := c.BindJSON(&movieReview); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    if result := configs.DB.Where(&movieReview).First(&movieReview); result.Error != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
       return
    }
        
    c.JSON(http.StatusOK, map[string]interface{}{"movieReview": movieReview})
  }
}

func DeleteMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieReview models.MovieReview

    if err := c.BindJSON(&movieReview); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    if result := configs.DB.Delete(&movieReview); result.Error != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
       return
    }
    
    c.JSON(http.StatusOK, map[string]interface{}{"result": "Movie Review successfully deleted"})
  }
}
