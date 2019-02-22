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

	// PNG Lossless Compression
	compressedImageFilePathForPNG := "./Lena_compressed.png"
	gocv.IMWriteWithParams(compressedImageFilePathForPNG, mat, []int{gocv.IMWritePngCompression, 0})

	saveImage := gocv.IMRead(compressedImageFilePathForPNG, gocv.IMReadAnyColor)
	if &saveImage == nil {
		log.Panic("Can not read Image file : ", compressedImageFilePathForPNG)
		return
	}

	// JPG Lossy Compression
	compressedImageFilePathForJPG := "./Lena_compressed.jpg"
	gocv.IMWriteWithParams(compressedImageFilePathForJPG, mat, []int{gocv.IMWriteJpegQuality, 0})
}
