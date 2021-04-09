package glc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	g "glc-infinite/glc"
	e "glc-infinite/handler/err"
	"glc-infinite/handler/header"
	"glc-infinite/handler/recover"
)

type GlcController struct {
}

func NewGlcController() *GlcController {
	return &GlcController{}
}

func (gc GlcController) IsInfite(w http.ResponseWriter, r *http.Request) error {
	file, _, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Error getting file : %v", err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return fmt.Errorf("Error reading file : %v", err)
	}

	glc := g.GLC{}
	err = json.Unmarshal(buf.Bytes(), &glc)
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

func (gc GlcController) Handle(mux *http.ServeMux) {
	mux.Handle("/glc/isFinite", header.HeaderHandler(recover.RecoverHandler(e.ErrorHandler(gc.IsInfite))))
}
