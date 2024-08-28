package database

import (
	"fmt"
	"menu/model/drinks"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func initDB() {
	// Set up the connection string.
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	fmt.Println("Connecting to database with:", connection)

	// Connect to the database.
	db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// Automatically migrate the schema.
	err = db.AutoMigrate(&drinks.Drinks{})
	if err != nil {
		panic("failed to migrate database schema: " + err.Error())
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		initDB()
	}
	return db
}
