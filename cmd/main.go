package main

import (
	"log"
	"net/http"

	"glc-infinite/controller/glc"
	"glc-infinite/handler/header"
	"glc-infinite/handler/recover"

	"github.com/gorilla/mux"
)

// main function program
func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(header.HeaderHandler)
	router.Use(recover.RecoverHandler)

	gc := glc.NewGlcController(router)
	gc.Handle()

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
