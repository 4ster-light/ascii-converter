package ascii

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"strings"

	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var (
	errInvalidOptions        = errors.New("invalid options")
	errUnsupportedOutputType = errors.New("unsupported output type")
)

func ConvertImage(imageBytes []byte, formColorOption string, formOutputType string) (interface{}, error) {
	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	// Resize the image to a smaller size for ASCII conversion
	img = resize.Resize(80, 0, img, resize.Lanczos3)

	// Convert to grayscale if black and white option is selected
	if formColorOption == "bw" {
		img = convertToGrayscale(img)
	}

	switch formOutputType {
	case "text":
		return imageToAscii(img), nil
	case "image":
		return ConvertAsciiToImage(imageToAscii(img))
	default:
		return nil, errUnsupportedOutputType
	}
}

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

func imageToAscii(img image.Image) string {
	bounds := img.Bounds()
	ascii := ""
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			ascii += getAsciiChar(c.Y)
		}
		ascii += "\n"
	}
	return ascii
}

func getAsciiChar(grayScale uint8) string {
	chars := []string{"@", "#", "S", "%", "?", "*", "+", ";", ":", ",", "."}
	index := int(grayScale) * len(chars) / 256
	return chars[index]
}

func ConvertAsciiToImage(ascii string) (image.Image, error) {
	lines := strings.Split(ascii, "\n")
	width := len(lines[0])
	height := len(lines)

	img := image.NewRGBA(image.Rect(0, 0, width*7, height*15))
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
	}

	for y, line := range lines {
		for x, char := range line {
			d.Dot = fixed.Point26_6{X: fixed.Int26_6(x * 7 * 64), Y: fixed.Int26_6((y + 1) * 15 * 64)}
			d.DrawString(string(char))
		}
	}

	return img, nil
}
