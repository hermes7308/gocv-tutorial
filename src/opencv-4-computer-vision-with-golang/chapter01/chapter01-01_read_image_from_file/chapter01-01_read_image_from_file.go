package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	mat := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if mat.Empty() {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	log.Println("read", imageFilePath)
	log.Println("size:", mat.Size())
	log.Println("type:", mat.Type())

	mat = gocv.IMRead(imageFilePath, gocv.IMReadGrayScale)
	if &mat == nil {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	log.Println("read", imageFilePath)
	log.Println("size:", mat.Size())
	log.Println("type:", mat.Type())
}
