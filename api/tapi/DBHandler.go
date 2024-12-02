package tapi

// todo: write the data to rds db
//  todo: <entity> operations are done in rds,
import (
	"log"
	ops "sudhagar/glad/ops/db"
)

func WriteToDB(record any) (string, error) {
	db, err := ops.GetDB()
	if err != nil {
		log.Println("there is an error fetching the db", err)
	}
	if db == nil {
		log.Println("db is nil")
	}
	log.Println("inserting record now:", record)
	result := db.Create(record)
	if result.Error != nil {
		log.Println("error occurred in the write process", result.Error)
		return "", result.Error
	}
	return "success", nil
}
