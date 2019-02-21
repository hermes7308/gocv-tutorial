package main

import (
	"image"
	"image/color"
	"log"
	"math/rand"

	"gocv.io/x/gocv"
)

var (
	width  int
	height int
)

func main() {
	imageFilePath := "../../data/Lena.png"
	mat := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if mat.Empty() {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	window := gocv.NewWindow("Drawing")
	width, height = mat.Rows(), mat.Cols()

	gocv.Circle(&mat, getRandomPoint(), 40, color.RGBA{255, 0, 0, 0}, 1)

	for {
		window.IMShow(mat)

		key := window.WaitKey(3)
		if key == 27 {
			break
		}
	}
}

func getRandomPoint() image.Point {
	return image.Point{X: rand.Intn(width), Y: rand.Intn(height)}
}
