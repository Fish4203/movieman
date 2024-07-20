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

type bulkMedia struct {
  Movies  []models.MovieUnion
  Books   []models.BookUnion
  Games   []models.GameUnion
  Shows   []models.ShowUnion
}

func BulkAdd() gin.HandlerFunc {
  return func(c *gin.Context) {
    var bulkMedia           bulkMedia
    var movies              []models.Movie
    var moviesExternal      []models.MovieExternal
    var books               []models.Book
    var booksExternal       []models.BookExternal
    var games               []models.Game
    var gamesExternal       []models.GameExternal
    var shows               []models.Show
    var showsExternal       []models.ShowExternal
    var seasons             []models.ShowSeason
    var seasonsExternal     []models.ShowSeasonExternal
    var episodes            []models.ShowEpisode
    var episodesExternal    []models.ShowEpisodeExternal

    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    err = json.Unmarshal(body, &bulkMedia)
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    for i := 0; i < len(bulkMedia.Movies); i++ {
      movies = append(movies, bulkMedia.Movies[i].Movie)
      moviesExternal = append(moviesExternal, bulkMedia.Movies[i].GetExternal())
    }

    for i := 0; i < len(bulkMedia.Books); i++ {
      books = append(books, bulkMedia.Books[i].Book)
      booksExternal = append(booksExternal, bulkMedia.Books[i].GetExternal())
    }

    for i := 0; i < len(bulkMedia.Games); i++ {
      games = append(games, bulkMedia.Games[i].Game)
      gamesExternal = append(gamesExternal, bulkMedia.Games[i].GetExternal())
    }

    for i := 0; i < len(bulkMedia.Shows); i++ {
      shows = append(shows, bulkMedia.Shows[i].Show)
      showsExternal = append(showsExternal, bulkMedia.Shows[i].GetExternal())
    
      bulkSeasons := bulkMedia.Shows[i].BulkSeasons
      for j := 0; j < len(bulkSeasons); j++ {
        seasons = append(seasons, bulkSeasons[j].ShowSeason)
        seasonsExternal = append(seasonsExternal, bulkSeasons[j].GetExternal())
     
        bulkEpisodes := bulkSeasons[j].BulkEpisodes
        for k := 0; k < len(bulkEpisodes); k++ {
          episodes = append(episodes, bulkEpisodes[k].ShowEpisode)
          episodesExternal = append(episodesExternal, bulkEpisodes[k].GetExternal())
        }
      }
    }

    if err := createBulkMediaTransaction(
      movies,
      moviesExternal,
      books,
      booksExternal,
      games,
      gamesExternal,
      shows,
      showsExternal,
      seasons,
      seasonsExternal,
      episodes,
      episodesExternal,
    ); err != nil {
      c.JSON(http.StatusCreated, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"result": "success"})
    }
  }
}


func createBulkMediaTransaction(
  movies              []models.Movie,
  moviesExternal      []models.MovieExternal,
  books               []models.Book,
  booksExternal       []models.BookExternal,
  games               []models.Game,
  gamesExternal       []models.GameExternal,
  shows               []models.Show,
  showsExternal       []models.ShowExternal,
  seasons             []models.ShowSeason,
  seasonsExternal     []models.ShowSeasonExternal,
  episodes            []models.ShowEpisode,
  episodesExternal    []models.ShowEpisodeExternal,
  ) error {
  tx := configs.DB.Begin()
  tx.SavePoint("start")

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&movies); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&books); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&games); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&shows); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&seasons); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&episodes); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Commit(); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&moviesExternal); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&booksExternal); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&gamesExternal); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&showsExternal); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&seasonsExternal); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Clauses(clause.OnConflict{ UpdateAll: true }).Create(&episodesExternal); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  if result := tx.Commit(); result.Error != nil {
    tx.RollbackTo("start")
    return result.Error
  }

  return nil
}

