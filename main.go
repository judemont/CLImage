package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
	_ "image/jpg"
	"log"
	"os"
	"github.com/gookit/color"
	"github.com/nfnt/resize"
	"github.com/urfave/cli/v2"
)

const VERSION = "v1.0.22"

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
	var imagePath string
	var imgWidth int

    cli.VersionPrinter = func(cCtx *cli.Context) {
        fmt.Println(cCtx.App.Version)
    }

	app := &cli.App{
        Name:  "CLI Image Displayer",
        Usage: "Display images in your terminal, with colored characters.",
		Version: VERSION,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "width",
				Aliases: []string{"W", "w"},
				Value: 50,
				Usage: "Image width (in characters)",
				Destination: &imgWidth,
			},
		},
		Action: func(cCtx *cli.Context) error {
            imagePath = cCtx.Args().Get(0)
            return nil
        },
    }

	if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
	if imagePath == ""{
		os.Exit(0)
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
