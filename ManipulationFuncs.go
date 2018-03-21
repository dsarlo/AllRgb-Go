package main

import (
	"image"
	"image/color"
	"math/rand"
)

func randomNonUniquePic(img *image.RGBA, colorArray []rgbColor) {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			rand := rand.Intn(max24Bit)
			chosenColor := colorArray[rand]
			img.Set(h, w, color.RGBA{chosenColor.red, chosenColor.green, chosenColor.blue, 255})
		}
	}
}

func randomUniquePic(img *image.RGBA, colorArray []rgbColor, partialShuffle bool) {
	colorArray = shuffleColorArray(colorArray, partialShuffle) //Shuffle the array first for uniqueness (half shuffle = true in parameter (cool effect))
	iterator := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			chosenColor := colorArray[iterator]
			img.Set(h, w, color.RGBA{chosenColor.red, chosenColor.green, chosenColor.blue, 255})
			iterator++
		}
	}
}

func uniqueFIFOPic(img *image.RGBA, colorArray []rgbColor) {
	iterator := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			chosenColor := colorArray[iterator]
			img.Set(h, w, color.RGBA{chosenColor.red, chosenColor.green, chosenColor.blue, 255})
			iterator++
		}
	}
}

//HELPER FUNCTIONS
func shuffleColorArray(colorArray []rgbColor, partialShuffle bool) []rgbColor {
	shuffleConstant := 2
	if partialShuffle {
		shuffleConstant = 1
	}
	for index := 0; index < len(colorArray)*shuffleConstant; index++ {
		firstRand := rand.Intn(max24Bit)
		secondRand := rand.Intn(max24Bit)
		for secondRand == firstRand {
			secondRand = rand.Intn(max24Bit)
		}
		a := colorArray[firstRand]
		b := colorArray[secondRand]
		//swap them
		colorArray[secondRand] = a
		colorArray[firstRand] = b
	}
	return colorArray
}
