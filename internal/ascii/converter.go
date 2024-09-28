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
	"sync"

	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/webp"
)

type convert struct {
	Color bool
	Wg    sync.WaitGroup
	Mx    sync.Mutex
}

const (
	asciiChars = "`.',-~:;=+*#%@M"                     // ASCII characters to use for the image rendering
	width      = 500                                   // Width of the ASCII art in characters
	fontWidth  = 7                                     // Width of each character in the font
	fontHeight = 13                                    // Height of each character in the font
	fontRatio = float64(fontWidth)/float64(fontHeight) // Font aspect ratio
)

// ConvertImage converts an image byte slice to ASCII art
func ConvertImage(imageBytes []byte, preserveColor bool) (string, error) {
	img, err := decodeImage(imageBytes)
	if err != nil {
		return "", err
	}

	c := convert{preserveColor, sync.WaitGroup{}, sync.Mutex{}}
	resizedImg := resizeImage(img)
	asciiArt := c.ToAscii(resizedImg)

	return asciiArt, nil
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
	resizedHeight := uint(float64(width) * ratio * fontRatio)
	resizedWidth := uint(width)
	return resize.Resize(resizedWidth, resizedHeight, img, resize.Lanczos3)
}

// ToAscii converts an image to ASCII characters
func (conv *convert) ToAscii(img image.Image) string {
	bounds := img.Bounds()
	out := make([]string, bounds.Dy())

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		conv.Wg.Add(1)
		go func(y int) {
			defer conv.Wg.Done()

			var str string
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				c := img.At(x, y)
				grayColor := color.GrayModel.Convert(c).(color.Gray)
				asciiIndex := int(grayColor.Y) * (len(asciiChars) - 1) / 255
				char := string(asciiChars[asciiIndex])

				if conv.Color {
					r, g, b, _ := c.RGBA()
					hexColor := fmt.Sprintf("#%02X%02X%02X", r>>8, g>>8, b>>8)
					char = fmt.Sprintf("<x-char style=\"color:%s\">%s</x-char>", hexColor, char)
					// char = fmt.Sprintf("%s%s<", hexColor, char)
				} else {
					hexColor := "#FFFFFF"
					char = fmt.Sprintf("<x-char style=\"color:%s\">%s</x-char>", hexColor, char)
				}
				str += char
			}
			// conv.Mx.Lock()
			out[y] = str
			// conv.Mx.Unlock()
		}(y)
	}
	conv.Wg.Wait()

	return strings.Join(out, "\n")
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
		subline := strings.Split(line, "<x-char")
		t := 0
		for _, tagged := range subline[1:] {
			tagged = "<x-char" + tagged
			textColor := parseColor(tagged)

			symbol := string(tagged[len(tagged)-9])
			drawChar(img, textColor, t*fontWidth, (y+1)*fontHeight, symbol)
			t += 1
		}
		// for x, ch := range line {
		// 	var textColor color.RGBA
		// 	if preserveColor {
		// 		textColor = parseColor(x, y, lines)
		// 	} else {
		// 		textColor = color.RGBA{255, 255, 255, 255} // White text for non-colored output
		// 	}
		// 	drawChar(img, textColor, x*fontWidth, (y+1)*fontHeight, string(ch))
		// }
	}

	// Save the image as PNG
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}

func parseColor(line string) color.RGBA {
	fmt.Println(line)
	// Parse the hex color
	clr := line[14:22]
	fmt.Println(clr)
	colorValue, err := strconv.ParseUint(clr, 16, 32)
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
