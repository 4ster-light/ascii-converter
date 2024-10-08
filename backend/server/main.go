package main

import (
	"ascii-converter/backend/internal/handlers"
	"path/filepath"

	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.AllowAll().Handler)

	// Routes
	router.Get("/", handlers.HomeHandler)
	router.Post("/convert-to-ascii", handlers.ConvertToImageHandler)
	router.Post("/convert-to-image", handlers.ConvertAsciiToImageHandler)

	// Static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "build/frontend"))
	fileServer(router, "/", filesDir)

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

// Serves static files from the given root path
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	// Serve static files
	fs := http.StripPrefix(path, http.FileServer(root))
	r.Get(path+"*", func(w http.ResponseWriter, r *http.Request) {
		// Check if the file exists in the static directory
		if _, err := root.Open(strings.TrimPrefix(r.URL.Path, path)); err != nil {
			// If not found, serve index.html
			workDir, _ := os.Getwd() // Get current working directory
			http.ServeFile(w, r, filepath.Join(workDir, "build/frontend", "index.html"))
			return
		}
		// If file is found, serve it
		fs.ServeHTTP(w, r)
	})
}
