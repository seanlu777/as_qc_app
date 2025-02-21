package db

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(dsn string) error {
	err := initializeDBConnection(dsn)
	if err != nil {
		return fmt.Errorf("failed to initialize db connection: %v", err)
	}
	err = runMigrations()
	if err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}
	return nil
}

// initializeDBConnection initializes the DB connection
func initializeDBConnection(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errors.Is(err, gorm.ErrInvalidDB) || errors.Is(err, gorm.ErrRecordNotFound) {
		// Create the database if it does not exists
		if err := createDatabaseIfNotExists(dsn); err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to open connection todatabase: %v", err)
		}
	}
	return nil
}

func createDatabaseIfNotExists(dsn string) error {
	// Create a new connection to the default database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open connection to default database: %v", err)
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	// Create the database
	result := db.Exec("CREATE DATABASE " + os.Getenv("DB_NAME"))
	if result.Error != nil {
		return fmt.Errorf("failed to create database: %v", result.Error)
	}
	fmt.Println("Successfully created database")
	return nil
}

// runMigrations runs all database migrations
func runMigrations() error {
	err := DB.AutoMigrate(&Test{})
	if err != nil {
		return fmt.Errorf("failed to migrate db: %v", err)
	}

	// Ensure User table exists
	if !DB.Migrator().HasTable(&Test{}) {
		if err := DB.AutoMigrate(&Test{}); err != nil {
			return fmt.Errorf("failed to migrate User table: %v", err)
		}
	}
	return nil
}
