package main

import (
	"flag"
	"log"
	"path/filepath"
)

var archiveFile string
var outputDir string

func main() {
	flag.Parse()

	err := ReadConfig(configFile)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	archiveFile = C.General["archive_file"]
	outputDir = C.General["output_dir"]

	SetDownloader()
	Download()
	for podcast := range C.Podcasts {
		DownloadThumbnail(podcast)
		GenerateFeed(filepath.Join(outputDir, podcast))
	}
}
