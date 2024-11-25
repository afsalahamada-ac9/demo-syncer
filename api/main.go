/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package main

import (
	"log"
	"net/http"
	api "sudhagar/glad/api/endpoints"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sync/import", api.ImportSFData)
	router.HandleFunc("/sync/export", api.ExportRDSData)
	log.Println(http.ListenAndServe(":4010", router))
}
