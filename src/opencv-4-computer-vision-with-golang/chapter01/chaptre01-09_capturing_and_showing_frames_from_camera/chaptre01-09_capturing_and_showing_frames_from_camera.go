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

	window := gocv.NewWindow("Camera")
	mat := gocv.NewMat()

	for {
		capture.Read(&mat)
		window.IMShow(mat)

		key := window.WaitKey(1)
		if key == 27 {
			break
		}
	}

	err = capture.Close()
	if err != nil {
		log.Panic("Can not close Camera")
	}

	err = window.Close()
	if err != nil {
		log.Panic("Can not close Window")
	}
}
