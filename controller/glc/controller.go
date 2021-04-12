package glc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"glc-infinite/glc"
	"glc-infinite/httperrors"
)

type GlcController struct {
}

func NewGlcController() *GlcController {
	return &GlcController{}
}

func (gc GlcController) IsInfite(w http.ResponseWriter, r *http.Request) error {
	var glc glc.GLC

	err := json.NewDecoder(r.Body).Decode(&glc)
	if err != nil {
		return httperrors.NewHTTPError(err, 400, "Bad request: invalid JSON")
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
