package main

import (
	"log"
	"os"
	"image"
	"image/jpeg"
	"image/color"
	"reflect"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Usage: go run <program> <input_image> <output_image>")
	}
	
	// read file
	file, error := os.Open(os.Args[1])
	if error != nil { log.Fatal(error); }
	defer file.Close()

	// decode image
	imageSrc, error := jpeg.Decode(file)
	if error != nil { log.Fatal(error);}

	imageRect := imageSrc.Bounds()

	w, h := imageRect.Max.X, imageRect.Max.Y
	log.Printf("W: %d, H: %d", w, h)

	test := imageSrc.ColorModel()
	log.Println(reflect.TypeOf(test))
	
	// change color
	var imageRGBA *image.RGBA = image.NewRGBA(imageRect)
	var RGBAColor color.RGBA
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := imageSrc.At(x, y)
			RGBAColor = imageRGBA.ColorModel().Convert(oldColor).(color.RGBA)
			RGBAColor.G = 0
			RGBAColor.B = 0
			imageRGBA.SetRGBA(x, y, RGBAColor)
		}
	}
		
	// set write out and encode
	outPath, error := os.Create(os.Args[2])
	if error != nil { log.Fatal(error);}
	jpeg.Encode(outPath, imageRGBA, nil)
	
}
