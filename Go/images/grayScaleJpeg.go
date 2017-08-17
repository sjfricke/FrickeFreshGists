package main

import (
	"log"
	"os"
	"image"
	"image/jpeg"
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
	imageSrc, _, error := image.Decode(file)
	if error != nil { log.Fatal(error);}

	bounds := imageSrc.Bounds()

	w, h := bounds.Max.X, bounds.Max.Y
	log.Printf("W: %d, H: %d", w, h)

	// Grey scale
	gray := image.NewGray(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := imageSrc.At(x, y)
//			grayColor := gray.ColorModel().Convert(oldColor)
			gray.Set(x, y, oldColor)
		}
	}
		
	// set write out and encode
	outPath, error := os.Create(os.Args[2])
	if error != nil { log.Fatal(error);}
	jpeg.Encode(outPath, gray, nil)
	
}
