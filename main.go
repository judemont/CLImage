package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
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

	rgba := convertToRGBA(imgData)
	fmt.Println(rgba.At(0, 0)) // Example: accessing the color at pixel (0, 0)
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
