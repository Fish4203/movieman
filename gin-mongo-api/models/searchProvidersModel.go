package models

import (
  "gorm.io/gorm"
)

type DataProvider struct {
  gorm.Model
  Name          string            `gorm:"not null"`
  Description   string            `gorm:"not null"`

  MovieExternal []MovieExternal
}
