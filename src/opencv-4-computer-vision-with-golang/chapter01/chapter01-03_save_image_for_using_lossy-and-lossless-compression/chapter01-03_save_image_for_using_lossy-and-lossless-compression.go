package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	img := gocv.IMRead(imageFilePath, gocv.IMReadAnyColor)
	if img.Empty() {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	// PNG Lossless Compression
	compressedImageFilePathForPNG := "./Lena_compressed.png"
	gocv.IMWriteWithParams(compressedImageFilePathForPNG, img, []int{gocv.IMWritePngCompression, 0})

	save_img := gocv.IMRead(compressedImageFilePathForPNG, gocv.IMReadAnyColor)
	if &save_img == nil {
		log.Panic("Can not read Image file : ", compressedImageFilePathForPNG)
		return
	}

	// JPG Lossy Compression
	compressedImageFilePathForJPG := "./Lena_compressed.jpg"
	gocv.IMWriteWithParams(compressedImageFilePathForJPG, img, []int{gocv.IMWriteJpegQuality, 0})
}
