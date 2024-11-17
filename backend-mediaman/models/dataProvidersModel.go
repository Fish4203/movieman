package models

import (
  "gorm.io/gorm"
)

type DataProvider struct {
  gorm.Model
  Name          string    `                   binding:"required"`
  Description   string    `                   binding:"required"`
  BaseUrl       string    `json:"base_url"    binding:"required,url"`
}
