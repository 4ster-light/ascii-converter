package ascii

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/nfnt/resize"
)

// ASCII character set for grayscale mapping (from darkest to lightest)
var asciiChars = []byte("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")

// ConvertImage converts an image to ASCII art
func ConvertImage(imageBytes []byte, colorOption string) (string, error) {
	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return "", err
	}

	// Resize the image to a width of 80 pixels, maintaining aspect ratio
	img = resize.Resize(80, 0, img, resize.Lanczos3)

	// Generate and return ASCII art based on the color option
	switch colorOption {
	case "bw":
		return imageToAscii(convertToGrayscale(img)), nil
	case "color":
		return imageToColorAscii(img), nil
	default:
		return imageToAscii(img), nil // Default to grayscale ASCII
	}
}

// convertToGrayscale converts a color image to grayscale
func convertToGrayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayImg.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return grayImg
}

// imageToAscii converts an image to grayscale ASCII art
func imageToAscii(img image.Image) string {
	bounds := img.Bounds()
	var sb strings.Builder
	sb.Grow((bounds.Dx() + 1) * bounds.Dy()) // Pre-allocate space for efficiency, including newlines

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			sb.WriteByte(getAsciiChar(c.Y))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// imageToColorAscii converts an image to color ASCII art
func imageToColorAscii(img image.Image) string {
	bounds := img.Bounds()
	var sb strings.Builder
	sb.Grow((bounds.Dx() + 1) * bounds.Dy() * 20) // Pre-allocate space, including color codes

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			char := getAsciiChar(uint8((r + g + b) / 3 >> 8))
			sb.WriteString(getColoredChar(char, uint8(r>>8), uint8(g>>8), uint8(b>>8)))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// getAsciiChar maps a grayscale value to an ASCII character
func getAsciiChar(grayScale uint8) byte {
	index := int(grayScale) * (len(asciiChars) - 1) / 255
	return asciiChars[index]
}

// getColoredChar returns an ASCII character with ANSI color codes
func getColoredChar(char byte, r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c\x1b[0m", r, g, b, char)
}
