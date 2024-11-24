package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
  CreatedAt     time.Time
  UpdatedAt     time.Time
  DeletedAt     gorm.DeletedAt    `                                                       gorm:"index"`
  
  ID            uint              `json:"id"                                              gorm:"primaryKey"`
  
  Title         string            `json:"title"               binding:"required"          gorm:"index"` 
  Date          string            `json:"date"                binding:"required"`
  Budget        *uint             `json:"budget,omitempty"`
  Length        *uint             `json:"length,omitempty"`
  
  Externals     []MovieExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:MovieID;References:ID"` 
  Reviews       []MovieReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:MovieID;References:ID"` 
}

type MovieExternal struct {
  CreatedAt         time.Time
  UpdatedAt         time.Time
  DeletedAt         gorm.DeletedAt                                                       `gorm:"index"`

  MovieID           *uint         `json:"movieID,omitempty"                               gorm:"index,not null"`

  ExternalID        string        `json:"externalID"          binding:"required"          gorm:"primaryKey"`
  DataProviderID    uint          `json:"dataProviderID"      binding:"required"          gorm:"primaryKey"`
 
  WatchPlatforms    []string      `json:"watchPlatforms"                                  gorm:"serializer:json"`
  Genre             []string      `json:"genre"                                           gorm:"serializer:json"`
  Links             []string      `json:"links"               binding:"dive,url"          gorm:"serializer:json"`
  Description       string        `json:"description"         binding:"required"`
  ReviewScore       uint          `json:"reviewScore"         binding:"required,lte=100"` 
  AgeRating         *string       `json:"ageRating,omitempty"`
}

type MovieReview struct {
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt      `                                                       gorm:"index"`

  MovieID     uint                `json:"movieID"             binding:"required"          gorm:"primaryKey"`
  UserID      uint                `json:"userID"              binding:"required"          gorm:"primaryKey"`

  Rating      uint                `json:"rating"              binding:"required,lte=100"`
  Progress    uint                `json:"progress"            binding:"required,lte=100"`
  Notes       string              `json:"notes"               binding:"required"`
}

type MovieUnion struct {
  Movie         Movie             `json:"movie"               binding:"required"`
  MovieExternal MovieExternal     `json:"external"            binding:"required"`
}


