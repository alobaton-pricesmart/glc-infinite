package main

import (
	"log"
	"net/http"

	"glc-infinite/controller/glc"

	"github.com/gorilla/mux"
)

// main function program
func main() {
	router := mux.NewRouter().StrictSlash(true)

	gc := glc.NewGlcController(router)
	gc.Handle()

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
