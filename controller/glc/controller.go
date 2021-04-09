package glc

import (
	"encoding/json"
	"fmt"
	"net/http"

	g "glc-infinite/glc"
	e "glc-infinite/handler/err"

	"github.com/gorilla/mux"
)

type GlcController struct {
	router *mux.Router
}

func NewGlcController(router *mux.Router) *GlcController {
	return &GlcController{router}
}

func (gc GlcController) IsInfite(w http.ResponseWriter, r *http.Request) error {
	var glc g.GLC

	err := json.NewDecoder(r.Body).Decode(&glc)
	if err != nil {
		return e.NewHTTPError(err, 400, "Bad request: invalid JSON")
	}

	finite := glc.IsFinite()

	response := make(map[string]interface{})
	response["finite"] = finite

	body, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("Response body write error : %v", err)
	}

	w.WriteHeader(200)
	w.Write(body)
	return nil
}

func (gc GlcController) Handle() {
	gc.router.Handle("/glc/isFinite", e.ErrorHandler(gc.IsInfite)).Methods("POST")
}
