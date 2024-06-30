package configs

import (
	"context"
	"fmt"
	"gin-mongo-api/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitClient() (*gorm.DB)   {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  db, err := gorm.Open(postgres.New(postgres.Config{
    DSN: fmt.Sprintf("user=apiClient password=%s dbname=movieman host=%s port=9432 sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST")),
    PreferSimpleProtocol: true, // disables implicit prepared statement usage
  }), &gorm.Config{})
  if err != nil {
    log.Fatal("Error loading db")
  }

  db.AutoMigrate(&models.User{}, &models.Movie{}, &models.MovieExternal{}, &models.MovieReview{})

  return db
}
