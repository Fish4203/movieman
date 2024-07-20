package models

import "github.com/gin-gonic/gin"

type Show struct {
  Media

  Budget        uint              `json:"budget"`
  Rating        string            `json:"rating"`
  
  ExternalInfo  []ShowExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
  Review        []ShowReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
  
  Seasons       []ShowSeason     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"`
  Episodes      []ShowEpisode    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"`
}

type ShowExternal struct {
  MediaExternal
}

type ShowReview struct {
  MediaReview
}

type ShowUnion struct {
  Show
  ShowExternal

  BulkSeasons       []ShowSeasonUnion  `json:"seasons"` 
}

func (show ShowUnion) GetExternal() ShowExternal {
  external := show.ShowExternal
  external.Title = show.Show.Title
  external.Date  = show.Show.Date

  return external
}

func (show *Show) Save(c *gin.Context) error {
  return saveMedia(c, show)
}

func (show *Show) Get(c *gin.Context) error {
  return getMedia(c, show)
}

func (show *Show) Delete(c *gin.Context) error {
  return deleteMedia(c, show)
}

func (show *ShowReview) Save(c *gin.Context) error {
  return saveReview(c, show)
}

func (show *ShowReview) Get(c *gin.Context) error {
  return getReview(c, show)
}

func (show *ShowReview) Delete(c *gin.Context) error {
  return deleteReview(c, show)
}