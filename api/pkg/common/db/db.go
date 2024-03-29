package db

import (
	"log"

	"github.com/b0nbon1/debt-roundup/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
  db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

  if err != nil {
    log.Fatalln(err)
  }

  db.AutoMigrate(&models.Debt{})
  db.AutoMigrate(&models.User{})

  return db
}

