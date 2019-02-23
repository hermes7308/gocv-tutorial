package main

import (
	"gocv.io/x/gocv"
	"log"
)

const (
	WIDTH   = 480
	HEIGHT  = 640
	CHANNEL = 3
)

type Matrix []uint8

func main() {
	matrix := make(Matrix, WIDTH*HEIGHT*CHANNEL)

	// white
	matrix.changeColor(255, 255, 255)

	mat, err := gocv.NewMatFromBytes(HEIGHT, WIDTH, gocv.MatTypeCV8UC3, matrix)
	if err != nil {
		log.Panic("Can not change bytes to Mat")
		return
	}

	windowForWhite := gocv.NewWindow("White")
	defer windowForWhite.Close()
	windowForWhite.IMShow(mat)

	// red
	matrix.changeColor(0, 0, 255)

	mat, err = gocv.NewMatFromBytes(HEIGHT, WIDTH, gocv.MatTypeCV8UC3, matrix)
	if err != nil {
		log.Panic("Can not change bytes to Mat")
		return
	}

	windowForRed := gocv.NewWindow("Red")
	defer windowForRed.Close()
	windowForRed.IMShow(mat)

	// black
	matrix.changeColor(0, 0, 0)

	mat, err = gocv.NewMatFromBytes(HEIGHT, WIDTH, gocv.MatTypeCV8UC3, matrix)
	if err != nil {
		log.Panic("Can not change bytes to Mat")
		return
	}

	windowForBlack := gocv.NewWindow("Black")
	defer windowForBlack.Close()
	windowForBlack.IMShow(mat)

	// some white
	x := 240
	for _, y := range []int{160, 320, 480} {
		index := (WIDTH*y + x) * CHANNEL // getPixelIndex
		matrix[index] = 255
		matrix[index+1] = 255
		matrix[index+2] = 255
	}

	mat, err = gocv.NewMatFromBytes(HEIGHT, WIDTH, gocv.MatTypeCV8UC3, matrix)
	if err != nil {
		log.Panic("Can not change bytes to Mat")
		return
	}

	windowForSomeWhite := gocv.NewWindow("Some White")
	defer windowForSomeWhite.Close()
	windowForSomeWhite.IMShow(mat)

	// black to blue
	for i := 0; i < len(matrix); i += 3 {
		matrix[i] = 255
	}

	mat, err = gocv.NewMatFromBytes(HEIGHT, WIDTH, gocv.MatTypeCV8UC3, matrix)
	if err != nil {
		log.Panic("Can not change bytes to Mat")
		return
	}

	windowForBlackToBlue := gocv.NewWindow("Black to Blue")
	defer windowForBlackToBlue.Close()
	windowForBlackToBlue.IMShow(mat)

	// center vertical white line
	x = 320
	for y := 0; y < HEIGHT; y++ {
		index := (WIDTH*y + x) * CHANNEL // getPixelIndex
		matrix[index] = 255
		matrix[index+1] = 255
		matrix[index+2] = 255
	}

	mat, err = gocv.NewMatFromBytes(HEIGHT, WIDTH, gocv.MatTypeCV8UC3, matrix)
	if err != nil {
		log.Panic("Can not change bytes to Mat")
		return
	}

	windowForCenterVerticalWhiteLine := gocv.NewWindow("Center Vertical White Line")
	defer windowForCenterVerticalWhiteLine.Close()
	windowForCenterVerticalWhiteLine.IMShow(mat)

	// certain region
	x1, y1 := 100, 100
	x2, y2 := 200, 600

	for x = x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			index := (WIDTH*y + x) * CHANNEL // getPixelIndex
			matrix[index] = 255
			matrix[index+1] = 255
			matrix[index+2] = 255

		}
	}

	mat, err = gocv.NewMatFromBytes(HEIGHT, WIDTH, gocv.MatTypeCV8UC3, matrix)
	if err != nil {
		log.Panic("Can not change bytes to Mat")
		return
	}

	windowForCertainRegion := gocv.NewWindow("Certain region")
	defer windowForCertainRegion.Close()
	windowForCertainRegion.IMShow(mat)

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}

func (matrix *Matrix) changeColor(blue, green, red uint8) {
	for i := 0; i < len(*matrix); i += 3 {
		(*matrix)[i] = blue    // Blue
		(*matrix)[i+1] = green // Green
		(*matrix)[i+2] = red   // Red
	}
}
