package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	windowForTrackbar := gocv.NewWindow("Trackbar Window")
	trackbarR := windowForTrackbar.CreateTrackbar("R", 255)
	trackbarG := windowForTrackbar.CreateTrackbar("G", 255)
	trackbarB := windowForTrackbar.CreateTrackbar("B", 255)

	mat := gocv.NewMatWithSize(300, 300, gocv.MatTypeCV8UC3)

	log.Println("-----------------------")
	log.Println("Rows:", mat.Rows())
	log.Println("Cols:", mat.Cols())

	for {
		// BGR
		blue := float64(trackbarB.GetPos())
		green := float64(trackbarG.GetPos())
		red := float64(trackbarR.GetPos())

		log.Println("-----------------------")
		log.Println("blue:", blue)
		log.Println("green:", green)
		log.Println("red:", red)

		mat.SetTo(gocv.NewScalar(blue, green, red, 0.0))

		windowForTrackbar.IMShow(mat)

		key := windowForTrackbar.WaitKey(3)
		if key == 27 {
			break
		}
	}

	windowForTrackbar.Close()
}
