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
	"strconv"
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

	FormatJPEG = "jpeg"
	FormatPNG = "png"
	FormatWEBP = "webp"
)

// Map image formats to detect functions
var detectFuncMap = map[string]func([]byte) bool{
	FormatJPEG: isJPEG,
	FormatPNG:  isPNG,
	FormatWEBP: isWEBP,
}

// Detects if the byte slice is a JPEG image
func isJPEG(imageBytes []byte) bool {
	return bytes.HasPrefix(imageBytes, []byte("\xFF\xD8\xFF"))
}

// Detects if the byte slice is a PNG image
func isPNG(imageBytes []byte) bool {
	return bytes.HasPrefix(imageBytes, []byte("\x89PNG\r\n\x1a\n"))
}

// Detects if the byte slice is a WEBP image
func isWEBP(imageBytes []byte) bool {
	return bytes.HasPrefix(imageBytes, []byte("RIFF")) && bytes.Contains(imageBytes[8:16], []byte("WEBP"))
}

// ConvertImage converts an image byte slice to ASCII art
func ConvertImage(imageBytes []byte, preserveColor bool) (string, error) {
	img, err := decodeImage(imageBytes)
	if err != nil {
		return "", err
	}

	resizedImg := resizeImage(img)
	asciiArt := convertToAscii(resizedImg, preserveColor)

	return asciiArt, nil
}

// Map image formats to decode functions
var decodeFuncMap = map[string]func(*bytes.Reader) (image.Image, error){
	FormatJPEG: func(r *bytes.Reader) (image.Image, error) { return jpeg.Decode(r) },
	FormatPNG:  func(r *bytes.Reader) (image.Image, error) { return png.Decode(r) },
	FormatWEBP: func(r *bytes.Reader) (image.Image, error) { return webp.Decode(r) },
}

// decodeImage detects the image format and decodes it
func decodeImage(imageBytes []byte) (image.Image, error) {
	format, err := detectImageFormat(imageBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to detect image format: %w", err)
	}

	decodeFunc, found := decodeFuncMap[format]
	if !found {
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	reader := bytes.NewReader(imageBytes)

	img, err := decodeFunc(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}

// detectImageFormat determines the format of the image based on its byte signature
func detectImageFormat(imageBytes []byte) (string, error) {
	for format, detectFunc := range detectFuncMap {
		if detectFunc(imageBytes){
			return format, nil
		}
	}
	return "", fmt.Errorf("unsupported image format")
}

// resizeImage resizes the input image while maintaining aspect ratio
func resizeImage(img image.Image) image.Image {
	ratio := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	height := int(float64(width) * ratio * 0.5)
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}

// convertToAscii converts an image to ASCII characters
func convertToAscii(img image.Image, preserveColor bool) string {
	bounds := img.Bounds()
	var sb strings.Builder

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			grayColor := color.GrayModel.Convert(c).(color.Gray)
			asciiIndex := int(grayColor.Y) * (len(asciiChars) - 1) / 255
			char := string(asciiChars[asciiIndex])

			if preserveColor {
				r, g, b, _ := c.RGBA()
				hexColor := fmt.Sprintf("#%02X%02X%02X", r>>8, g>>8, b>>8)
				sb.WriteString(fmt.Sprintf("<x-char style=\"color:%s\">%s</x-char>", hexColor, char))
			} else {
				sb.WriteString(char)
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

// Convert ASCII art to an image
func ConvertAsciiToImage(asciiArt string, outputPath string, preserveColor bool) error {
	lines := strings.Split(strings.TrimSpace(asciiArt), "\n")

	// Calculate the dimensions of the image
	maxLineLength := 0
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}

	imgWidth := fontWidth * maxLineLength
	imgHeight := fontHeight * len(lines)

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Fill the image with a black background
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

	// Draw the ASCII text onto the image
	for y, line := range lines {
		for x, ch := range line {
			var textColor color.RGBA
			if preserveColor {
				textColor = parseColor(x, y, lines)
			} else {
				textColor = color.RGBA{255, 255, 255, 255} // White text for non-colored output
			}
			drawChar(img, textColor, x*fontWidth, (y+1)*fontHeight, string(ch))
		}
	}

	// Save the image as PNG
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}

func parseColor(x, y int, lines []string) color.RGBA {
	// Extract color from the x-char tag
	line := lines[y]
	startTag := strings.LastIndex(line[:x], "<x-char style=\"color:")
	if startTag == -1 {
		return color.RGBA{255, 255, 255, 255} // Default to white if no color found
	}
	endTag := strings.Index(line[startTag:], "\">")
	if endTag == -1 {
		return color.RGBA{255, 255, 255, 255} // Default to white if no color found
	}
	colorStr := line[startTag+len("<x-char style=\"color:") : startTag+endTag]

	// Parse the hex color
	colorValue, err := strconv.ParseUint(colorStr[1:], 16, 32)
	if err != nil {
		return color.RGBA{255, 255, 255, 255} // Default to white if parsing fails
	}

	return color.RGBA{
		R: uint8((colorValue >> 16) & 0xFF),
		G: uint8((colorValue >> 8) & 0xFF),
		B: uint8(colorValue & 0xFF),
		A: 255,
	}
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
