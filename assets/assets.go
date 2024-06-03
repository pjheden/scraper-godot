package assets

import "time"

// Asset stores information about a godot asset.
type Asset struct {
	ID int
	// Created       time.Time
	// Updated       time.Time
	Title         string
	Description   string
	Creator       string
	Version       string
	RepositoryURL string
	Stars         int
	FirstCommit   time.Time
	LatestCommit  time.Time
}
