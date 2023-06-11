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
import flag "github.com/spf13/pflag"


func main() {
	imgPath, imgWidth := getFlags()

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

func getFlags() (string, int){
	var imagePath string
	flag.StringVarP(&imagePath, "image","i" ,"", "The relative or absolute path of the image to be used. (REQUIRED)")
	flag.BoolP("help", "h", false, "Print help")
	var width int
	flag.IntVarP(&width, "width", "w", 50, "The image width")
	flag.Parse()
	if imagePath == "" {
		flag.PrintDefaults()
		log.Fatal("The image path is required (--image or -i)")
	}
	return imagePath, width
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


func consoleSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	heigth, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return heigth, width
}