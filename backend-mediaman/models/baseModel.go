package models

import (
	"backend-mediaman/configs"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type Media struct {
  Title       string          `json:"title" gorm:"primaryKey"`
  Date        string          `json:"date"  gorm:"primaryKey"`
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt  `             gorm:"index"`
}

type MediaExternal struct {
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt  `             gorm:"index"`

  Title       string        
  Date        string      

  ExternalID        string    `json:"externalId"      gorm:"primaryKey"`
  DataProviderID    uint      `json:"dataProvider"    gorm:"primaryKey"`
  
  WatchPlatforms    []string  `json:"watchPlatforms"  gorm:"serializer:json"`
  Genre             []string  `json:"genre"           gorm:"serializer:json"`
  Links             []string  `json:"links"           gorm:"serializer:json"`
  Description       string    `json:"description"`
  ReviewScore       string    `json:"reviewScore"` 
}

type MediaReview struct {
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt  `             gorm:"index"`

  Title       string          `json:"title" gorm:"primaryKey"`
  Date        string          `json:"date"  gorm:"primaryKey"`
  UserID      uint            `json:"user"  gorm:"primaryKey"`

  Rating      uint            `json:"rating"`
  Progress    uint            `json:"progress"`
  Notes       string          `json:"notes"`
}

type MediaInterface interface {
  GetTitle() string
  SetTitle(value string)
  Save(c *gin.Context) error
  Delete(c *gin.Context) error
  Get(c *gin.Context) error
}

type ReviewInterface interface {
  Save(c *gin.Context) error
  Delete(c *gin.Context) error
  Get(c *gin.Context) error
  GetUserId() uint
  SetUserId(value uint) 
}

func (media Media) GetTitle() string {
  return media.Title
}

func (media *Media) SetTitle(value string) {
  (*media).Title = value
}

func (review MediaReview) GetUserId() uint {
  return review.UserID
}

func (review *MediaReview) SetUserId(value uint) {
  (*review).UserID = value
}

func saveMedia(c *gin.Context, media MediaInterface) error {
  if err := c.BindJSON(media); err != nil {
    return err
  }

  result := configs.DB.Save(media)
  return result.Error
}

func getMedia(c *gin.Context, media MediaInterface) error {
  if err := c.BindJSON(media); err != nil {
    return err
  }

  result := configs.DB.Preload(clause.Associations).First(media, media)
  return result.Error
}

func deleteMedia(c *gin.Context, media MediaInterface) error {
  if err := c.BindJSON(media); err != nil {
    return err
  }

  result := configs.DB.Delete(media)
  return result.Error
}

func SearchMedia[T MediaInterface](c *gin.Context, media T, mediaArr *[]T) error {
  if err := c.BindJSON(media); err != nil {
    return err
  }

  title := media.GetTitle()
  media.SetTitle("")
  
  result := configs.DB.Preload(clause.Associations).Where("title LIKE ?", "%" + title + "%").Where(&media).Find(mediaArr)
  return result.Error
}

func saveReview(c *gin.Context, review ReviewInterface) error {
  if err := c.BindJSON(review); err != nil {
    return err
  }

  userId := c.GetUint("user_id")
  if userId != review.GetUserId() {
    return errors.New("Can't edit another users review")
  }

  result := configs.DB.Save(review)
  return result.Error
}

func getReview(c *gin.Context, review ReviewInterface) error {
  if err := c.BindJSON(review); err != nil {
    return err
  }

  result := configs.DB.Preload(clause.Associations).First(review, review)
  return result.Error
}

func deleteReview(c *gin.Context, review ReviewInterface) error {
  if err := c.BindJSON(review); err != nil {
    return err
  }

  userId := c.GetUint("user_id")
  if userId != review.GetUserId() {
    return errors.New("Can't delete another users review")
  }

  result := configs.DB.Delete(review)
  return result.Error
}

