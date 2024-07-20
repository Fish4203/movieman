package controllers

import (
	"backend-mediaman/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShow() gin.HandlerFunc {
  return func(c *gin.Context) {
    var show models.Show

    if err := show.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"show": show})
    }
  }
}

func GetShows() gin.HandlerFunc {
  return func(c *gin.Context) {
    var shows []*models.Show
    var show models.Show

    if err := models.SearchMedia(c, &show, &shows); err != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": err})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"shows": shows})
    }
  }
}

func DeleteShow() gin.HandlerFunc {
  return func(c *gin.Context) {
    var show models.Show

    if err := show.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"show": show})
    }
  }
}

func CreateShowReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var showReview models.ShowReview

    if err := showReview.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"review": showReview})
    } 
  }
}

func GetShowReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var showReview models.ShowReview

    if err := showReview.Get(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"review": showReview})
    } 
  }
}

func DeleteShowReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var showReview models.ShowReview

    if err := showReview.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"status": "deleted"})
    } 
  }
}
