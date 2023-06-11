package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/gookit/color"
	"github.com/nfnt/resize"
)

func main() {
	imageFile, err := os.Open("test_image.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()

	imgData, _, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}

	newWidth := 50
	newHeight := 25

	resizedImg := resizeImg(newWidth, newHeight, imgData)

	rgba := convertToRGBA(resizedImg)

	for r := 0; r < newHeight; r++ {
		for c := 0; c < newWidth; c++ {
			pixColor := rgba.At(c, r)
			r, g, b, _ := pixColor.RGBA()
			pixRgbColorFormatted := color.RGB(uint8(r>>8), uint8(g>>8), uint8(b>>8), true)
			pixRgbColorFormatted.Print(" ")
		}
		fmt.Println()
	}
}

func resizeImg(newWidth, newHeight int, imgData image.Image) image.Image {
	resizedImg := resize.Resize(uint(newWidth), uint(newHeight), imgData, resize.Lanczos3)
	return resizedImg
}

func convertToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba
}
