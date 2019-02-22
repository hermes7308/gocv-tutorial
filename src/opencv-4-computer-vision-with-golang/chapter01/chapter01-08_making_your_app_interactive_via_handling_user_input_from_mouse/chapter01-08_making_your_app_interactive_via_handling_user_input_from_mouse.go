package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	mat := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if mat.Empty() {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	matCopied := gocv.NewMat()
	mat.CopyTo(&matCopied)

	for {
		// gocv has yet no mouse event handler
		// so I use selectROI for capture image
		rectangle := gocv.SelectROI("Original Image", mat)
		regionMat := matCopied.Region(rectangle)

		if !regionMat.Empty() {
			log.Println(regionMat)
			capture := gocv.NewWindow("Capture")
			capture.IMShow(regionMat)
		}
	}
}
