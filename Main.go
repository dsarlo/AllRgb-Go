package main

import (
	"image"
	"image/png"
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
var max24Bit = 16777216

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

	//SELECT WHICH FUNCTION TO RUN THAT WILL CREATE THE IMAGE
	//randomNonUniquePic(&img, colorArray)
	//randomUniquePic(&img, colorArray, false)
	uniqueFIFOPic(&img, colorArray)

	file, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	png.Encode(file, &img)
}
