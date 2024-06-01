package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pjheden/scraper-godot/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Index()).ServeHTTP(w, r)
	})

	fmt.Println("Listening to 8080...")
	http.ListenAndServe(":8080", r)
}
