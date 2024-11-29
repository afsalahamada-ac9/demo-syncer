package sf_handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sudhagar/glad/api/tapi"
	"sudhagar/glad/entity"
)

func CourseHandler(w http.ResponseWriter, r *http.Request) {
	var courses []entity.Course
	parsed_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("there was an error parsing the body")
	}
	err = json.Unmarshal(parsed_body, &courses)
	if err != nil {
		log.Println("there was an error unmarshalling the body")
	}
	for _, course := range courses {
		_, err := tapi.WriteToDB(&course)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(course)
		}
	}
	log.Println(courses)
}
