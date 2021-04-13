package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/juju/errors"
	"gopkg.in/yaml.v2"

	"glc-infinite/pkg/service"
)

var Settings SettingsRoot

type SettingsRoot struct {
	Database Database `yaml:"database"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
}

func ParseSettings() error {
	path := ".apps/glc/config/config.yml"
	if service.IsLocal() {
		path = "apps/glc/config/config.dev.yml"
	}

	p, err := filepath.Abs(path)
	if err != nil {
		return errors.Trace(err)
	}

	f, err := os.Open(p)
	if err != nil {
		return errors.Trace(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Trace(err)
	}

	if err := yaml.Unmarshal(content, &Settings); err != nil {
		return errors.Trace(err)
	}

	return nil
}
