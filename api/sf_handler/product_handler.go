package sf_handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	tapi "sudhagar/glad/api/tapi"
	test_entity "sudhagar/glad/entity/sf_entity"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	var response []test_entity.Product
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("there was an error reading the body")
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Println("error in unmarshal process", err)
	}
	defer r.Body.Close()
	log.Println("response:", string(resp))
	for _, record := range response {
		_, err := tapi.WriteToDB(&record)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(record)
		}
	}

}
