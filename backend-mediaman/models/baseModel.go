package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type Media struct {
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt      `                                                   gorm:"index"`
  
  ID          uint                `json:"id"                                          gorm:"primaryKey"`
  
  Title       string              `json:"title"           binding:"required"          gorm:"not null"` 
  Date        string              `json:"date"            binding:"required"          gorm:"not null"`
}

type MediaExternal struct {
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt      `                                                   gorm:"index"`

  MediaID           uint          `json:"mediaID"                                     gorm:"index"`

  ExternalID        string        `json:"externalID"      binding:"required"          gorm:"primaryKey"`
  DataProviderID    uint          `json:"dataProvider"    binding:"required"          gorm:"primaryKey"`
 
  DataProvider      DataProvider  `                                                   gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;ForeignKey:ID;References:DataProviderID"`

  WatchPlatforms    []string      `json:"watchPlatforms"                              gorm:"serializer:json"`
  Genre             []string      `json:"genre"                                       gorm:"serializer:json"`
  Links             []string      `json:"links"           binding:"dive,url"          gorm:"serializer:json"`
  Description       string        `json:"description"     binding:"required"`
  ReviewScore       uint          `json:"reviewScore"     binding:"required,lte=100"` 
}

type MediaReview struct {
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt  `                                                   gorm:"index"`

  MediaID     uint            `json:"mediaID"         binding:"required"          gorm:"primaryKey"`
  UserID      uint            `json:"userID"          binding:"required"          gorm:"primaryKey"`

  Rating      uint            `json:"rating"          binding:"required,lte=100"`
  Progress    uint            `json:"progress"        binding:"required,lte=100"`
  Notes       string          `json:"notes"           binding:"required"`
}

type FullMediaInterface interface {
  GetTitle() string
  SetTitle(value string)
  GetDate() string
  SetDate(value string)
  GetExternalID() string
  SetExternalID(value string)
  GetDataProvider() uint
  SetDataProvider(value uint)
}

type ReviewInterface interface {
  GetUserID() uint
  SetUserID(value uint)
  GetMediaID() uint
  SetMediaID(value uint)

  Save(c *gin.Context) error
  Delete(c *gin.Context) error
  Get(c *gin.Context) error   
}

// media methods 
func (media Media) GetTitle() string {
  return media.Title
}

func (media *Media) SetTitle(value string) {
  (*media).Title = value
}

func (media Media) GetDate() string {
  return media.Date
}

func (media *Media) SetDate(value string) {
  (*media).Date = value
}

func (media Media) GetID() uint {
  return media.ID
}

func (media *Media) SetID(value uint) {
  (*media).ID = value
}

// review methods 
func (review MediaReview) GetUserId() uint {
  return review.UserID
}

func (review *MediaReview) SetUserId(value uint) {
  (*review).UserID = value
}

func (review MediaReview) GetMediaID() uint {
  return review.MediaID
}

func (review *MediaReview) SetMediaID(value uint) {
  (*review).MediaID = value
}


