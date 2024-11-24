package controllers

import (
	"backend-mediaman/configs"
	"backend-mediaman/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bulkMedia struct {
  Movies  []models.MovieUnion
  //Books   []models.BookUnion
  //Games   []models.GameUnion
  //Shows   []models.ShowUnion
}

func CreateBulk() gin.HandlerFunc {
  return func(c *gin.Context) {
    var media bulkMedia

    if err := c.ShouldBind(media); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    var dbErrors []error

    for i := 0; i < len(media.Movies); i++ {
      var movieData     models.Movie            = media.Movies[i].Movie
      var externalData  models.MovieExternal    = media.Movies[i].MovieExternal
      
      var movie     models.Movie
      var external  models.MovieExternal

      result := configs.DB. 
        Where(&models.MovieExternal{ExternalID: externalData.ExternalID, DataProviderID: externalData.DataProviderID}).
        Attrs(externalData).
        FirstOrCreate(&external)

      if result.Error != nil {
        dbErrors = append(dbErrors, result.Error)
        continue
      }

      if external.MovieID == nil {
        result = configs.DB.
          Where(&models.Movie{Title: movieData.Title, Date: movieData.Date}).
          Attrs(movieData).
          FirstOrCreate(&movie)

        *external.MovieID = movie.ID
      
        if result.Error != nil {
          dbErrors = append(dbErrors, result.Error)
          continue
        }
      }

      result = configs.DB.Save(&external)
    
      if result.Error != nil {
        dbErrors = append(dbErrors, result.Error)
        continue
      }
    }
  }
}


