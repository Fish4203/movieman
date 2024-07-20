package controllers

import (
	"backend-mediaman/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bulkMedia struct {
  models.Movie
  models.MovieExternal
}

func BulkAdd() gin.HandlerFunc {
  return func(c *gin.Context) {
    var bulkMedia bulkMedia

    if err := c.BindJSON(&bulkMedia); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    c.JSON(http.StatusCreated, map[string]interface{}{"moviesExternal": bulkMedia.MovieExternal, "movie": bulkMedia.Movie})
  }
}

