package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	capture, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Panic("Can not find Camera...")
		return
	}
	defer capture.Close()

	window := gocv.NewWindow("Camera")
	defer window.Close()

	mat := gocv.NewMat()

	for {
		capture.Read(&mat)
		window.IMShow(mat)

		key := window.WaitKey(1)
		if key == 27 {
			break
		}
	}
}
