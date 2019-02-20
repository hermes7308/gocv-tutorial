package main

import (
	"log"

	"gocv.io/x/gocv"
)

type aa struct {
	ab int
}

func main() {
	imageFilePath := "../../data/Lena.png"
	img := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if &img == nil {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	log.Println("read", imageFilePath)
	log.Println("size:", img.Size())
	log.Println("type:", img.Type())

	img = gocv.IMRead(imageFilePath, gocv.IMReadGrayScale)
	if &img == nil {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	log.Println("read", imageFilePath)
	log.Println("size:", img.Size())
	log.Println("type:", img.Type())
}
