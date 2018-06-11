package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var archiveFile string
var configFile string
var outputDir string

func init() {
	configDefault, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), "config.yml"))
	if err != nil {
		fmt.Println(err)
		return
	}
	archiveDefault, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), "archive"))
	if err != nil {
		fmt.Println(err)
		return
	}
	outputDirDefault, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), "output_dir"))
	if err != nil {
		fmt.Println(err)
		return
	}
	flag.StringVar(&archiveFile, "archive", archiveDefault, "archive file")
	flag.StringVar(&configFile, "config", configDefault, "configuration file")
	flag.StringVar(&outputDir, "output", outputDirDefault, "output directory")
}
