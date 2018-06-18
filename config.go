package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config holds information about multiple podcasts
// and general configuration.
type Config struct {
	Podcasts map[string]Podcast `yaml:"podcasts"`
	General  map[string]string  `yaml:"general"`
}

// C holds a read config
var C *Config

func ReadConfig(fp string) error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("%+v\n", err)
		return err
	}
	err = yaml.Unmarshal(data, &C)
	return nil
}
