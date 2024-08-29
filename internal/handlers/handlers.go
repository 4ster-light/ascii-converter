package handlers

import (
	"ascii-converter/internal/ascii"
	"ascii-converter/internal/templates"

	"image/png"
	"io"
	"net/http"
)

// Presents the page and has a form to upload an image
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home("🏠 Home").Render(r.Context(), w)
}

// Handles the form submission
func ConvertImageHandler(w http.ResponseWriter, r *http.Request) {
	maxMemory := int64(10 << 20) // 10 MB

	// Parse the form
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		http.Error(w, "Unable to parse the form", http.StatusInternalServerError)
		return
	}

	// Get the image
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get the image", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the image
	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read the image", http.StatusInternalServerError)
		return
	}

	// Get the options
	colorOption := r.FormValue("color_option")
	outputType := r.FormValue("output_type")

	// Convert the image and get the ASCII text
	asciiResult, err := ascii.ConvertImage(imageBytes, colorOption, outputType)
	if err != nil {
		http.Error(w, "Unable to convert the image", http.StatusInternalServerError)
		return
	}

	// Return the ASCII text or the image
	if outputType == "text" {
		w.Header().Set("Content-Type", "text/plain")
		if asciiText, ok := asciiResult.(string); ok {
			w.Write([]byte(asciiText))
		} else {
			http.Error(w, "Unexpected result type for text output", http.StatusInternalServerError)
		}
	} else {
		// Convert the ASCII to an image
		if asciiText, ok := asciiResult.(string); ok {
			asciiImage, err := ascii.ConvertAsciiToImage(asciiText)
			if err != nil {
				http.Error(w, "Unable to convert the ASCII to an image", http.StatusInternalServerError)
				return
			}

			// Write the image
			w.Header().Set("Content-Type", "image/png")
			png.Encode(w, asciiImage)
		} else {
			http.Error(w, "Unexpected result type for image output", http.StatusInternalServerError)
		}
	}
}
