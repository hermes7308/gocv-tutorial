package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	originalMat := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if originalMat.Empty() {
		log.Panic("Can not read image: " + imageFilePath)
		return
	}

	// gray
	grayMat := gocv.NewMat()
	gocv.CvtColor(originalMat, &grayMat, gocv.ColorBGRToGray)

	windowForGray := gocv.NewWindow("Gray Image")
	defer windowForGray.Close()
	windowForGray.IMShow(grayMat)

	// hsv
	hsvMat := gocv.NewMat()
	gocv.CvtColor(originalMat, &hsvMat, gocv.ColorBGRToHSV)

	windowForHsv := gocv.NewWindow("HSV Image")
	defer windowForHsv.Close()
	windowForHsv.IMShow(hsvMat)

	// hsv extension
	hsvImage := hsvMat.ToBytes()
	for i := 2; i < len(hsvImage); i += 3 {
		hsvImage[i] = hsvImage[i] * 2
	}

	hsvExtensionMat, err := gocv.NewMatFromBytes(hsvMat.Rows(), hsvMat.Cols(), gocv.MatTypeCV8UC3, hsvImage)
	if err != nil {
		log.Panic("Can not convert bytes to Mat")
		return
	}

	gocv.CvtColor(hsvExtensionMat, &hsvExtensionMat, gocv.ColorHSVToBGR)

	windowForHsvExtension := gocv.NewWindow("HSV Extension Image")
	defer windowForHsvExtension.Close()
	windowForHsvExtension.IMShow(hsvExtensionMat)

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}
