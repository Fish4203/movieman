package controllers

import (
	"backend-mediaman/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShowEpisode() gin.HandlerFunc {
  return func(c *gin.Context) {
    var episode models.ShowEpisode

    if err := episode.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"episode": episode})
    }
  }
}

func GetShowEpisodes() gin.HandlerFunc {
  return func(c *gin.Context) {
    var episodes []*models.ShowEpisode
    var episode models.ShowEpisode

    if err := models.SearchMedia(c, &episode, &episodes); err != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": err})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"episodes": episodes})
    }
  }
}

func DeleteShowEpisode() gin.HandlerFunc {
  return func(c *gin.Context) {
    var episode models.ShowEpisode

    if err := episode.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"episode": episode})
    }
  }
}

func CreateShowEpisodeReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var episodeReview models.ShowEpisodeReview

    if err := episodeReview.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"review": episodeReview})
    } 
  }
}

func GetShowEpisodeReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var episodeReview models.ShowEpisodeReview

    if err := episodeReview.Get(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"review": episodeReview})
    } 
  }
}

func DeleteShowEpisodeReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var episodeReview models.ShowEpisodeReview

    if err := episodeReview.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"status": "deleted"})
    } 
  }
}
