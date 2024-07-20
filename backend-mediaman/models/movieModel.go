package models

import "github.com/gin-gonic/gin"

type Movie struct {
  Media

  Budget        uint              `json:"budget"`
  Length        uint              `json:"length"`
  Rating        string            `json:"rating"`
  ExternalInfo  []MovieExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
  Review        []MovieReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
}

type MovieExternal struct {
  MediaExternal
}

type MovieReview struct {
  MediaReview
}

type MovieUnion struct {
  Movie
  MovieExternal
}

func (movie MovieUnion) GetExternal() MovieExternal {
  external := movie.MovieExternal
  external.Title = movie.Movie.Title
  external.Date  = movie.Movie.Date

  return external
}

func (movie *Movie) Save(c *gin.Context) error {
  return saveMedia(c, movie)
}

func (movie *Movie) Get(c *gin.Context) error {
  return getMedia(c, movie)
}

func (movie *Movie) Delete(c *gin.Context) error {
  return deleteMedia(c, movie)
}

func (movie *MovieReview) Save(c *gin.Context) error {
  return saveReview(c, movie)
}

func (movie *MovieReview) Get(c *gin.Context) error {
  return getReview(c, movie)
}

func (movie *MovieReview) Delete(c *gin.Context) error {
  return deleteReview(c, movie)
}
