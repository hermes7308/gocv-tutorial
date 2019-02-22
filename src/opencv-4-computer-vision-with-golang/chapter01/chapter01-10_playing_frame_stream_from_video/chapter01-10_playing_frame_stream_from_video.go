package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	capture, err := gocv.VideoCaptureFile("../../data/drop.avi")
	if err != nil {
		log.Panic("Can not find video")
		return
	}

	window := gocv.NewWindow("Video")
	mat := gocv.NewMat()

	for {
		ok := capture.Read(&mat)
		if !ok {
			log.Println("Reached the end of the video")
			break
		}

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
