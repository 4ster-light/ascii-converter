package handlers

import (
	"ascii-converter/internal/ascii"
	"ascii-converter/internal/templates"

	"fmt"
	"html"
	"io"
	"net/http"
)

// Presents the page and has a form to upload an image
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home("🏠 Home").Render(r.Context(), w)
}

// Handles the form submission
func ConvertImageHandler(w http.ResponseWriter, r *http.Request) {
	const maxMemory = 10 << 20 // 10 MB

	// Parse the form data
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		http.Error(w, "Unable to parse the form", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Get the uploaded file
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get the image", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file data
	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read the image", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Convert the image to ASCII
	asciiResult, err := ascii.ConvertImage(imageBytes)
	if err != nil {
		http.Error(w, "Unable to convert the image", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Set the content type and return the ASCII text
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<pre>%s</pre>", html.EscapeString(asciiResult))
}
