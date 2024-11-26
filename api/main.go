/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	api "sudhagar/glad/api/endpoints"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sync/import", api.ImportSFData)
	router.HandleFunc("/sync/export", api.ExportRDSData)
	PORT := 4010
	url := fmt.Sprintf(":%d", PORT)
	log.Println("listening at port", PORT)
	log.Println(http.ListenAndServe(url, router))
}
