package controllers

import (
	"encoding/json"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreateMovie() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movie models.Movie

    if err := c.BindJSON(&movie); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    if result := configs.DB.Create(&movie); result.Error != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error})
      return
    }
        
    c.JSON(http.StatusCreated, map[string]interface{}{"movie": movie})
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

    if result := configs.DB.Where(&movie).Find(&movies); result.Error != nil {
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
