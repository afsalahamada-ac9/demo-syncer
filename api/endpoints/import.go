package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sudhagar/glad/entity"
)

func ImportSFData(w http.ResponseWriter, r *http.Request) {
	// sf_api := "https://aol-dev--awspoc.sandbox.my.salesforce.com/services/data/v55.0/sobjects/Event__c"
	// token, err := api.GenerateTokens()
	// if err != nil {
	// 	log.Println("error generating the tokens")
	// }
	// req, err := http.NewRequest("GET", sf_api, nil)
	// if err != nil {
	// 	log.Println(err)
	// }
	// body := "Bearer " + token
	// req.Header.Set("Authorization", body)
	// client := http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println("error executing the request", err)
	// }
	parse, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("error parsing the response", err)
	}
	var result []entity.SF
	err = json.Unmarshal(parse, &result)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(result)
}
