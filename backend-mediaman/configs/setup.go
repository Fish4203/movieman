package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

  

func GetDB() (*gorm.DB) {
  db, err := gorm.Open(postgres.New(postgres.Config{
    DSN: fmt.Sprintf("user=postgres dbname=movieman host=%s port=5432 sslmode=disable ", os.Getenv("DB_HOST")),
    PreferSimpleProtocol: true, // disables implicit prepared statement usage
  }), &gorm.Config{})
  if err != nil {
    log.Fatal("Error loading db")
  }

  return db
}

var DB *gorm.DB = GetDB()
