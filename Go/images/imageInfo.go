package main

import (
	"log"
	"os"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"reflect"
)

func main() {
	file, error := os.Open("flower.jpg")
	if error != nil { log.Fatal(error); }
	defer file.Close()

	image, _, error := image.DecodeConfig(file)
	if error != nil { log.Fatal(error);}

	log.Printf("Width: %d - Height: %d", image.Width, image.Height)

	imageColor := image.ColorModel
	log.Println(reflect.TypeOf(imageColor))
}
