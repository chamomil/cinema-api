package db

import (
	"cinema/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func Init() {
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	DB.AutoMigrate(&models.Movie{}, &models.Genre{}, &models.User{}, &models.Feedback{})
}

func Start() {
	Init()
	Migrate()
}
