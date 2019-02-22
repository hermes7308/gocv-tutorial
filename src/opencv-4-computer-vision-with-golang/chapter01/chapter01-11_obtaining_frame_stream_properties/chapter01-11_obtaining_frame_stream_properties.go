package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gocv.io/x/gocv"
	"log"
)

func main() {
	videoPath := "../../data/drop.avi"
	err := printCaptureProperties(videoPath)
	if err != nil {
		log.Panic("Could not print capture properties: " + videoPath)
	}

	cameraNo := 0
	err = printCaptureProperties(cameraNo)
	if err != nil {
		log.Panic("Could not print capture properties: " + string(cameraNo))
	}
}

func printCaptureProperties(param interface{}) error {
	var capture *gocv.VideoCapture
	var err error

	switch param.(type) {
	case int:
		capture, err = gocv.VideoCaptureDevice(param.(int))
	case string:
		capture, err = gocv.VideoCaptureFile(param.(string))
	default:
		return errors.New(fmt.Sprintf("%v is not supported type", param))
	}

	if err != nil {
		log.Panic("Can not find capture source")
		return err
	}

	log.Println("--------------------")
	log.Println("Frame count:", capture.Get(gocv.VideoCaptureFrameCount))
	log.Println("Frame width:", capture.Get(gocv.VideoCaptureFrameWidth))
	log.Println("Frame height:", capture.Get(gocv.VideoCaptureFrameHeight))
	log.Println("Frame rate:", capture.Get(gocv.VideoCaptureFPS))

	return nil
}
