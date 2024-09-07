package ascii

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/nfnt/resize"
	"golang.org/x/image/webp"
)

const asciiChars = "`.',-~:;=+*#%@M"

// Converts a PNG byte slice to ASCII art
func ConvertImage(imageBytes []byte) (string, error) {
	format, err := detectImageFormat(imageBytes)
	if err != nil {
		return "", fmt.Errorf("failed to detect image format: %w", err)
	}

	var img image.Image
	switch format {
	case "jpeg":
		img, err = jpeg.Decode(bytes.NewReader(imageBytes))
	case "png":
		img, err = png.Decode(bytes.NewReader(imageBytes))
	case "webp":
		img, err = webp.Decode(bytes.NewReader(imageBytes))
	default:
		return "", fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	width := 500
	ratio := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	height := int(float64(width) * ratio * 0.5)

	img = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	asciiArt := convertToAscii(img)

	asciiArt = addBorder(asciiArt)

	return asciiArt, nil
}

func detectImageFormat(imageBytes []byte) (string, error) {
	if bytes.HasPrefix(imageBytes, []byte("\xFF\xD8\xFF")) {
		return "jpeg", nil
	} else if bytes.HasPrefix(imageBytes, []byte("\x89PNG\r\n\x1a\n")) {
		return "png", nil
	} else if bytes.HasPrefix(imageBytes, []byte("RIFF")) && bytes.Contains(imageBytes[8:16], []byte("WEBP")) {
		return "webp", nil
	}

	return "", fmt.Errorf("unsupported image format")
}

func convertToAscii(img image.Image) string {
	bounds := img.Bounds()
	var sb strings.Builder

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

func addBorder(asciiArt string) string {
	lines := strings.Split(asciiArt, "\n")
	wifth := len(lines[0]) + 4

	top := strings.Repeat("-", wifth)
	bordered := []string{top}

	for _, line := range lines {
		if line != "" {
			bordered = append(bordered, fmt.Sprintf("| %s |", line))
		}
	}

	bordered = append(bordered, top)
	return strings.Join(bordered, "\n")
}
