package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

  

func GetDB() (*gorm.DB) {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  db, err := gorm.Open(postgres.New(postgres.Config{
    DSN: fmt.Sprintf("user=apiClient password=%s dbname=movieman host=%s port=9432 sslmode=disable ", os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST")),
    PreferSimpleProtocol: true, // disables implicit prepared statement usage
  }), &gorm.Config{})
  if err != nil {
    log.Fatal("Error loading db")
  }

  return db
}

var DB *gorm.DB = GetDB()
