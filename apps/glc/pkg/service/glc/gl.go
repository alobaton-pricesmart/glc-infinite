package glc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"glc-infinite/apps/glc/pkg/glc"
	"glc-infinite/pkg/httperror"
)

type GlcService struct {
}

func NewGlcService() *GlcService {
	return &GlcService{}
}

func (gc GlcService) IsInfite(w http.ResponseWriter, r *http.Request) error {
	var glc glc.GLC

	err := json.NewDecoder(r.Body).Decode(&glc)
	if err != nil {
		return httperror.NewHTTPError(err, 400, "Bad request: invalid JSON")
	}

	finite, err := glc.IsFinite()
	if err != nil {
		return httperror.NewHTTPError(err, 400, "Error in the provided struct")
	}

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
