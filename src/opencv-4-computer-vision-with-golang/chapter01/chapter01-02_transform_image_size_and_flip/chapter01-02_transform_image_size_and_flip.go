package main

import (
	"image"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	img := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if img.Empty() {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	log.Println("original image size:", img.Size())

	width, height := 0.25, 0.5

	resized_img := gocv.NewMat()
	gocv.Resize(img, &resized_img, image.Point{X: 0, Y: 0}, width, height, gocv.InterpolationNearestNeighbor)
	log.Println("original image size:", resized_img.Size())

	flipped_img := gocv.NewMat()
	// Flip flips a 2D array around horizontal(0), vertical(1), or both axes(-1)
	// Reverses with respect to the x axis.
	gocv.Flip(img, &flipped_img, 0)
	// Reverses with respect to the y axis.
	gocv.Flip(img, &flipped_img, 1)
	// Reverses with respect to the both x, y axis.
	gocv.Flip(img, &flipped_img, -1)
}
