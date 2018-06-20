package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	//"time"
	"github.com/gorilla/feeds"
	"strings"
)

// Podcast holds information about a single podcast.
type Podcast struct {
	PlaylistID         string   `yaml:"playlist_id"`
	Name               string   `yaml:"name"`
	AdditionalEpisodes []string `yaml:"additional_episodes"`
	Logo               string   `yaml:"logo"`
}

// TODO: Compare podcast "updated time" to latest podcast episode "created time"
// and update podcast when applicable

// GenerateFeed generates a feed for a podcast
func GenerateFeed(podcastDir string) {
	log.Printf("Generating feed for %s\n", podcastDir)
	podcastKey := filepath.Base(podcastDir)
	feed := &feeds.Feed{
		Title: C.Podcasts[podcastKey].Name,
		Link:  &feeds.Link{Href: fmt.Sprintf("%s/%s", C.General["url_base"], podcastKey)},
		Image: &feeds.Image{
			Url: fmt.Sprintf("%s/%s/logo.jpg", C.General["url_base"], podcastKey),
		},
	}

	err := filepath.Walk(podcastDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("GenerateFeed: prevent panic by handling failure accessing a path %q: %v\n", podcastDir, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".mp3" {
			item := &feeds.Item{
				Title:   strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
				Link:    &feeds.Link{Href: fmt.Sprintf("%s/%s/%s", C.General["url_base"], podcastKey, info.Name())},
				Created: info.ModTime(),
				Enclosure: &feeds.Enclosure{
					Url:    fmt.Sprintf("%s/%s/%s", C.General["url_base"], podcastKey, info.Name()),
					Length: fmt.Sprintf("%d", info.Size()),
					Type:   "audio/mpeg",
				},
			}
			feed.Add(item)
			log.Printf("Adding %s to %s feed\n", info.Name(), podcastDir)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("GenerateFeeds: error walking the path %q: %v\n", podcastDir, err)
	}
	rss, err := feed.ToRss()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fp, err := os.Create(filepath.Join(outputDir, podcastKey, "feed.rss"))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Println("Writing feed to file")
	fp.WriteString(rss)
	fp.Close()
}
