package database

import (
	"fmt"
	"log"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "zulham"
	dbPort   = "5433"
	dbName   = "pglocal"
	db       *gorm.DB
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	var err error
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Terjadi Kesalahan saat koneksi ke db :", err)
	}

	fmt.Println("Berhasil tersambung ke db")
	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
