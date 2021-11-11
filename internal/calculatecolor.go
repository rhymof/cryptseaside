package internal

import (
	"image/color"
)

func Brend(c0, c1 color.Color, weight float64) color.Color {
	r0, g0, b0, _ := c0.RGBA()
	r1, g1, b1, _ := c1.RGBA()
	r0, g0, b0 = r0/256, g0/256, b0/256
	r1, g1, b1 = r1/256, g1/256, b1/256
	return color.RGBA{
		R: uint8(float64(r0)*weight + float64(r1)*(1-weight)),
		G: uint8(float64(g0)*weight + float64(g1)*(1-weight)),
		B: uint8(float64(b0)*weight + float64(b1)*(1-weight)),
		A: 255,
	}
}
