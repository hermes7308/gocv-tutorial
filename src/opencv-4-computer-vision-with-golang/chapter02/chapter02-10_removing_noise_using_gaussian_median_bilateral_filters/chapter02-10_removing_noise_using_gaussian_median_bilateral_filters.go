package main

import (
	"gocv.io/x/gocv"
	"image"
	"log"
	"math/rand"
)

type Image []float32

func main() {
	imageFilePath := "../../data/Lena.png"
	originalMat := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if originalMat.Empty() {
		log.Panic("Can not read image: " + imageFilePath)
		return
	}

	originalImage := originalMat.ToBytes()
	originalFloat32Image := convertToFloat32(originalImage)

	// noised
	noisedImage := createNoisedImage(originalFloat32Image)

	noisedMat, err := gocv.NewMatFromBytes(originalMat.Rows(), originalMat.Cols(), gocv.MatTypeCV8UC3, noisedImage)
	if err != nil {
		log.Panic("Can not convert bytes to Mat")
		return
	}

	windowForNoisedImage := gocv.NewWindow("Noised Image")
	defer windowForNoisedImage.Close()
	windowForNoisedImage.IMShow(noisedMat)

	// gaussian
	gaussianBlurMat := gocv.NewMat()
	gocv.GaussianBlur(noisedMat, &gaussianBlurMat, image.Point{X: 7, Y: 7}, 0, 0, gocv.BorderConstant)

	windowForGaussian := gocv.NewWindow("Gaussian Filtered Image")
	defer windowForGaussian.Close()
	windowForGaussian.IMShow(gaussianBlurMat)

	// median
	medianBlurMat := gocv.NewMat()
	gocv.MedianBlur(noisedMat, &medianBlurMat, 7)

	windowForMedian := gocv.NewWindow("Median Filtered Image")
	defer windowForMedian.Close()
	windowForMedian.IMShow(medianBlurMat)

	// bilateral
	bilateralFilterMat := gocv.NewMat()
	gocv.BilateralFilter(noisedMat, &bilateralFilterMat, -1, 0.3, 10)

	windowForBilateralFilter := gocv.NewWindow("BilateralFilter Image")
	defer windowForBilateralFilter.Close()
	windowForBilateralFilter.IMShow(bilateralFilterMat)

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}

func convertToFloat32(original []uint8) Image {
	floatArray := make(Image, len(original))
	for i := 0; i < len(original); i++ {
		floatArray[i] = float32(original[i]) / 255
	}

	return floatArray
}

func convertToUint8(original Image) []uint8 {
	floatArray := make([]uint8, len(original))
	for i := 0; i < len(original); i++ {
		floatArray[i] = uint8(original[i] * 255)
	}

	return floatArray
}

func createNoisedImage(image []float32) []uint8 {
	noised := make(Image, len(image))
	for i := 0; i < len(image); i++ {
		noise := image[i] + 0.2*rand.Float32()
		if noise > 1 {
			noise = 1
		}

		if noise < 0 {
			noise = 0
		}

		noised[i] = noise
	}

	return convertToUint8(noised)
}
