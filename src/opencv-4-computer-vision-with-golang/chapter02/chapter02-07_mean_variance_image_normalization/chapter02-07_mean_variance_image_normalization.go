package main

import (
	"gocv.io/x/gocv"
	"log"
	"math"
)

type Image []float32

func main() {
	imageFilePath := "../../data/Lena.png"
	originalMat := gocv.IMRead(imageFilePath, gocv.IMReadGrayScale)
	if originalMat.Empty() {
		log.Panic("Can not read image: " + imageFilePath)
		return
	}

	originalImage := originalMat.ToBytes()
	targetImage := make(Image, len(originalImage))
	for i := 0; i < len(targetImage); i++ {
		targetImage[i] = float32(originalImage[i]) / 255
	}

	targetImage.zeroAverageMatrix()
	log.Println(targetImage)

	targetImage.dispersionMatrix()
	log.Println(targetImage)
}

func (image *Image) zeroAverageMatrix() {
	average := image.average()

	for i := 0; i < len(*image); i++ {
		(*image)[i] = (*image)[i] - average
	}
}

func (image *Image) dispersionMatrix() {
	average := image.average()
	tempImage := image.clone()
	length := len(tempImage)

	for i := 0; i < length; i++ {
		tempImage[i] = float32(math.Pow(float64(tempImage[i]-average), 2))
	}

	sum := tempImage.sum()

	standardDeviation := sum / float32(length-1)

	for i := 0; i < length; i++ {
		(*image)[i] = (*image)[i] / standardDeviation
	}
}

func (image Image) sum() float32 {
	length := len(image)
	var sum float32 = 0
	for i := 0; i < length; i++ {
		sum = sum + image[i]
	}

	return sum
}

func (image Image) average() float32 {
	length := len(image)
	sum := image.sum()

	return sum / float32(length)
}

func (image Image) clone() Image {
	length := len(image)
	cloneImage := make([]float32, length)
	for i := 0; i < length; i++ {
		cloneImage[i] = image[i]
	}

	return cloneImage
}
