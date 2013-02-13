package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	bounds image.Rectangle
	pixels [][]uint8
}

func SampleImage(width, height int) *Image {
	return &Image{
		image.Rectangle{image.ZP, image.Point{width, height}},
		Pic(width, height),
	}
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return img.bounds
}

func (img Image) At(x, y int) color.Color {
	v := img.pixels[x][y]
	return color.RGBA{v, v, 255, 255}
}

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		ret[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			ret[i][j] = uint8(i ^ j + (i+j)/2)
		}
	}
	return ret
}

func main() {
	pic.ShowImage(SampleImage(200, 200))
}
