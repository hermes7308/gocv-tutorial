package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	printCapturePropertiesFromVideo("../../data/drop.avi")
	printCapturePropertiesFromCamera(0)
}

func printCapturePropertiesFromCamera(cameraNumber int) {
	log.Println("Created capture:", cameraNumber)
	capture, err := gocv.VideoCaptureDevice(0)
	printCaptureProperties(capture, err)
}

func printCapturePropertiesFromVideo(videoPath string) {
	log.Println("Created capture:", videoPath)
	capture, err := gocv.VideoCaptureFile(videoPath)
	printCaptureProperties(capture, err)
}

func printCaptureProperties(capture *gocv.VideoCapture, err error) {
	if err != nil {
		log.Panic("Can not find capture source")
		return
	}

	log.Println("Frame count:", capture.Get(gocv.VideoCaptureFrameCount))
	log.Println("Frame width:", capture.Get(gocv.VideoCaptureFrameWidth))
	log.Println("Frame height:", capture.Get(gocv.VideoCaptureFrameHeight))
	log.Println("Frame rate:", capture.Get(gocv.VideoCaptureFPS))
}
