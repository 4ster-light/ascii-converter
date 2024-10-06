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
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/webp"
)

const (
	asciiChars = "`.',-~:;=+*#%@M"
	width      = 500
	fontWidth  = 7
	fontHeight = 13
	fontRatio  = float64(fontWidth) / float64(fontHeight)
)

type convert struct {
	preserveColor bool
	wg            sync.WaitGroup
	// mu            sync.Mutex
}

// ConvertImage converts an image byte slice to ASCII art
func ConvertImage(imageBytes []byte, preserveColor bool) (string, error) {
	img, err := decodeImage(imageBytes)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	c := &convert{preserveColor: preserveColor}
	resizedImg := resizeImage(img)
	asciiArt := c.toAscii(resizedImg)

	return asciiArt, nil
}

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
		return nil, fmt.Errorf("failed to decode %s image: %w", format, err)
	}

	return img, nil
}

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

func resizeImage(img image.Image) image.Image {
	ratio := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	resizedHeight := uint(float64(width) * ratio * fontRatio)
	return resize.Resize(width, resizedHeight, img, resize.Lanczos3)
}

func (c *convert) toAscii(img image.Image) string {
	bounds := img.Bounds()
	out := make([]string, bounds.Dy())

	c.wg.Add(bounds.Dy())
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		go func(y int) {
			defer c.wg.Done()
			c.processRow(img, y, bounds, &out[y])
		}(y)
	}
	c.wg.Wait()

	return strings.Join(out, "\n")
}

func (c *convert) processRow(img image.Image, y int, bounds image.Rectangle, outRow *string) {
	var row strings.Builder
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		col := img.At(x, y)
		grayColor := color.GrayModel.Convert(col).(color.Gray)
		asciiIndex := int(grayColor.Y) * (len(asciiChars) - 1) / 255
		char := string(asciiChars[asciiIndex])

		if c.preserveColor {
			r, g, b, _ := col.RGBA()
			hexColor := fmt.Sprintf("#%02X%02X%02X", r>>8, g>>8, b>>8)
			char = fmt.Sprintf("<x-char style=\"color:%s\">%s</x-char>", hexColor, char)
		} else {
			char = fmt.Sprintf("<x-char>%s</x-char>", char)
		}
		row.WriteString(char)
	}
	*outRow = row.String()
}

// ConvertAsciiToImage converts ASCII art to an image
func ConvertAsciiToImage(asciiArt string, outputPath string, preserveColor bool) error {
	lines := strings.Split(strings.TrimSpace(asciiArt), "\n")
	imgWidth, imgHeight := calculateImageDimensions(lines)

	img := createImage(imgWidth, imgHeight)
	if err := drawAsciiOnImage(img, lines, preserveColor); err != nil {
		return fmt.Errorf("failed to draw ASCII on image: %w", err)
	}

	if err := saveImage(img, outputPath); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}

func calculateImageDimensions(lines []string) (int, int) {
	maxLineLength := 0
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}
	return fontWidth * maxLineLength, fontHeight * len(lines)
}

func createImage(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)
	return img
}

func drawAsciiOnImage(img *image.RGBA, lines []string, preserveColor bool) error {
	for y, line := range lines {
		sublines := strings.Split(line, "<x-char")
		for x, tagged := range sublines[1:] {
			tagged = "<x-char" + tagged
			var textColor color.RGBA
			var err error
			if preserveColor {
				textColor, err = parseColor(tagged)
				if err != nil {
					return fmt.Errorf("failed to parse color: %w", err)
				}
			} else {
				textColor = color.RGBA{255, 255, 255, 255} // White color if not preserving color
			}

			symbol := string(tagged[len(tagged)-9])
			drawChar(img, textColor, x*fontWidth, (y+1)*fontHeight, symbol)
		}
	}
	return nil
}

func parseColor(line string) (color.RGBA, error) {
	re := regexp.MustCompile(`color:(#[0-9A-Fa-f]{6})`)
	match := re.FindStringSubmatch(line)
	if len(match) < 2 {
		return color.RGBA{255, 255, 255, 255}, nil
	}

	colorValue, err := strconv.ParseUint(match[1][1:], 16, 32)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid color format: %w", err)
	}

	return color.RGBA{
		R: uint8((colorValue >> 16) & 0xFF),
		G: uint8((colorValue >> 8) & 0xFF),
		B: uint8(colorValue & 0xFF),
		A: 255,
	}, nil
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

func saveImage(img *image.RGBA, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	if err := png.Encode(outFile, img); err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}
