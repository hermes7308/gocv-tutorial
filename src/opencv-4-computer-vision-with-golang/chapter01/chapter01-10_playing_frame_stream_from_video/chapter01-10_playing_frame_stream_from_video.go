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
	defer capture.Close()

	window := gocv.NewWindow("Video")
	defer window.Close()

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
}
