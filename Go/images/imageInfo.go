package main

import (
	"fmt"
	"log"
	"os"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run imageInfo.go <input_image>")
	}
		
	file, error := os.Open(os.Args[1])
	if error != nil { log.Fatal(error); }
	defer file.Close()

	image, imageMime, error := image.Decode(file)
	if error != nil { log.Fatal(error);}

	imageColor := image.At(0,0)
	var imageColorType string
	
	switch imageColor.(type) {
	case color.Alpha:
		imageColorType = "Alpha"
	case color.Alpha16:
		imageColorType = "Alpha16"		
	case color.CMYK:
		imageColorType = "CMYK"
	case color.Gray:
		imageColorType = "Gray"
	case color.Gray16:
		imageColorType = "Gray16"
	case color.NRGBA:
		imageColorType = "NRGBA"
	case color.NRGBA64:
		imageColorType = "NRGBA64"
	case color.NYCbCrA:
		imageColorType = "NYCbCrA"
	case color.RGBA:
		imageColorType = "RGBA"
	case color.RGBA64:
		imageColorType = "RGBA64"
	case color.YCbCr:
		imageColorType = "YCbCr"
	default:
		imageColorType = "Non Supported Color"
	}
	
	fmt.Printf("%s:\n", os.Args[1])
	fmt.Printf("\tType: %s\n", imageMime)
	fmt.Printf("\tWidth: %dpx - Height: %dpx\n", image.Bounds().Max.X, image.Bounds().Max.Y)
	fmt.Printf("\tColor: %s\n", imageColorType) 
}
