package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	api "sudhagar/glad/api/services"
)

// todo: the below function is returning an error from SF side, check with it.
func ExportRDSData(w http.ResponseWriter, r *http.Request) {
	// sendToSf, err := ioutil.ReadAll(r.Body)
	sf_api := "https://aol-dev--awspoc.sandbox.my.salesforce.com/services/apexrest/handleAolEvent"
	var jsonData map[string]interface{}
	//jsonData, err := json.Marshal(string(sendToSf)) // since jsonData is of type []byte, we've to parse it as *bytes.Buffer which implements io.Reader(which is the expected type of the body)
	// log.Println("this is the json data:", string(jsonData))
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		log.Println("there was an error decoding the object in the request", err)
	}
	parsed_json, err := json.Marshal(jsonData)
	if err != nil {
		log.Println("there was an error parsing the json, check parsed_json", err)
	}
	log.Println("parsed json:", parsed_json)
	req, err := http.NewRequest("POST", sf_api, bytes.NewBuffer(parsed_json)) // todo: fix now
	if err != nil {
		log.Println("error creating the request")
	}
	AUTH_TOKEN, err := api.GenerateTokens()
	if err != nil {
		log.Println("error generating the tokens")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+string(AUTH_TOKEN))
	log.Println(AUTH_TOKEN)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("there was an error in the request", err)
	}
	fmt.Printf("token type: %T", AUTH_TOKEN)
	var SfResult string
	parse, err := ioutil.ReadAll(resp.Body)
	log.Println("parse:", string(parse))
	if err != nil {
		log.Println("error decoding the object", err)
	}
	err = json.Unmarshal(parse, &SfResult)
	if err != nil {
		log.Println(err)
	}
	log.Println("data was sent successfully")
	json.NewEncoder(w).Encode(SfResult)
}
