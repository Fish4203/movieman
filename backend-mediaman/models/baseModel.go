package models

import (
	"backend-mediaman/configs"
	"errors"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type Media struct {
  Title       string          `json:"title" gorm:"not null;primaryKey"`
  Date        string          `json:"date"  gorm:"not null;primaryKey"`
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

  Title       string          `             gorm:"primaryKey"`
  Date        string          `             gorm:"primaryKey"`
  UserID      uint            `json:"user"  gorm:"primaryKey"`

  Progress    uint            `json:"progress"`
  Rating      uint            `json:"rating"`
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

func saveMedia(c *gin.Context, media MediaInterface) error {
  if err := c.BindJSON(media); err != nil {
    return err
  }

  if result := configs.DB.Save(media); result.Error != nil {
    return result.Error
  }

  return nil
}

func getMedia(c *gin.Context, media MediaInterface) error {
  if err := c.BindJSON(media); err != nil {
    return err
  }

  if result := configs.DB.Preload(clause.Associations).First(media, media); result.Error != nil {
    return result.Error
  }

  return nil
}

func deleteMedia(c *gin.Context, media MediaInterface) error {
  if err := c.BindJSON(media); err != nil {
    return err
  }

  if result := configs.DB.Delete(media); result.Error != nil {
    return result.Error
  }

  return nil
}

func searchMedia(c *gin.Context, media MediaInterface) (interface{}, error) {
  mediaArr := reflect.MakeSlice(reflect.TypeOf(media).Elem(), 0, 0)

  if err := c.BindJSON(media); err != nil {
    return mediaArr, err
  }

  title := media.GetTitle()
  media.SetTitle("")
    
  result := configs.DB.Preload(clause.Associations).Where("title LIKE ?", "%" + title + "%").Find(mediaArr)

  return mediaArr, result.Error
}

func saveReview(c *gin.Context, review ReviewInterface) error {
  if err := c.BindJSON(review); err != nil {
    return err
  }

  userId := c.GetUint("user_id")
  if userId != review.GetUserId() {
    return errors.New("Can't edit another users review")
  }

  if result := configs.DB.Save(review); result.Error != nil {
    return result.Error
  }

  return nil
}

func getReview(c *gin.Context, review ReviewInterface) error {
  if err := c.BindJSON(review); err != nil {
    return err
  }

  if result := configs.DB.Preload(clause.Associations).First(review, review); result.Error != nil {
    return result.Error
  }

  return nil
}

func deleteReview(c *gin.Context, review ReviewInterface) error {
  if err := c.BindJSON(review); err != nil {
    return err
  }

  userId := c.GetUint("user_id")
  if userId != review.GetUserId() {
    return errors.New("Can't delete another users review")
  }

  if result := configs.DB.Delete(review); result.Error != nil {
    return result.Error
  }

  return nil
}

