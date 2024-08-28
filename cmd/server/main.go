package main

import (
	"ascii-converter/internal/templates"

	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home().Render(r.Context(), w)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", baseHandler)

	log.Println("Server running on: localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
