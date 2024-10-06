package handlers

import (
	"ascii-converter/backend/internal/ascii"
	"log"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const maxMemory = 10 << 20 // 10 MB

// Presents the page and has a form to upload an image
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	indexPath := filepath.Join(workDir, "build", "frontend", "index.html")
	http.ServeFile(w, r, indexPath)
}

// Handles the form submission to display ASCII art
func ConvertToImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Content-Type:", r.Header.Get("Content-Type"))

	// Parse the form
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		http.Error(w, "Unable to parse the form", http.StatusInternalServerError)
		fmt.Println("Error parsing form:", err)
		return
	}

	// Get the image from the form
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get the image", http.StatusInternalServerError)
		fmt.Println("Error getting image:", err)
		return
	}
	defer file.Close()

	// Read the image bytes
	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read the image", http.StatusInternalServerError)
		fmt.Println("Error reading image:", err)
		return
	}

	// If the preserve_color form value is true, preserve the original colors
	preserveColor := r.FormValue("preserve_color") == "on"

	// Convert the image to ASCII art
	ascii, err := ascii.ConvertImage(imageBytes, preserveColor)
	if err != nil {
		http.Error(w, "Unable to convert the image", http.StatusInternalServerError)
		fmt.Println("Error converting image:", err)
		return
	}

	// Send the ASCII art as a JSON response
	w.Header().Set("Content-Type", "text/html")
	json.NewEncoder(w).Encode(ascii)
}

// Handles ASCII to image conversion and serves the download
func ConvertAsciiToImageHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the input JSON
	var input struct {
		AsciiArt      string `json:"ascii-art"`
		PreserveColor bool   `json:"preserve_color"`
	}

	// Parse the JSON
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Decode the ASCII art from base64
	asciiArtBytes, err := base64.StdEncoding.DecodeString(input.AsciiArt)
	if err != nil {
		http.Error(w, "Invalid ASCII art input", http.StatusBadRequest)
		return
	}

	// Convert the ASCII art to a string
	asciiArt := string(asciiArtBytes)

	// Define the output path for the generated image file
	outputFile := filepath.Join(os.TempDir(), "ascii_image.png")

	// Convert the ASCII art to an image and save it
	err = ascii.ConvertAsciiToImage(asciiArt, outputFile, input.PreserveColor)
	if err != nil {
		http.Error(w, "Error generating image", http.StatusInternalServerError)
		return
	}

	// Serve the generated image as a download response
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_image.png")
	w.Header().Set("Content-Type", "image/png")
	file, err := os.Open(outputFile)
	if err != nil {
		http.Error(w, "Error serving image", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Write the image file to the response body
	io.Copy(w, file)
}
