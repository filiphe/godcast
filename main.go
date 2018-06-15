package main

import (
	"flag"
	"log"
	"path/filepath"
)

func main() {
	flag.Parse()

	err := ReadConfig(configFile)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	SetDownloader()
	Download()
	for podcast := range C.Podcasts {
		DownloadThumbnail(podcast)
		GenerateFeed(filepath.Join(outputDir, podcast))
	}
}
