package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

type rgbColor struct {
	red   uint8
	green uint8
	blue  uint8
}

func main() {

	colorArray := []rgbColor{}

	maxUint8 := 256

	for r := 0; r < maxUint8; r++ {
		for g := 0; g < maxUint8; g++ {
			for b := 0; b < maxUint8; b++ {
				red := uint8(r)
				green := uint8(g)
				blue := uint8(b)
				colorArray = append(colorArray, rgbColor{red, green, blue})
			}
		}
	}

	width := 4096
	height := 4096
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	max24Bit := 16777216

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			rand := rand.Intn(max24Bit)
			chosenColor := colorArray[rand]
			img.Set(h, w, color.RGBA{chosenColor.red, chosenColor.green, chosenColor.blue, 255})
		}
	}

	file, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	png.Encode(file, img)
}
