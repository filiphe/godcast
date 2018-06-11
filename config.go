package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config holds information about multiple podcasts
// and general configuration.
type Config struct {
	Podcasts map[string]Podcast `yaml:"podcasts"`
	General  map[string]string  `yaml:"general"`
}

func ReadConfig(fp string) (*Config, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c := Config{}
	err = yaml.Unmarshal(data, &c)
	return &c, nil
}
