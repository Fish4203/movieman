package controllers

import (
	"backend-mediaman/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShowSeason() gin.HandlerFunc {
  return func(c *gin.Context) {
    var season models.ShowSeason

    if err := season.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"season": season})
    }
  }
}

func GetShowSeasons() gin.HandlerFunc {
  return func(c *gin.Context) {
    var seasons []*models.ShowSeason
    var season models.ShowSeason

    if err := models.SearchMedia(c, &season, &seasons); err != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": err})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"seasons": seasons})
    }
  }
}

func DeleteShowSeason() gin.HandlerFunc {
  return func(c *gin.Context) {
    var season models.ShowSeason

    if err := season.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"season": season})
    }
  }
}

func CreateShowSeasonReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var seasonReview models.ShowSeasonReview

    if err := seasonReview.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"review": seasonReview})
    } 
  }
}

func GetShowSeasonReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var seasonReview models.ShowSeasonReview

    if err := seasonReview.Get(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"review": seasonReview})
    } 
  }
}

func DeleteShowSeasonReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var seasonReview models.ShowSeasonReview

    if err := seasonReview.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"status": "deleted"})
    } 
  }
}
