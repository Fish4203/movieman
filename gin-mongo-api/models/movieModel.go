package models

import (
  "gorm.io/gorm"
)


type Movie struct {
  gorm.Model
  // basic info
  Title         string            `json:"title"             gorm:"not null"`
  Date          string            `json:"date"              gorm:"not null"`
  Budget        uint              `json:"budget,omitempty"`
  Length        uint              `json:"length,omitempty"`
  Rating        string            `json:"rating,omitempty"`
  ExternalInfo  []MovieExternal   
  Review        []MovieReview
}

type MovieExternal struct {
  gorm.Model
  MovieID           uint
  DataProviderID    uint

  ExternalPlatform  string    `json:"external_platform"       gorm:"not null"`

  Genre             []string  `json:"genre,omitempty"`
  Description       string    `json:"description,omitempty"`
  ReviewScore       string    `json:"review_score,omitempty"` 
  Platforms         []string  `json:"platforms,omitempty"`
  Links             []string  `json:"links,omitempty"`
}

type MovieReview struct {
  gorm.Model
  UserID        uint
  MovieID       uint

  WatchPercent  uint
  Rating        uint
  Notes         string
}
