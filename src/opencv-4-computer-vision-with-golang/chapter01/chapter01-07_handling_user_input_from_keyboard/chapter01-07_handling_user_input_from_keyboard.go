package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"log"
	"math/rand"
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

	window := gocv.NewWindow("Handling user input from keyboard")
	width, height = mat.Rows(), mat.Cols()

	matCopied := gocv.NewMat()
	mat.CopyTo(&matCopied)

	running := true

	for running {
		window.IMShow(matCopied)

		key := gocv.WaitKey(0)

		switch key {
		case int('p'):
			for i := 0; i < 10; i++ {
				gocv.Circle(&matCopied, getRandomPoint(), 3, color.RGBA{255, 0, 0, 0}, -1)
			}
		case int('l'):
			gocv.Line(&matCopied, getRandomPoint(), getRandomPoint(), color.RGBA{0, 255, 0, 0}, 3)
		case int('r'):
			gocv.Rectangle(&matCopied, getRandomRectangle(), color.RGBA{0, 0, 255, 0}, 3)
		case int('e'):
			gocv.Ellipse(&matCopied, getRandomPoint(), getRandomPoint(), rand.Float64(), 0, 360, color.RGBA{255, 255, 0, 0}, 3)
		case int('t'):
			gocv.PutText(&matCopied, "OpenCV", getRandomPoint(), gocv.FontHersheySimplex, 1, color.RGBA{0, 0, 0, 0}, 3)
		case int('c'):
			mat.CopyTo(&matCopied)
		case 27:
			running = false
		}
	}

}

func getRandomPoint() image.Point {
	return image.Point{X: rand.Intn(width), Y: rand.Intn(height)}
}

func getRandomRectangle() image.Rectangle {
	return image.Rect(rand.Intn(width), rand.Intn(height), rand.Intn(width), rand.Intn(height))
}
