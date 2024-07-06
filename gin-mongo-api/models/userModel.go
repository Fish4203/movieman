package models

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name          string            `gorm:"not null,unique"`
  Email         string            `gorm:"not null,unique"`
  Password      string            `gorm:"not null"`
  Role          string            `gorm:"not null"`

  MovieReview   []MovieReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
