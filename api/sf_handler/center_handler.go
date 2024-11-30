package sf_handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sudhagar/glad/api/tapi"
	entity "sudhagar/glad/entity/sf_entity"
)

func CenterHandler(w http.ResponseWriter, r *http.Request) {
	var centers []entity.Center
	parsed_response, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("an error occurred")
	}
	err = json.Unmarshal(parsed_response, &centers)
	if err != nil {
		log.Println("an error occurred in the unmarshalling opf the centers")
	}
	for _, record := range centers {
		value := record.Value
		_, err := tapi.WriteToDB(record.NewCenter(value.Ext_id, value.Tenant_id, value.Ext_name, value.Address, value.Geo_Location, value.Capacity, value.Mode, value.Webpage, value.Is_national_center, value.Is_enabled, value.Created_at, value.Updated_at))
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(record)
		}
	}
	log.Println(centers)
}
