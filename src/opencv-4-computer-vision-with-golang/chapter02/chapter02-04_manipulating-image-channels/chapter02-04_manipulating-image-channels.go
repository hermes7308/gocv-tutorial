package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	originalMat := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if originalMat.Empty() {
		log.Panic("Can not read image: " + imageFilePath)
		return
	}

	image := originalMat.ToBytes()
	width := originalMat.Rows()
	height := originalMat.Cols()
	channel := originalMat.Channels()

	// swap red and blue(BGR)
	for row := 0; row < width; row++ {
		for col := 0; col < height; col++ {
			index := (width*row + col) * channel
			image[index], image[index+2] = image[index+2], image[index]
		}
	}

	matForSwap, err := gocv.NewMatFromBytes(width, height, gocv.MatTypeCV8UC3, image)
	if err != nil {
		log.Panic("Can not swap red and blue", err)
		return
	}

	windowForSwapRedAndBlue := gocv.NewWindow("Swap red and blue")
	defer windowForSwapRedAndBlue.Close()
	windowForSwapRedAndBlue.IMShow(matForSwap)

	// back to original and scale
	for row := 0; row < width; row++ {
		for col := 0; col < height; col++ {
			index := (width*row + col) * channel
			image[index], image[index+2] = image[index+2], image[index]

			// blue scale
			scaleForBlue := float32(image[index]) * 0.9
			image[index] = uint8(scaleForBlue)

			// green scale
			scaleForGreen := float32(image[index+1]) * 1.1
			image[index] = uint8(scaleForGreen)
		}
	}

	matForScale, err := gocv.NewMatFromBytes(width, height, gocv.MatTypeCV8UC3, image)
	if err != nil {
		log.Panic("Can not scale", err)
		return
	}

	windowForScale := gocv.NewWindow("Scale Image")
	defer windowForScale.Close()
	windowForScale.IMShow(matForScale)

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}
