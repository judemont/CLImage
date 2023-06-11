package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"

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

	// Redimensionner l'image à 500x500 pixels
	newWidth := 500
	newHeight := 500
	resizedImg := resize.Resize(uint(newWidth), uint(newHeight), imgData, resize.Lanczos3)

	rgba := convertToRGBA(resizedImg)
	fmt.Println(rgba.At(0, 100))
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
