package handlers

import (
	"ascii-converter/internal/ascii"
	"ascii-converter/internal/templates"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const maxMemory = 10 << 20 // 10 MB

// Presents the page and has a form to upload an image
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home("✰ArtSCII✰").Render(r.Context(), w)
}

// Handles the form submission to display ASCII art and the download button
func ConvertImageHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		http.Error(w, "Unable to parse the form", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get the image", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read the image", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	preserveColor := r.FormValue("preserve_color") == "on"

	coloredAscii, err := ascii.ConvertImage(imageBytes, true)
	if err != nil {
		http.Error(w, "Unable to convert the image", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	plainAscii, err := ascii.ConvertImage(imageBytes, false)
	if err != nil {
		http.Error(w, "Unable to convert the image", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Remove HTML tags from plainAscii
	plainAscii = strings.ReplaceAll(plainAscii, "<x-char style=\"color:#000000\">", "")
	plainAscii = strings.ReplaceAll(plainAscii, "</x-char>", "")

	displayAscii := plainAscii
	if preserveColor {
		displayAscii = coloredAscii
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<pre>%s</pre>
		<form action="/convert-to-image" method="POST">
			<input type="hidden" name="ascii-art" value="%s" />
			<input type="hidden" name="preserve_color" value="%t" />
			<button type="submit">Download as Image</button>
		</form>
	`, displayAscii, base64.StdEncoding.EncodeToString([]byte(plainAscii)), preserveColor)
}

// Handles ASCII to image conversion and serves the download
func ConvertAsciiToImageHandler(w http.ResponseWriter, r *http.Request) {
	asciiArtBase64 := r.FormValue("ascii-art")
	if asciiArtBase64 == "" {
		http.Error(w, "Missing ASCII art input", http.StatusBadRequest)
		return
	}

	preserveColor := r.FormValue("preserve_color") == "on"

	asciiArtBytes, err := base64.StdEncoding.DecodeString(asciiArtBase64)
	if err != nil {
		http.Error(w, "Invalid ASCII art input", http.StatusBadRequest)
		return
	}

	asciiArt := string(asciiArtBytes)

	// Define the output path for the generated image
	outputFile := filepath.Join(os.TempDir(), "ascii_image.png")

	// Convert ASCII to image and save it
	err = ascii.ConvertAsciiToImage(asciiArt, outputFile, preserveColor)
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
