package main

import (
	"gocv.io/x/gocv"
	"io/ioutil"
	"log"
	"math/rand"
)

const (
	WIDTH   = 300
	HEIGHT  = 300
	CHANNEL = 3
)

func main() {
	filePath := "./random.cvs"

	newImage := createRandomImage(WIDTH, HEIGHT, CHANNEL)

	// save file
	err := ioutil.WriteFile(filePath, newImage, 0644)
	if err != nil {
		log.Panic("Can not write file", err)
		return
	}

	loadImage, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic("Can not read file", err)
		return
	}

	// show image
	mat, err := gocv.NewMatFromBytes(WIDTH, HEIGHT, gocv.MatTypeCV8UC3, loadImage)
	if err != nil {
		log.Panic("Can not convert bytes to Mat", err)
		return
	}

	window := gocv.NewWindow("Random Image")
	defer window.Close()
	window.IMShow(mat)

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}

func createRandomImage(rows, cols, channel int) []uint8 {
	image := make([]uint8, rows*cols*channel)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for ch := 0; ch < channel; ch++ {
				index := (rows*row+col)*channel + ch
				image[index] = uint8(rand.Intn(255))
			}
		}
	}
	return image
}
