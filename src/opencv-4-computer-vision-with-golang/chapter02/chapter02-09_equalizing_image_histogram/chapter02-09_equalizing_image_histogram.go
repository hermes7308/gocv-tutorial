package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
)

func main() {
	imageFilePath := "../../data/Lena.png"
	originalMat := gocv.IMRead(imageFilePath, gocv.IMReadGrayScale)
	if originalMat.Empty() {
		log.Panic("Can not read image: " + imageFilePath)
		return
	}

	// grey equalizeHist
	greyEqualizeHistMat := gocv.NewMat()
	gocv.EqualizeHist(originalMat, &greyEqualizeHistMat)

	windowForGreyEqualizeHist := gocv.NewWindow("Window For Grey EqualizeHist")
	defer windowForGreyEqualizeHist.Close()
	windowForGreyEqualizeHist.IMShow(greyEqualizeHistMat)

	// plot
	greyEqualizeHistImage := greyEqualizeHistMat.ToBytes()

	v := make(plotter.Values, len(greyEqualizeHistImage))
	for i := 0; i < len(v); i++ {
		v[i] = float64(greyEqualizeHistImage[i])
	}

	plot, err := plot.New()
	if err != nil {
		log.Println("Can not create plot", err)
		return
	}

	plot.Title.Text = fmt.Sprintf("pixel value")

	histogram, err := plotter.NewHist(v, len(v))
	if err != nil {
		log.Panic("Can not create histogram", err)
		return
	}

	histogram.Normalize(1)
	plot.Add(histogram)

	if err := plot.Save(4*vg.Inch, 4*vg.Inch, "pixel_value.png"); err != nil {
		log.Panic(err)
		return
	}

	// wait
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("Pressed ESC")
			break
		}
	}
}
