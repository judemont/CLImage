package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
	"log"
	"os"
	"github.com/gookit/color"
	"github.com/nfnt/resize"
	"github.com/alecthomas/kingpin/v2"
)

const VERSION = "v1.1.0"

func main() {
	imgPath, imageWidth := getArgs()

	imageFile, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()

	imgData, _, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}

	newWidth := imageWidth
	newHeight := newWidth/2

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

func getArgs() (string, int){
	var(
		imagePath = kingpin.Arg("image path", "The path of the image.").Required().String()
		imgWidth = kingpin.Flag("width", "Image width (in characters)").Short('w').Default("50").Int()
	)
	kingpin.Version(VERSION)
	kingpin.Parse()
	return *imagePath, *imgWidth
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
