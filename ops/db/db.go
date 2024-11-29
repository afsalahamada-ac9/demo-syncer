package infra

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	connection_str := "host=localhost user=postgres password=1234 port=5432 dbname=glad sslmode=disable" // todo: move to .env
	db, err := gorm.Open(postgres.Open(connection_str), &gorm.Config{})
	if err != nil {
		log.Println("there was an error with the database", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("there was an error configuring the db")
	}
	// set up connection pooling
	sqlDB.SetMaxIdleConns(10)           // max number of connections in idle connection pool
	sqlDB.SetMaxOpenConns(100)          // max number of open connections
	sqlDB.SetConnMaxLifetime(time.Hour) // max amount of time a connection may be reused
	if err != nil {
		log.Println("there was an error connecting to the database")
		return nil, err
	}
	return db, nil
}
