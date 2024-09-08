package handlers

import (
	"ascii-converter/internal/ascii"
	"ascii-converter/internal/templates"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const maxMemory = 10 << 20 // 10 MB

// Presents the page and has a form to upload an image
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home("✰ArtSCII✰").Render(r.Context(), w)
}

// Handles the form submission to display ASCII art and the download button
func ConvertImageHandler(w http.ResponseWriter, r *http.Request) {
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

	// Return ASCII art with a download button
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<pre>%s</pre>
		<form action="/convert-to-image" method="POST">
			<input type="hidden" name="ascii-art" value="%s" />
			<button type="submit">Download as Image</button>
		</form>
	`, html.EscapeString(asciiResult), html.EscapeString(asciiResult))
}

// Handles ASCII to image conversion and serves the download
func ConvertAsciiToImageHandler(w http.ResponseWriter, r *http.Request) {
	asciiArt := r.FormValue("ascii-art")
	if asciiArt == "" {
		http.Error(w, "Missing ASCII art input", http.StatusBadRequest)
		return
	}

	// Define the output path for the generated image
	outputFile := filepath.Join(os.TempDir(), "ascii_image.png")

	// Convert ASCII to image and save it
	err := ascii.ConvertAsciiToImage(asciiArt, outputFile)
	if err != nil {
		http.Error(w, "Error generating image", http.StatusInternalServerError)
		return
	}

	// Serve the generated image as a download
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_image.png")
	w.Header().Set("Content-Type", "image/png")
	file, err := os.Open(outputFile)
	if err != nil {
		http.Error(w, "Error serving image", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Write the image file to the response
	io.Copy(w, file)
}
