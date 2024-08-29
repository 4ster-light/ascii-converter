package handlers

import (
	"ascii-converter/internal/ascii"
	"ascii-converter/internal/templates"
	"encoding/base64"
	"fmt"
	"html"

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
		fmt.Fprintf(w, "<pre>%s</pre>", html.EscapeString(asciiResult.(string)))
	} else {
		asciiImage, err := ascii.ConvertAsciiToImage(asciiResult.(string))
		if err != nil {
			http.Error(w, "Unable to convert ASCII to image", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "<img src='data:image/png;base64,")
		pngEncoder := base64.NewEncoder(base64.StdEncoding, w)
		png.Encode(pngEncoder, asciiImage)
		pngEncoder.Close()
		fmt.Fprintf(w, "' alt='ASCII Art'>")
	}
}
