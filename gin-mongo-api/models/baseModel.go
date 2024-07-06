package models

import (
	"time"

	"gorm.io/gorm"
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

  Title       string      `gorm"not null"`
  Date        string      `gorm"not null"`

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
