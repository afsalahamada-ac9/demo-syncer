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

func ExportRDSData(w http.ResponseWriter, r *http.Request) {
	sendToSf, err := ioutil.ReadAll(r.Body)
	sf_api := "https://aol-dev--awspoc.sandbox.my.salesforce.com/services/apexrest/handleAolEvent"
	jsonData, err := json.Marshal(sendToSf) // since jsonData is of type []byte, we've to parse it as *bytes.Buffer which implements io.Reader(which is the expected type of the body)
	//log.Println(string(sendToSf), string(jsonData))
	if err != nil {
		log.Println("there is an error in the input file", err)
	}
	req, err := http.NewRequest("POST", sf_api, bytes.NewBuffer(jsonData))
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
	var SfResult map[string]interface{}
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
