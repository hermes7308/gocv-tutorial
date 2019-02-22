package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	capture, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Panic("Can not find capture source")
		return
	}

	frameWidth := int(capture.Get(gocv.VideoCaptureFrameWidth))
	frameHeight := int(capture.Get(gocv.VideoCaptureFrameHeight))

	log.Println("Frame width:", frameWidth)
	log.Println("Frame height:", frameHeight)

	videoWriter, err := gocv.VideoWriterFile("./captured_video.avi", "X264", 25, frameWidth, frameHeight, true)
	if err != nil {
		log.Panic("Can not open video writer")
		return
	}

	window := gocv.NewWindow("Frame")
	mat := gocv.NewMat()

	for {
		capture.Read(&mat)

		if mat.Empty() {
			log.Fatal("Can not get frame")
			break
		}

		err = videoWriter.Write(mat)
		if err != nil {
			log.Panic("Can not write frame")
			break
		}

		window.IMShow(mat)

		key := window.WaitKey(1)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}
