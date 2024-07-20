package controllers

import (
	"backend-mediaman/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGame() gin.HandlerFunc {
  return func(c *gin.Context) {
    var game models.Game

    if err := game.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"game": game})
    }
  }
}

func GetGames() gin.HandlerFunc {
  return func(c *gin.Context) {
    var games []*models.Game
    var game models.Game

    if err := models.SearchMedia(c, &game, &games); err != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": err})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"games": games})
    }
  }
}

func DeleteGame() gin.HandlerFunc {
  return func(c *gin.Context) {
    var game models.Game

    if err := game.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"game": game})
    }
  }
}

func CreateGameReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var gameReview models.GameReview

    if err := gameReview.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"review": gameReview})
    } 
  }
}

func GetGameReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var gameReview models.GameReview

    if err := gameReview.Get(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"review": gameReview})
    } 
  }
}

func DeleteGameReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var gameReview models.GameReview

    if err := gameReview.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"status": "deleted"})
    } 
  }
}
