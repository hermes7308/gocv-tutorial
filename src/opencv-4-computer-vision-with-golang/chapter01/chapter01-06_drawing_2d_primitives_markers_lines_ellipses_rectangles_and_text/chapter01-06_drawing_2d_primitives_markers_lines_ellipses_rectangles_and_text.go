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

	// Circle
	gocv.Circle(&mat, getRandomPoint(), 40, color.RGBA{255, 0, 0, 0}, 1)
	gocv.Circle(&mat, getRandomPoint(), 5, color.RGBA{255, 0, 0, 0}, 2)
	gocv.Circle(&mat, getRandomPoint(), 40, color.RGBA{255, 85, 85, 0}, 2)
	gocv.Circle(&mat, getRandomPoint(), 40, color.RGBA{255, 170, 170, 0}, 2)

	// Line
	gocv.Line(&mat, getRandomPoint(), getRandomPoint(), color.RGBA{0, 255, 0, 0}, 1)
	gocv.Line(&mat, getRandomPoint(), getRandomPoint(), color.RGBA{85, 255, 85, 0}, 3)
	gocv.Line(&mat, getRandomPoint(), getRandomPoint(), color.RGBA{255, 170, 170, 0}, 3)

	// ArrowLine
	gocv.ArrowedLine(&mat, getRandomPoint(), getRandomPoint(), color.RGBA{0, 0, 255, 0}, 3)

	// Rectangle
	gocv.Rectangle(&mat, getRandomRectangle(), color.RGBA{255, 255, 0, 0}, 3)

	// Ellipse
	gocv.Ellipse(&mat, getRandomPoint(), getRandomPoint(), rand.Float64(), 0, 360, color.RGBA{255, 255, 255, 0}, 3)

	// Text
	gocv.PutText(&mat, "OpenCV", getRandomPoint(), gocv.FontHersheySimplex, 1, color.RGBA{0, 0, 0, 0}, 3)

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

func getRandomRectangle() image.Rectangle {
	return image.Rect(rand.Intn(width), rand.Intn(height), rand.Intn(width), rand.Intn(height))
}
