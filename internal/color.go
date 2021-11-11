package internal

import (
	"image/color"
	"math"
)

func Noise(c color.Color, n int) color.Color {
	r := float64(n & 7)
	g := float64((n >> 3) & 7)
	b := float64((n >> 5) & 7)
	cr, cg, cb, _ := c.RGBA()
	crf := float64(cr) / 256
	cgf := float64(cg) / 256
	cbf := float64(cb) / 256
	return color.RGBA{
		R: uint8(math.Abs(crf - r)),
		G: uint8(math.Abs(cgf - g)),
		B: uint8(math.Abs(cbf - b)),
		A: 255,
	}
}

func Star(n int) color.Color {
	r := uint8(n & 4)
	g := uint8((n >> 2) & 4)
	b := uint8((n >> 4) & 4)
	return color.RGBA{
		R: 255 - r*r,
		G: 255 - g*g,
		B: 255 - b*b,
		A: 255,
	}
}
