package models

import "github.com/gin-gonic/gin"

type ShowEpisode struct {
  Media

  EpisodeDate   string            `json:"episodeDate"    gorm:"not null"`
  EpisodeNumber uint              `json:"episodeNumber"  gorm:"not null;primaryKey"`
  EpisodeTitle  string            `json:"episodeTitle"`

  SeasonDate    string            `json:"seasonDate"    gorm:"not null"`
  SeasonNumber  uint              `json:"seasonNumber"  gorm:"not null;primaryKey"`
  SeasonTitle   string            `json:"seasonTitle"`

  Budget        uint              `json:"budget"`
  Rating        string            `json:"rating"`
  
  ExternalInfo  []ShowEpisodeExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date,SeasonNumber,EpisodeNumber;References:Title,Date,SeasonNumber,EpisodeNumber"` 
  Review        []ShowEpisodeReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date,SeasonNumber,EpisodeNumber;References:Title,Date,SeasonNumber,EpisodeNumber"` 
}

type ShowEpisodeExternal struct {
  SeasonNumber  uint              `json:"seasonNumber"  gorm:"not null;primaryKey"`
  EpisodeNumber uint              `json:"episodeNumber"  gorm:"not null;primaryKey"`
  
  MediaExternal
}

type ShowEpisodeReview struct {
  SeasonNumber  uint              `json:"seasonNumber"  gorm:"not null;primaryKey"`
  EpisodeNumber uint              `json:"episodeNumber"  gorm:"not null;primaryKey"`
  
  MediaReview
}

func (episode *ShowEpisode) Save(c *gin.Context) error {
  return saveMedia(c, episode)
}

func (episode *ShowEpisode) Get(c *gin.Context) error {
  return getMedia(c, episode)
}

func (episode *ShowEpisode) Delete(c *gin.Context) error {
  return deleteMedia(c, episode)
}

func (episode *ShowEpisodeReview) Save(c *gin.Context) error {
  return saveReview(c, episode)
}

func (episode *ShowEpisodeReview) Get(c *gin.Context) error {
  return getReview(c, episode)
}

func (episode *ShowEpisodeReview) Delete(c *gin.Context) error {
  return deleteReview(c, episode)
}
