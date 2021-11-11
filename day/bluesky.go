package day

import (
	"image"
	"image/color"

	"github.com/rhymof/cryptseaside"
	"github.com/rhymof/cryptseaside/internal"
)

func Bluesky(rgba *image.RGBA, nano int64) *image.RGBA {
	seeds := internal.Sep(nano)
	for y := 0; y < 256; y++ {
		seedy := seeds[(11*y)%56]
		for x := 0; x < 256; x++ {
			seedx := seeds[(x+3*y)%56]
			switch {
			case y < cryptseaside.HorizonY:
				height := float64(y) / cryptseaside.HorizonY
				rgba.Set(x, y, internal.Noise(colorBluesky(height), seedx))
			case y < cryptseaside.GroundY:
				rgba.Set(x, y, internal.Noise(colorBluesea(seedy), seedx))
			default:
				rgba.Set(x, y, internal.Noise(colorGlound(seedy), seedx))
			}
		}
	}
	return rgba
}

func colorBluesea(n int) color.Color {
	r := n & 15
	g := (n >> 2) & 15
	b := (n >> 4) & 15
	return color.RGBA{
		R: uint8(26 + 2*r),
		G: uint8(126 + 2*g),
		B: uint8(168 + 3*b),
		A: 255,
	}
}

func colorBluesky(height float64) color.Color {
	return color.RGBA{
		R: uint8(46 + (140-46)*height/2),
		G: uint8(86 + (185-86)*height/2),
		B: uint8(156 + (240-156)*height/2),
		A: 255,
	}
}

func colorGlound(n int) color.Color {
	r := uint8(n & 7)
	g := uint8((n >> 3) & 7)
	b := uint8((n >> 5) & 7)
	basecolor := r + g + b + 63
	return color.RGBA{
		R: basecolor + r,
		G: basecolor + g,
		B: basecolor + b,
		A: 255,
	}
}
