package controller

import (
	"net/http"

	"glc-infinite/controller/glc"
	"glc-infinite/middleware"

	"github.com/gorilla/mux"
)

func Routes() error {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.HeaderHandler)
	router.Use(middleware.RecoverHandler)

	gc := glc.NewGlcController()
	router.Handle("/glc/isFinite", middleware.ErrorHandler(gc.IsInfite)).Methods("POST")

	return http.ListenAndServe(":8080", router)
}
