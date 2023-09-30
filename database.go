package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func handleDatabaseConnection() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "root"
	dbname := "hacktiv8_assignment3"

	connectionString := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Error connecting to database", err)
	}

	log.Println("Database connection established")
}

func GetDatabaseConnection() *gorm.DB {
	if db == nil {
		handleDatabaseConnection()
	}

	return db
}
