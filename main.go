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
		downloadLink := fmt.Sprintf("%s%s", c.General["playlist_base"], value.PlaylistID)
		cmd, err := ydl.Download(downloadLink)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		cmd.Wait()
		for _, ID := range value.AdditionalEpisodes {
			downloadLink = fmt.Sprintf("%s%s", c.General["video_base"], ID)
			cmd, err = ydl.Download(downloadLink)
			if err != nil {
				log.Fatalf("%+v\n", err)
			}
			cmd.Wait()
		}
	}
}
