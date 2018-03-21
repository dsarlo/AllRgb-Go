package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

//Custom struct for holding the unique colors generated on start
type rgbColor struct {
	red   uint8
	green uint8
	blue  uint8
}

//Global height and width of image being created
//4096x4096 is for 24 bit uniqueness
var width = 4096
var height = 4096

func randomNonUniquePic(img image.RGBA, colorArray []rgbColor) {
	max24Bit := 16777216
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			rand := rand.Intn(max24Bit)
			chosenColor := colorArray[rand]
			img.Set(h, w, color.RGBA{chosenColor.red, chosenColor.green, chosenColor.blue, 255})
		}
	}
}

func uniqueFIFOPic(img image.RGBA, colorArray []rgbColor) {
	iterator := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			chosenColor := colorArray[iterator]
			img.Set(h, w, color.RGBA{chosenColor.red, chosenColor.green, chosenColor.blue, 255})
			iterator++
		}
	}
}

func main() { //Builds the unique color array (24 bit) and calls the method of your chosen design
	//TODO possibly add command line args?

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

	img := *image.NewRGBA(image.Rect(0, 0, width, height))

	//randomNonUniquePic(img, colorArray)
	uniqueFIFOPic(img, colorArray)

	file, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	png.Encode(file, &img)
}
