package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pjheden/scraper-godot/views"
	"github.com/pjheden/scraper-godot/views/components"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		assetsTable := components.Table{
			Columns: []components.TableColumn{
				{Name: "Title", IsHref: false},
				{Name: "Creator", IsHref: false},
				{Name: "Version", IsHref: false},
				{Name: "Stars", IsHref: false},
				{Name: "First Commit", IsHref: false},
				{Name: "Latest Commit", IsHref: false},
				{Name: "Repository URL", IsHref: true, ConvertValueFunc: func(_ string) string { return "" }},
			},
			Rows: [][]string{
				{"Godot", "Godot Engine", "3.3.3", "2,000", "2021-01-01", "2021-01-01", "https://github.com/pjheden"},
				{"Godot", "Godot Engine", "3.3.3", "2,000", "2021-01-01", "2021-01-01", "https://github.com/pjheden"},
				{"Godot", "Godot Engine", "3.3.3", "2,000", "2021-01-01", "2021-01-01", "https://github.com/pjheden"},
			},
		}

		templ.Handler(views.Index(assetsTable)).ServeHTTP(w, r)
	})

	fmt.Println("Listening to 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
