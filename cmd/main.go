package main

import (
	"log"

	"glc-infinite/controller"
)

// main function program
func main() {
	err := controller.Routes()
	if err != nil {
		log.Fatal(err)
	}
}
