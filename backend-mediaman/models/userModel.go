package models

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name          string          `json:"name,omitempty"      binding:"required"          gorm:"unique;not null"`
  Email         *string         `json:"email,omitempty"     binding:"omitempty,email"   gorm:"unique;not null"`
  Password      *string         `json:"password,omitempty"                              gorm:"not null"`
  Role          string          `json:"role,omitempty"                                  gorm:"not null"`

  MovieReview   []MovieReview   `json:"movieReview" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
