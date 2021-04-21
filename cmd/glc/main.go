package main

import (
	"encoding/json"
	"glc-infinite/pkg/glc"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// main function program
func main() {
	args := os.Args[1:]
	fileName := args[0]

	if len(fileName) == 0 {
		log.Error("fileName is required")
		return
	}

	log.WithField("fileName", fileName).Info("reading the file...")

	// read the file
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.WithField("error", err).Error("error reading the file")
		return
	}

	// Just for testing purpose...
	// log.WithField("body", string(body)).Info()

	// Mapeamos el archivo en una estructura.
	g := glc.GLC{}
	err = json.Unmarshal([]byte(body), &g)
	if err != nil {
		log.WithField("error", err).Error("error unmarshaling the file")
		return
	}

	// Just for testing purpose...
	// log.WithField("dfa", d).Info()

	// Validamos si el GLC es finito o infinito.
	finite, err := g.IsFinite()
	if err != nil {
		log.WithField("error", err).Error("error executing g.IsFinite()")
		return
	}

	log.WithField("finite", finite).Info("result")
}
