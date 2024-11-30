package sf_handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sudhagar/glad/api/tapi"
	"sudhagar/glad/entity"
)

func TenantHandler(w http.ResponseWriter, r *http.Request) {
	var tenants []entity.Tenant
	parsed_response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	err = json.Unmarshal(parsed_response, &tenants)
	if err != nil {
		log.Println("there was an error unmarshalling the response")
	}
	defer r.Body.Close()
	for _, record := range tenants {
		_, err := tapi.WriteToDB(&record)
		if err == nil {
			json.NewEncoder(w).Encode(record)
			log.Println("insertion successful")
		} else {
			json.NewEncoder(w).Encode(err)
		}
	}
	log.Println(tenants)
}
