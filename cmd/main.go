package main

import (
	"net/http"

	"glc-infinite/controller/glc"
)

// main function program
func main() {
	mux := http.NewServeMux()

	gc := glc.NewGlcController()
	gc.Handle(mux)

	http.ListenAndServe(":8080", mux)
}
