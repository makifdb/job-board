package database

import (
	"job-site/pkg/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB gorm connector
var db *gorm.DB
var err error

// ConnectDB connect to db
func ConnectDB(url string) *gorm.DB {
	db, err = gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  url,
				PreferSimpleProtocol: true,
			}), &gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	log.Println("Database connected")

	return db
}

// Migrate migrate db
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entities.Organization{},
		&entities.User{},
		&entities.Location{},
		&entities.Job{},
	)
}

// GetDB get db
func GetDB() *gorm.DB {
	return db
}
