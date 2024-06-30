package models

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name          string            `gorm:"not null"`
  Email         string            `gorm:"not null"`
  Password      string            `gorm:"not null"`
  Role          string            `gorm:"not null"`

  MovieReview   []MovieReview
}
