package main

import (
	"gocv.io/x/gocv"
	"log"
	"math"
)

const GAMMA float64 = 0.5

func main() {
	imageFilePath := "../../data/Lena.png"
	originalMat := gocv.IMRead(imageFilePath, gocv.IMReadGrayScale)
	if originalMat.Empty() {
		log.Panic("Can not read image: " + imageFilePath)
		return
	}

	windowForOriginal := gocv.NewWindow("Original Image")
	defer windowForOriginal.Close()
	windowForOriginal.IMShow(originalMat)

	// corrected Image
	correctedImage := originalMat.ToBytes()

	for i := 0; i < len(correctedImage); i++ {
		correctedValue := uint8(math.Pow(float64(correctedImage[i])/255, GAMMA) * 255)
		correctedImage[i] = correctedValue
	}

	correctedMat, err := gocv.NewMatFromBytes(originalMat.Rows(), originalMat.Cols(), gocv.IMReadGrayScale, correctedImage)
	if err != nil {
		log.Panic("Can not convert bytes to Mat")
		return
	}

	windowForCorrect := gocv.NewWindow("Corrected Image")
	defer windowForCorrect.Close()
	windowForCorrect.IMShow(correctedMat)

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}
