package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShowEpisode struct {
  Title       string          `json:"title" `
  
  Number      uint            `json:"number"  gorm:"not null;primaryKey"` 
  Date        string          `json:"date"    gorm:"not null;primaryKey"`
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt  `             gorm:"index"`

  Budget        uint              `json:"budget"`
  Rating        string            `json:"rating"`
  
  ExternalInfo  []ShowSeasonExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Number,Date;References:Number,Date"` 
  Review        []ShowSeasonReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Number,Date;References:Number,Date"` 
}

type ShowEpisodeExternal struct {
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt  `             gorm:"index"`

  Number      uint        
  Date        string      

  ExternalID        string    `json:"externalId"      gorm:"primaryKey"`
  DataProviderID    uint      `json:"dataProvider"    gorm:"primaryKey"`
  
  WatchPlatforms    []string  `json:"watchPlatforms"  gorm:"serializer:json"`
  Genre             []string  `json:"genre"           gorm:"serializer:json"`
  Links             []string  `json:"links"           gorm:"serializer:json"`
  Description       string    `json:"description"`
  ReviewScore       string    `json:"reviewScore"` 
}

type ShowEpisodeReview struct {
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt  `             gorm:"index"`

  Number      uint            `             gorm:"primaryKey"`
  Date        string          `             gorm:"primaryKey"`
  UserID      uint            `json:"user"  gorm:"primaryKey"`

  Progress    uint            `json:"progress"`
  Rating      uint            `json:"rating"`
  Notes       string          `json:"notes"`
}

func (episode ShowEpisode) GetTitle() string {
  return fmt.Sprint(episode.Number)
}

func (media *ShowEpisode) SetTitle(value string) {
  (*media).Title = value
}

func (review ShowEpisodeReview) GetUserId() uint {
  return review.UserID
}

func (review *ShowEpisodeReview) SetUserId(value uint) {
  (*review).UserID = value
}

func (showEpisode *ShowEpisode) Save(c *gin.Context) error {
  return saveMedia(c, showEpisode)
}

func (showEpisode *ShowEpisode) Get(c *gin.Context) error {
  return getMedia(c, showEpisode)
}

func (showEpisode *ShowEpisode) Delete(c *gin.Context) error {
  return deleteMedia(c, showEpisode)
}

func (showEpisode *ShowEpisodeReview) Save(c *gin.Context) error {
  return saveReview(c, showEpisode)
}

func (showEpisode *ShowEpisodeReview) Get(c *gin.Context) error {
  return getReview(c, showEpisode)
}

func (showEpisode *ShowEpisodeReview) Delete(c *gin.Context) error {
  return deleteReview(c, showEpisode)
}
