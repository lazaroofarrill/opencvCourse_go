package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math/rand"
)

func main() {
	path := "Resources/shapes.png"
	img := gocv.IMRead(path, gocv.IMReadUnchanged)

	imgGray, imgBlur, imgCanny, imgDil := gocv.NewMat(), gocv.NewMat(), gocv.NewMat(), gocv.NewMat()
	imgContours := gocv.NewMat()

	gocv.CvtColor(img, &imgGray, gocv.ColorBGRToGray)
	gocv.GaussianBlur(imgGray, &imgBlur, image.Point{X: 3, Y: 3}, 3, 0, gocv.BorderConstant)
	gocv.Canny(imgBlur, &imgCanny, 25, 75)

	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Point{X: 3, Y: 3})
	gocv.Dilate(imgCanny, &imgDil, kernel)

	drawContours(imgDil, imgContours)

	win := gocv.NewWindow("img")
	win2 := gocv.NewWindow("Image Gray")
	win3 := gocv.NewWindow("Image Blur")
	win4 := gocv.NewWindow("Image Canny")
	win5 := gocv.NewWindow("Image Dil")
	win6 := gocv.NewWindow("Image Contours")

	for {
		win.IMShow(img)
		win2.IMShow(imgGray)
		win3.IMShow(imgBlur)
		win4.IMShow(imgCanny)
		win5.IMShow(imgDil)
		win6.IMShow(imgContours)
		gocv.WaitKey(1)
	}
}

/*
Do all kinds of crazy stuff to the image
*/
func drawContours(img gocv.Mat, dst gocv.Mat) {
	hierarchy := gocv.NewMat()
	contours := gocv.FindContoursWithParams(img, &hierarchy, gocv.RetrievalTree, gocv.ChainApproxSimple)
	println(hierarchy.Type().String())
	hierarchyDims := hierarchy.Size()

	counter := 0
	var hierarchyMatrix [][]int32
	for j := 0; j < hierarchyDims[1]*4; j += 4 {
		hierarchyMatrix = append(hierarchyMatrix, []int32{})
		for count := 0; count < 4; count++ {
			hierarchyMatrix[j/4] = append(hierarchyMatrix[j/4], hierarchy.GetIntAt(0, j+count))
			fmt.Printf("%3v", hierarchyMatrix[j/4][count])
		}
		counter++
		println()
	}
	println(counter)

	err := dst.Close()
	if err != nil {
		panic(err)
	}

	sizes := img.Size()
	dst = gocv.NewMatWithSizes(sizes, gocv.MatTypeCV64FC4)
	conPoly := gocv.NewPointsVector()
	objectType := ""
	rand.Seed(1400)
	for i := 0; i < contours.Size(); i++ {
		area := gocv.ContourArea(contours.At(i))

		parent := hierarchyMatrix[i][3]
		if area > 1000 && parent == 1 {
			peri := gocv.ArcLength(contours.At(i), true)
			conPoly.Append(gocv.ApproxPolyDP(contours.At(i), 0.02*peri, true))
			gocv.DrawContours(&dst, conPoly, conPoly.Size()-1, color.RGBA{B: 69, R: 255}, 1)
			rect := gocv.BoundingRect(conPoly.At(conPoly.Size() - 1))

			gocv.Rectangle(&dst, rect, color.RGBA{G: 255, A: 5}, 1)
			switch conPoly.At(conPoly.Size() - 1).Size() {
			case 3:
				objectType = "Tri"
			case 4:
				aspRatio := float32(rect.Dx()) / float32(rect.Dy())
				if aspRatio > 0.95 && aspRatio < 1.05 {
					objectType = "Square"
				} else {
					objectType = "Rect"
				}
			default:
				objectType = "Circle"
			}
			//println(objectType)
			gocv.PutText(&dst, objectType, rect.Min, gocv.FontHersheyPlain, 1, color.RGBA{R: 255, G: 255, B: 255}, 1)
		}
	}
	fmt.Printf("%v significant contours found", conPoly.Size())
}
