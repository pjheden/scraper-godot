package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

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

func main() {
	// godotengine collector
	godotCollector := colly.NewCollector(
		colly.AllowedDomains("godotengine.org"),
		colly.CacheDir("./godot_cache"),
	)

	githubCollector := godotCollector.Clone()
	githubCollector.AllowedDomains = []string{"github.com"}

	var currentAsset Asset
	assets := []Asset{}

	godotCollector.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	godotCollector.OnHTML("a[href].asset-header", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		e.Request.Visit(link)
	})

	godotCollector.OnHTML("body", func(e *colly.HTMLElement) {
		if e.Request.URL == nil || !strings.HasPrefix(e.Request.URL.String(), "https://godotengine.org/asset-library/asset/") {
			return
		}

		currentAsset = Asset{}

		{ // Get the ID from the URL.
			url := e.Request.URL.String()
			urlParts := strings.Split(url, "/")

			idInt, err := strconv.Atoi(urlParts[len(urlParts)-1])
			if err != nil {
				log.Fatalf("failed to convert url id to int: %s", err)
			}

			currentAsset.ID = idInt
		}

		// TODO: read other fields.

		e.ForEach("a[href].btn.btn-default", func(_ int, el *colly.HTMLElement) {
			link := el.Attr("href")

			if !strings.HasPrefix(link, "https://github.com/") {
				return
			}

			currentAsset.RepositoryURL = link

			githubCollector.Visit(link)
		})

		assets = append(assets, currentAsset)
	})

	githubCollector.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	githubCollector.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("span#repo-stars-counter-star", func(_ int, el *colly.HTMLElement) {
			stars, err := strconv.Atoi(el.Text)
			if err != nil {
				log.Fatalf("failed to convert stars to int: %s", err)
			}

			currentAsset.Stars = stars
		})
	})

	// Start scraping
	if err := godotCollector.Visit("https://godotengine.org/asset-library/asset"); err != nil {
		panic(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(assets)

}
