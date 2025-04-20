package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nrmadi02/go_react_monorepo/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	tz := os.Getenv("DB_TIMEZONE")

	dsn := "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + name +
		" port=" + port +
		" sslmode=disable TimeZone=" + tz

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}
	DB = database
}

func MigrateDatabase() {

	log.Println("Migrating database...")

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed!", err)
	}
}
