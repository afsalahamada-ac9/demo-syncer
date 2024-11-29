package sf_handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	tapi "sudhagar/glad/api/tapi"
	test_entity "sudhagar/glad/entity/sf_entity"
)

func TimingHandler(w http.ResponseWriter, r *http.Request) {
	var response []test_entity.Timing
	parse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("there was an error parsing the request body", err)
	}
	defer r.Body.Close()
	err = json.Unmarshal(parse, &response)
	if err != nil {
		log.Println("there was an error unmarshalling the request body", err)
	}
	for _, record := range response {
		_, err := tapi.WriteToDB(&record)
		if err == nil {
			json.NewEncoder(w).Encode(record)
		} else {
			json.NewEncoder(w).Encode(err)
		}
	}
	json.NewEncoder(w).Encode(response)
}
