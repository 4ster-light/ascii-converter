package main

import (
	"ascii-converter/internal/handlers"

	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Routes
	router.Get("/", handlers.HomeHandler)
	router.Post("/convert-to-ascii", handlers.ConvertImageHandler)
	router.Post("/convert-to-image", handlers.ConvertAsciiToImageHandler)

	// Get the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Starting server on: http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
