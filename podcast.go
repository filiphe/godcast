package main

// Podcast holds information about a single podcast.
type Podcast struct {
	PlaylistID string `yaml:"playlist_id"`
	Name       string `yaml:"name"`
}
