package controllers

import (
	"backend-mediaman/configs"
	"backend-mediaman/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieReview models.MovieReview

    if err := c.ShouldBind(movieReview); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    userID := c.GetInt("userID")
    role := c.GetString("role")
    
    if userID != int(movieReview.UserID) && role != "admin" {
      c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "review userID dosn't match token userID, are you trying to edit another users review?"})
      return 
    }
    
    result := configs.DB.Save(&movieReview)
    if result.Error != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": result.Error.Error()})
      return
    }
    
    c.JSON(http.StatusCreated, map[string]interface{}{"review": movieReview})
  }
}

func GetMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieReviews []models.MovieReview
    var result *gorm.DB 

    movieID, err := strconv.Atoi(c.Param("movieID"))
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    userID, err := strconv.Atoi(c.Query("userID"))
    if err != nil {
      result = configs.DB.Where(&models.MovieReview{MediaReview: models.MediaReview{MediaID: uint(movieID)}}).Find(&movieReviews)
    } else {
      result = configs.DB.Where(&models.MovieReview{MediaReview: models.MediaReview{MediaID: uint(movieID), UserID: uint(userID)}}).Find(&movieReviews)    
    }

    if result.Error != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": result.Error.Error()})
      return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"reviews": movieReviews}) 
  }
}

func DeleteMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    movieID, err := strconv.Atoi(c.Param("movieID"))
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    userIDParam, err := strconv.Atoi(c.Param("userID"))
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    userIDToken := c.GetInt("userID")
    role := c.GetString("role")
    
    if userIDParam != userIDToken && role != "admin" {
      c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "review userID dosn't match token userID, are you trying to delete another users review?"})
      return 
    }

    result := configs.DB.Where(&models.MovieReview{MediaReview: models.MediaReview{MediaID: uint(movieID), UserID: uint(userIDParam)}}).Delete(&models.MovieReview{})    

    if result.Error != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": result.Error.Error()})
      return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"deletedReviews": result.RowsAffected})  
  }
}

func GetMovie() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movie []models.Movie

    movieID, err := strconv.Atoi(c.Param("movieID"))
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    result := configs.DB.Preload("ExternalInfo").Preload("Review").First(&movie, movieID)
    if result.Error != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": result.Error.Error()})
      return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"movie": movie}) 
  }
}

func DeleteMovie() gin.HandlerFunc {
  return func(c *gin.Context) {
    movieID, err := strconv.Atoi(c.Param("movieID"))
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    result := configs.DB.Delete(&models.Movie{}, movieID)
    if result.Error != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": result.Error.Error()})
      return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"deletedMovies": result.RowsAffected})
  }
}

func CreateMovieReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var movieUnion models.MovieUnion

    if err := c.ShouldBind(movieUnion); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    
  }
}



