package infra

import (
	"log"
	"sudhagar/glad/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// use embedded and embeddedPrefix from gorm to allow it to easily parse embedded structs, otherwise it will create separate tables for each of them
func DbModifier(data entity.EventValue) {
	connection_str := "host=localhost user=postgres password=1234 port=5432 dbname=glad sslmode=disable" // todo: move to .env
	db, err := gorm.Open(postgres.Open(connection_str), &gorm.Config{})
	if err != nil {
		log.Println("there was an error connecting to the database")
	}
	err = db.AutoMigrate(&data)
	if err != nil {
		log.Println("failed to migrate", err)
	}
	db.Create(data)
}
