package database

import (
	"fmt"
	"log"
	"os"
	"technical-test-II/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// koneksi db MySQL dengan GORM
func StartDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbport, dbname)
	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(domain.User{}, domain.Transaction{}, domain.Token{})
}

func GetDB() *gorm.DB {
	return db
}
