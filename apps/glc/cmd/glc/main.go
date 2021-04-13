package main

import (
	log "github.com/sirupsen/logrus"

	"glc-infinite/apps/glc/pkg/config"
	"glc-infinite/apps/glc/pkg/service/glc"
	"glc-infinite/pkg/middleware"
	"glc-infinite/pkg/service"

	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
	glcService := glc.NewGlcService()
	r.Handle("/isFinite", middleware.ErrorHandler(glcService.IsInfite)).Methods("POST")
}

// main function program
func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	s := service.Init("glc")
	s.ConfigureRouting()

	registerRoutes(s.RoutingRouter())

	s.Run()
}
