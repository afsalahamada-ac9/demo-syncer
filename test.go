package main

import (
	"log"
	"net/http"
	handler "sudhagar/glad/api/sf_handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/account", handler.AccountHandler)
	log.Println("now listening at port 4001")
	log.Println(http.ListenAndServe(":4001", router))
}
