package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShowSeason struct {
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

  Episodes      []ShowEpisode         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Number,Date;References:Number,Date"`
}

type ShowSeasonExternal struct {
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

type ShowSeasonReview struct {
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

func (season ShowSeason) GetTitle() string {
  return fmt.Sprint(season.Number)
}

func (media *ShowSeason) SetTitle(value string) {
  (*media).Title = value
}

func (review ShowSeasonReview) GetUserId() uint {
  return review.UserID
}

func (review *ShowSeasonReview) SetUserId(value uint) {
  (*review).UserID = value
}

func (showSeason *ShowSeason) Save(c *gin.Context) error {
  return saveMedia(c, showSeason)
}

func (showSeason *ShowSeason) Get(c *gin.Context) error {
  return getMedia(c, showSeason)
}

func (showSeason *ShowSeason) Delete(c *gin.Context) error {
  return deleteMedia(c, showSeason)
}

func (showSeason *ShowSeasonReview) Save(c *gin.Context) error {
  return saveReview(c, showSeason)
}

func (showSeason *ShowSeasonReview) Get(c *gin.Context) error {
  return getReview(c, showSeason)
}

func (showSeason *ShowSeasonReview) Delete(c *gin.Context) error {
  return deleteReview(c, showSeason)
}
