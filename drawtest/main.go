package main

import (
	"log"
	"image"
	"image/color"
	"github.com/llgcode/draw2d/draw2dimg"
)

var ICON_SIZE = 72

func main() {
	drawFactionProgress()
}

func drawFactionProgress() {
	dest := image.NewRGBA(image.Rect(0, 0, ICON_SIZE, ICON_SIZE))
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetLineWidth(5)

	progress := 1500.0 / 2000.0
	distance := 204.0

	dashLength := progress * distance

	log.Println(dashLength)

	gc.SetLineDash([]float64{dashLength, 300}, 0)

	gc.MoveTo(72/2, 0)
	gc.LineTo(72,72/2)
	gc.LineTo(72/2,72)
	gc.LineTo(0,72/2)
	gc.Stroke()

	draw2dimg.SaveToPngFile("hello.png", dest)
}