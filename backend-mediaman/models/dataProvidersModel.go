package models

import (
	"time"

	"gorm.io/gorm"
)

type DataProvider struct {
  CreatedAt     time.Time
  UpdatedAt     time.Time
  DeletedAt     gorm.DeletedAt    `                                                   gorm:"index"`
  
  ID            uint              `json:"id"                                          gorm:"primaryKey"`
  
  Name          string            `json:"name"                binding:"required"      gorm:"unique;index"`
  Description   string            `json:"description"         binding:"required"`
  BaseUrl       string            `json:"base_url,omitempty"  binding:"required,url"`

  HasMovies     bool              `json:"hasMovies"     binding:"required"`
  Movies        []MovieExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:DataProviderID;References:ID"` 
}
