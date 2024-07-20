package models

import "github.com/gin-gonic/gin"

type ShowSeason struct {
  Media

  SeasonDate    string            `json:"seasonDate"    gorm:"not null"`
  SeasonTitle   string            `json:"seasonTitle"`
  SeasonNumber  uint              `json:"seasonNumber"  gorm:"primaryKey"`

  Budget        uint              `json:"budget"`
  Rating        string            `json:"rating"`
  
  ExternalInfo  []ShowSeasonExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date,SeasonNumber;References:Title,Date,SeasonNumber"` 
  Review        []ShowSeasonReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date,SeasonNumber;References:Title,Date,SeasonNumber"` 
 
  Episodes      []ShowEpisode          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date,SeasonNumber;References:Title,Date,SeasonNumber"` 
}

type ShowSeasonExternal struct {
  SeasonNumber  uint              `json:"seasonNumber"  gorm:"primaryKey"`
  MediaExternal
}

type ShowSeasonReview struct {
  SeasonNumber  uint              `json:"seasonNumber"  gorm:"primaryKey"`
  MediaReview
}

func (season *ShowSeason) Save(c *gin.Context) error {
  return saveMedia(c, season)
}

func (season *ShowSeason) Get(c *gin.Context) error {
  return getMedia(c, season)
}

func (season *ShowSeason) Delete(c *gin.Context) error {
  return deleteMedia(c, season)
}

func (season *ShowSeasonReview) Save(c *gin.Context) error {
  return saveReview(c, season)
}

func (season *ShowSeasonReview) Get(c *gin.Context) error {
  return getReview(c, season)
}

func (season *ShowSeasonReview) Delete(c *gin.Context) error {
  return deleteReview(c, season)
}
