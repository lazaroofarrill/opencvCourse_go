package main

import (
	. "gocv.io/x/gocv"
	"image"
)

const h = 350
const w = 250

func main() {
	path := "Resources/cards.jpg"

	img := IMRead(path, IMReadUnchanged)

	src := []image.Point{{X: 529, Y: 142}, {X: 771, Y: 190}, {X: 405, Y: 395}, {X: 674, Y: 457}}
	dst := []image.Point{{X: 0, Y: 0}, {X: w, Y: 0}, {X: 0, Y: h}, {X: w, Y: h}}

	srcV := NewPointVectorFromPoints(src)
	dstV := NewPointVectorFromPoints(dst)

	tm := GetPerspectiveTransform(srcV, dstV)
	warpedImage := NewMat()

	WarpPerspective(img, &warpedImage, tm, image.Point{X: w, Y: h})

	win := NewWindow("cards")
	size := warpedImage.Size()
	win.ResizeWindow(size[1], size[0])

	for {
		win.IMShow(warpedImage)
		WaitKey(0)
	}
}
