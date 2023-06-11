package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"github.com/gookit/color"
	"github.com/nfnt/resize"
	"github.com/urfave/cli/v2"
)



func main() {
	imgPath, imgWidth := getArgs()

	imageFile, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()

	imgData, _, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}

	newWidth := imgWidth
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
	var imagePath string
	var imgWidth int
	app := &cli.App{
        Name:  "Image Displayer",
        Usage: "Display images in your terminal, with colored characters.",
		Action: func(cCtx *cli.Context) error {
            imagePath = cCtx.Args().Get(0)
            return nil
        },
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "width",
				Value: 50,
				Usage: "Image width",
				Destination: &imgWidth,
			},
		},
    }

	if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
	return imagePath, imgWidth
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
