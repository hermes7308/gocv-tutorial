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

	// original Image
	log.Println("size:", originalMat.Size())
	log.Println("type:", originalMat.Type()) // MatTypeCV8UC3
	log.Println("channel:", originalMat.Channels())

	windowForOriginal := gocv.NewWindow("Window Original")
	defer windowForOriginal.Close()
	windowForOriginal.IMShow(originalMat)

	// scale Image
	originalImage := originalMat.ToBytes()
	newImage := make([]uint8, len(originalImage))
	for i := 0; i < len(originalImage); i++ {
		clipValue := float32(originalImage[i]) / 255
		if clipValue*2 > 1 {
			newImage[i] = 1
		} else {
			newImage[i] = 0
		}
	}

	for i := 0; i < len(newImage); i++ {
		newImage[i] = newImage[i] * 255
	}

	scaleMat, err := gocv.NewMatFromBytes(originalMat.Rows(), originalMat.Cols(), gocv.MatTypeCV8UC3, newImage)
	if err != nil {
		log.Panic("Can not convert bytes to Mat")
		return
	}

	windowForScale := gocv.NewWindow("Window Scale")
	defer windowForScale.Close()
	windowForScale.IMShow(scaleMat)

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}
