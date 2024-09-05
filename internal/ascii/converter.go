package ascii

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"strings"

	"github.com/nfnt/resize"
)

const asciiChars = ".:;=+*#%@"

// Converts a PNG byte slice to ASCII art
func ConvertImage(imageBytes []byte) (string, error) {
	// Decode the PNG image
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return "", fmt.Errorf("unable to decode the image: %w", err)
	}

	img = resize.Resize(100, 0, img, resize.Lanczos3)

	asciiArt := convertToAscii(img)

	asciiArt = addBorder(asciiArt)

	return asciiArt, nil
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
