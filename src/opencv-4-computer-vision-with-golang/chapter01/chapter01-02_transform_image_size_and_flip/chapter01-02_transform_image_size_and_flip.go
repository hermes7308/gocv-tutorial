package main

import (
	"image"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	mat := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if mat.Empty() {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	log.Println("original image size:", mat.Size())

	width, height := 0.25, 0.5

	resizeImage := gocv.NewMat()
	gocv.Resize(mat, &resizeImage, image.Point{X: 0, Y: 0}, width, height, gocv.InterpolationNearestNeighbor)
	log.Println("original image size:", resizeImage.Size())

	flippedImage := gocv.NewMat()
	// Flip flips a 2D array around horizontal(0), vertical(1), or both axes(-1)
	// Reverses with respect to the x axis.
	gocv.Flip(mat, &flippedImage, 0)
	// Reverses with respect to the y axis.
	gocv.Flip(mat, &flippedImage, 1)
	// Reverses with respect to the both x, y axis.
	gocv.Flip(mat, &flippedImage, -1)
}
