package main

import (
	. "gocv.io/x/gocv"
	"image"
)

func main() {
	path := "Resources/lena.png"
	img := IMRead(path, IMReadColor)

	if img.Empty() {
		panic("img not read")
	}

	imgGray := NewMat()
	imgBlur := NewMat()
	imgCanny := NewMat()
	imgDilation := NewMat()
	imgErotion := NewMat()
	kernel := GetStructuringElement(MorphRect, image.Point{X: 5, Y: 5})

	CvtColor(img, &imgGray, ColorBGRToGray)
	GaussianBlur(img, &imgBlur, image.Point{X: 7, Y: 7}, 5, 0, BorderConstant)
	Canny(imgBlur, &imgCanny, 25, 85)
	Dilate(imgCanny, &imgDilation, kernel)
	Erode(imgDilation, &imgErotion, kernel)

	win := NewWindow("img")
	sizes := img.Size()
	win.ResizeWindow(sizes[0], sizes[1])
	win2 := NewWindow("imgGray")
	win2.ResizeWindow(sizes[0], sizes[1])

	win3 := NewWindow("img blur")
	win3.ResizeWindow(sizes[0], sizes[1])

	win4 := NewWindow("img dilation")
	win4.ResizeWindow(sizes[0], sizes[1])

	win5 := NewWindow("img erotion")
	win5.ResizeWindow(sizes[0], sizes[1])

	for {
		win.IMShow(img)
		win2.IMShow(imgGray)
		win3.IMShow(imgCanny)
		win4.IMShow(imgDilation)
		win5.IMShow(imgErotion)
		WaitKey(0)
	}
}
