package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
		log.Printf("Downloading %s\n", value.Name)
		ydl.Options.Output.Value = fmt.Sprintf("%s/%s/%%(upload_date)s-%%(title)s.%%(ext)s", outputDir, key)
		downloadLink := fmt.Sprintf("%s%s", C.General["playlist_base"], value.PlaylistID)
		log.Println("Executing download command")
		cmd, err := ydl.Download(downloadLink)
		go io.Copy(os.Stdout, ydl.Stdout)
		go io.Copy(os.Stderr, ydl.Stderr)
		defer cmd.Wait()
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		//cmd.Wait()
		log.Println("Download command completed")
		for _, ID := range value.AdditionalEpisodes {
			downloadLink = fmt.Sprintf("%s%s", C.General["video_base"], ID)
			cmd, err = ydl.Download(downloadLink)
			defer cmd.Wait()
			if err != nil {
				log.Fatalf("%+v\n", err)
			}
			//cmd.Wait()
		}
	}
}

func DownloadThumbnail(podcast string) {
	log.Println("Downloading thumbnail")
	logoUrl := C.Podcasts[filepath.Base(podcast)].Logo
	if logoUrl != "" {
		fp, err := os.Create(fmt.Sprintf("%s/%s/logo.jpg", outputDir, podcast))
		if err != nil {
			log.Printf("Logo file create error: %+v\n", err)
		}
		resp, err := http.Get(logoUrl)
		if err != nil {
			log.Printf("Logo file GET error: %+v\n", err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Read logo GET body error: %+v\n", err)
		}
		fp.Write(data)
		fp.Close()

	} else {
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
			log.Fatalf("%+v\n", err)
		}
		cmd.Wait()
	}
	log.Println("Downloaded thumbnail")
}
