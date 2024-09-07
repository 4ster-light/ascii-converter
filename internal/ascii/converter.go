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
// Convert ASCII art to an image

func ConvertAsciiToImage(asciiArt string, outputPath string) error {
	lines := strings.Split(asciiArt, "\n")

	// Calculate the dimensions of the image
	fontWidth := 7
	fontHeight := 13
	imgWidth := fontWidth * len(lines[0])
	imgHeight := fontHeight * len(lines)

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Fill the image with white background
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// Draw the ASCII text onto the image
	black := color.RGBA{0, 0, 0, 255}

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