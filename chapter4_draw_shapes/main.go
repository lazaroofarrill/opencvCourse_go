package main

import (
	. "gocv.io/x/gocv"
	"image"
	. "image/color"
)

func main() {
	img := NewMatWithSizeFromScalar(Scalar{Val1: 255, Val2: 255, Val3: 255},
		512, 512, MatTypeCV8UC3)
	defer img.Close()

	Circle(&img, image.Point{X: 256, Y: 256}, 155, RGBA{R: 255, G: 69}, -1)
	Rectangle(&img, image.Rectangle{
		Min: image.Point{X: 130, Y: 226},
		Max: image.Point{X: 382, Y: 286},
	}, RGBA{R: 255, G: 255, B: 255}, -1)
	Line(&img, image.Point{X: 130, Y: 296}, image.Point{X: 382, Y: 296}, RGBA{R: 255, G: 255, B: 255}, 2)

	PutText(&img, "Murtaza's Workshop", image.Point{X: 137, Y: 262}, FontHersheyDuplex,
		0.75, RGBA{R: 255, G: 69}, 2)

	win := NewWindow("img")
	sizes := img.Size()
	win.ResizeWindow(sizes[1], sizes[0])
	for {
		win.IMShow(img)
		WaitKey(0)
	}
}
