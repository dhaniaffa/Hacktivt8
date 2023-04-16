package database

import (
	"fmt"
	"log"
	"myGram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbName   = "myGram"
	db       *gorm.DB
	err      error
)

func StartDatabase() {
	conf := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	dsn := conf

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Koneksi ke database error, ", err)
	}

	fmt.Println("Koneksi database sukses")

	db.Debug().AutoMigrate(&models.User{}, &models.Comment{}, &models.Photo{}, &models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
