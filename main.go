package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()

	c, err := ReadConfig(configFile)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	SetDownloader()
	for key, value := range c.Podcasts {
		ydl.Options.Output.Value = fmt.Sprintf("%s/%s/%%(upload_date)s-%%(title)s.%%(ext)s", outputDir, key)
		download_link := fmt.Sprintf("%s%s", c.General["playlist_base"], value.PlaylistID)
		fmt.Println(download_link)
		cmd, err := ydl.Download(download_link)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		cmd.Wait()
	}
}
