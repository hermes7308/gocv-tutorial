package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	img := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if &img == nil {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	window := gocv.NewWindow("Original Image")
	window.IMShow(img)
	gocv.WaitKey(5000) // if you want to show continually, 5000 -> -1
}
