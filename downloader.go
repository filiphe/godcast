package main

import (
	"github.com/brianallred/goydl"
)

var ydl = goydl.NewYoutubeDl()

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
