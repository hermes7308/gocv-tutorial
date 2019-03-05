package main

import (
	"gocv.io/x/gocv"
	"image"
	"log"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	originalMat := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if originalMat.Empty() {
		log.Panic("Can not read image: " + imageFilePath)
		return
	}

	gaussianBlurMat := gocv.NewMat()
	gocv.GaussianBlur(originalMat, &gaussianBlurMat, image.Point{X: 3, Y: 3}, 0, 0, gocv.BorderDefault)

	grayMat := gocv.NewMat()
	gocv.CvtColor(gaussianBlurMat, &grayMat, gocv.ColorBGRToGray)

	windowForGray := gocv.NewWindow("Gray Image")
	defer windowForGray.Close()
	windowForGray.IMShow(grayMat)

	// dx
	dx := gocv.NewMat()
	gocv.Sobel(grayMat, &dx, gocv.MatTypeCV16S, 1, 0, 5, 1, 0, gocv.BorderDefault)
	gocv.ConvertScaleAbs(dx, &dx, 1, 0)

	windowForDx := gocv.NewWindow("dx Image")
	defer windowForDx.Close()
	windowForDx.IMShow(dx)

	// dy
	dy := gocv.NewMat()
	gocv.Sobel(grayMat, &dy, gocv.MatTypeCV16S, 0, 1, 5, 1, 0, gocv.BorderDefault)
	gocv.ConvertScaleAbs(dy, &dy, 1, 0)

	windowForDy := gocv.NewWindow("dy Image")
	defer windowForDy.Close()
	windowForDy.IMShow(dy)

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}
