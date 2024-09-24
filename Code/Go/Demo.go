package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/png"
	"math"
	"math/big"
	"strings"
)

type ACaptcha struct{}

func (a *ACaptcha) GenerateCaptcha(width, height int) (string, error) {
	// Generate a random code
	code, err := a.generateRandomCode(5)
	if err != nil {
		return "", err
	}

	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw the code on the image
	drawString(img, code, 10, 10, color.Black)

	// Add some noise to the image
	for i := 0; i < 10; i++ {
		x1 := int(math.Floor(rand.Float64() * float64(width)))
		y1 := int(math.Floor(rand.Float64() * float64(height)))
		x2 := int(math.Floor(rand.Float64() * float64(width)))
		y2 := int(math.Floor(rand.Float64() * float64(height)))
		drawLine(img, x1, y1, x2, y2, color.Black)
	}

	// Encode the image to a base64 string
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func (a *ACaptcha) generateRandomCode(length int) (string, error) {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
	 code[i] = characters[n.Int64()]
	}
	return string(code), nil
}

func drawString(img *image.RGBA, text string, x, y int, color color.Color) {
	font := &image.Font{}
	draw.Draw(img, image.Rect(x, y, x+font.Width*len(text), y+font.Height), &image.Uniform{color}, image.ZP, draw.Src)
	for i, c := range text {
		draw.Draw(img, image.Rect(x+font.Width*i, y, x+font.Width*(i+1), y+font.Height), &image.Uniform{color}, image.ZP, draw.Src)
	}
}

func drawLine(img *image.RGBA, x1, y1, x2, y2 int, color color.Color) {
	draw.Draw(img, image.Rect(x1, y1, x2, y2), &image.Uniform{color}, image.ZP, draw.Src)
}
