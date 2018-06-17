package main

import (
	"flag"
	"log"
	"os/user"
	"path/filepath"
)

var configDefault string
var configFile string

func init() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return
	}
	configDefault, err := filepath.Abs(filepath.Join(currentUser.HomeDir, ".config", "godcast", "config.yml"))
	if err != nil {
		log.Fatal(err)
		return
	}
	flag.StringVar(&configFile, "config", configDefault, "configuration file")
}
