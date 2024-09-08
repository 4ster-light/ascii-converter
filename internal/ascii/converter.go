package ascii

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/webp"
)

const (
	asciiChars = "`.',-~:;=+*#%@M" // ASCII characters to use for the image rendering
	width      = 500               // Width of the ASCII art in characters
	fontWidth  = 7                 // Width of each character in the font
	fontHeight = 13                // Height of each character in the font
)

// ConvertImage converts an image byte slice to ASCII art
func ConvertImage(imageBytes []byte) (string, error) {
	img, err := decodeImage(imageBytes)
	if err != nil {
		return "", err
	}

	resizedImg := resizeImage(img)
	asciiArt := convertToAscii(resizedImg)
	borderedAsciiArt := addBorder(asciiArt)

	return borderedAsciiArt, nil
}

// decodeImage detects the image format and decodes it
func decodeImage(imageBytes []byte) (image.Image, error) {
	format, err := detectImageFormat(imageBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to detect image format: %w", err)
	}

	reader := bytes.NewReader(imageBytes)
	var img image.Image

	switch format {
	case "jpeg":
		img, err = jpeg.Decode(reader)
	case "png":
		img, err = png.Decode(reader)
	case "webp":
		img, err = webp.Decode(reader)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}

// detectImageFormat determines the format of the image based on its byte signature
func detectImageFormat(imageBytes []byte) (string, error) {
	switch {
	case bytes.HasPrefix(imageBytes, []byte("\xFF\xD8\xFF")):
		return "jpeg", nil
	case bytes.HasPrefix(imageBytes, []byte("\x89PNG\r\n\x1a\n")):
		return "png", nil
	case bytes.HasPrefix(imageBytes, []byte("RIFF")) && bytes.Contains(imageBytes[8:16], []byte("WEBP")):
		return "webp", nil
	default:
		return "", fmt.Errorf("unsupported image format")
	}
}

// resizeImage resizes the input image while maintaining aspect ratio
func resizeImage(img image.Image) image.Image {
	ratio := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	height := int(float64(width) * ratio * 0.5)
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}

// convertToAscii converts an image to ASCII characters
func convertToAscii(img image.Image) string {
	bounds := img.Bounds()
	var sb strings.Builder

	// Iterate over each pixel in the image and convert it to grayscale and then to ASCII
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y))
			grayColor, _ := c.(color.Gray)
			asciiIndex := int(grayColor.Y) * (len(asciiChars) - 1) / 255
			sb.WriteByte(asciiChars[asciiIndex])
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

// addBorder adds a border around the ASCII art
func addBorder(asciiArt string) string {
	lines := strings.Split(asciiArt, "\n")
	width := len(lines[0]) + 4

	top := strings.Repeat("-", width)
	bordered := []string{top}

	for _, line := range lines {
		if line != "" {
			bordered = append(bordered, fmt.Sprintf("| %s |", line))
		}
	}

	bordered = append(bordered, top)
	return strings.Join(bordered, "\n")
}

// Convert ASCII art to an image
func ConvertAsciiToImage(asciiArt string, outputPath string) error {
	lines := strings.Split(asciiArt, "\n")

	// Calculate the dimensions of the image
	imgWidth := fontWidth * (len(lines[0]) - 1)
	imgHeight := fontHeight * len(lines)

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Fill the image with white background
	// Shouldn't it be better to not have a background since it's a PNG?
	// Right now same color as website
	white := color.RGBA{26, 26, 26, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// Draw the ASCII text onto the image
	black := color.RGBA{255, 255, 255, 255}

	for y, line := range lines {
		for x, ch := range line {
			drawChar(img, black, x*fontWidth, (y+1)*fontHeight, string(ch))
		}
	}

	// Save the image as PNG or JPEG
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if strings.HasSuffix(outputPath, ".png") {
		err = png.Encode(outFile, img)
	} else if strings.HasSuffix(outputPath, ".jpg") || strings.HasSuffix(outputPath, ".jpeg") {
		err = jpeg.Encode(outFile, img, nil)
	} else {
		return fmt.Errorf("unsupported file format")
	}

	return err
}

// Draws a character on the image at the specified position
func drawChar(img *image.RGBA, col color.Color, x, y int, s string) {
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(s)
}
