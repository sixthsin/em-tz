package migrations

import (
	"log"
	"os"
	"sub-aggregator/internal/subs"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AutoMigrate() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&subs.Subscription{})
	log.Println("Migrated successfully")
}
