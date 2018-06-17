package main

import (
	"fmt"
	"log"

	"github.com/brianallred/goydl"
)

var ydl = goydl.NewYoutubeDl()

// SetDownloader sets goydl options
func SetDownloader() {
	ydl.Options.AddMetadata.Value = true
	ydl.Options.AudioFormat.Value = "mp3"
	ydl.Options.AudioQuality.Value = "192"
	ydl.Options.DownloadArchive.Value = archiveFile
	ydl.Options.EmbedThumbnail.Value = true
	ydl.Options.ExtractAudio.Value = true
	ydl.Options.Format.Value = "bestaudio/best"
	ydl.Options.KeepVideo.Value = false
	ydl.Options.IgnoreErrors.Value = true
	ydl.Options.NoMtime.Value = false
	ydl.Options.WriteThumbnail.Value = false
}

// Download downloads all Podcasts into their respective directories
func Download() {
	for key, value := range C.Podcasts {
		ydl.Options.Output.Value = fmt.Sprintf("%s/%s/%%(upload_date)s-%%(title)s.%%(ext)s", outputDir, key)
		downloadLink := fmt.Sprintf("%s%s", C.General["playlist_base"], value.PlaylistID)
		cmd, err := ydl.Download(downloadLink)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		cmd.Wait()
		for _, ID := range value.AdditionalEpisodes {
			downloadLink = fmt.Sprintf("%s%s", C.General["video_base"], ID)
			cmd, err = ydl.Download(downloadLink)
			if err != nil {
				log.Fatalf("%+v\n", err)
			}
			cmd.Wait()
		}
	}
}

func DownloadThumbnail(podcast string) {
	youtubeDL := goydl.NewYoutubeDl()
	youtubeDL.Options.IgnoreErrors.Value = true
	youtubeDL.Options.NoOverwrites.Value = false
	youtubeDL.Options.Output.Value = fmt.Sprintf("%s/%s/logo.%%(ext)s", outputDir, podcast)
	youtubeDL.Options.PlaylistItems.Value = "1"
	youtubeDL.Options.SkipDownload.Value = true
	youtubeDL.Options.WriteThumbnail.Value = true

	downloadLink := fmt.Sprintf("%s%s", C.General["playlist_base"], C.Podcasts[podcast].PlaylistID)
	cmd, err := youtubeDL.Download(downloadLink)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}
