package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w int
	h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}
func (i Image) At(x, y int) color.Color {
	m := uint8((x * y) * (y * y))
	return color.RGBA{m, m, 255, 255}
}
func main() {
	m := Image{200, 200}
	pic.ShowImage(m)
}
