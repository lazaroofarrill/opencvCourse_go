package main

import (
	"gocv.io/x/gocv"
	"image"
)

func main() {
	path := "Resources/lena.png"

	img := gocv.IMRead(path, gocv.IMReadUnchanged)
	imgResize := gocv.NewMat()
	defer img.Close()
	defer imgResize.Close()

	println(img.Size())
	gocv.Resize(img, &imgResize, image.Point{}, 0.5, 0.5, gocv.InterpolationLinear)

	winOg := gocv.NewWindow("og")
	defer winOg.Close()
	winOg.ResizeWindow(img.Size()[0], img.Size()[1])
	winRes := gocv.NewWindow("resized")
	defer winRes.Close()
	winRes.ResizeWindow(imgResize.Size()[1], imgResize.Size()[0])

	roi := image.Rectangle{
		Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 300, Y: 450},
	}

	region := img.Region(roi)

	cropped := region.Clone()
	region.Close()
	defer cropped.Close()

	winCrop := gocv.NewWindow("cropped")
	defer winCrop.Close()
	sizes := cropped.Size()
	winCrop.ResizeWindow(sizes[1], sizes[0])

	for {
		winOg.IMShow(img)
		winRes.IMShow(imgResize)
		winCrop.IMShow(cropped)
		gocv.WaitKey(0)
	}
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
