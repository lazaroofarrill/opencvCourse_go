package main

import "gocv.io/x/gocv"

func main() {
	path := "Resources/shapes.png"

	img := gocv.IMRead(path, gocv.IMReadUnchanged)
	imgHSV := gocv.NewMat()
	mask := gocv.NewMat()

	gocv.CvtColor(img, &imgHSV, gocv.ColorBGRToHSV)
	hmin, smin, vmin := 0.0, 110.0, 153.0
	hmax, smax, vmax := 19.0, 240.0, 255.0
	lower := gocv.Scalar{Val1: hmin, Val2: smin, Val3: vmin}
	upper := gocv.Scalar{Val1: hmax, Val2: smax, Val3: vmax}
	gocv.InRangeWithScalar(imgHSV, lower, upper, &mask)

	win := gocv.NewWindow("lambo")
	sizes := img.Size()
	win.ResizeWindow(sizes[1], sizes[0])
	win2 := gocv.NewWindow("hsv")
	win2.ResizeWindow(sizes[1], sizes[0])
	win3 := gocv.NewWindow("mask")
	win3.ResizeWindow(sizes[1], sizes[0])

	imgControls := gocv.NewWindow("Trackbars")
	imgControls.ResizeWindow(640, 200)
	hueMin := imgControls.CreateTrackbar("H min", 179)
	hueMin.SetPos(0)
	hueMax := imgControls.CreateTrackbar("H max", 179)
	hueMax.SetPos(19)
	satMin := imgControls.CreateTrackbar("S min", 255)
	satMin.SetPos(110)
	satMax := imgControls.CreateTrackbar("S max", 255)
	satMax.SetPos(240)
	valMin := imgControls.CreateTrackbar("V min", 255)
	valMin.SetPos(153)
	valMax := imgControls.CreateTrackbar("V max", 255)
	valMax.SetPos(255)

	for {
		hmin, smin, vmin := float64(hueMin.GetPos()), float64(satMin.GetPos()), float64(valMin.GetPos())
		hmax, smax, vmax := float64(hueMax.GetPos()), float64(satMax.GetPos()), float64(valMax.GetPos())
		lower := gocv.Scalar{Val1: hmin, Val2: smin, Val3: vmin}
		upper := gocv.Scalar{Val1: hmax, Val2: smax, Val3: vmax}
		gocv.InRangeWithScalar(imgHSV, lower, upper, &mask)

		win.IMShow(img)
		win2.IMShow(imgHSV)
		win3.IMShow(mask)
		gocv.WaitKey(1)
	}
}
