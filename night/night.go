package night

import (
	"image"
	"image/color"
	"math"

	"github.com/rhymof/cryptseaside"
	"github.com/rhymof/cryptseaside/internal"
)

func Night(rgba *image.RGBA, nano int64) *image.RGBA {
	nebulaX, nebulaY := int(nano%255), int(nano%cryptseaside.HorizonY)/2+cryptseaside.HorizonY/10
	seeds := internal.Sep(nano)
	for y := 0; y < 256; y++ {
		seedy := seeds[y%cryptseaside.CountSeeds]
		for x := 0; x < 256; x++ {
			seedx := seeds[(x+y)%cryptseaside.CountSeeds]
			switch {
			case y < cryptseaside.HorizonY:
				rgba.Set(x, y, internal.Noise(colorNightSky((cryptseaside.GroundY-float64(y))/cryptseaside.GroundY, seeds), seedx))
				if (32*x+seedx)%(253+2*y) == 1 {
					rgba.Set(x, y, colorStar(seedx))
				}
				if math.Pow(float64(nebulaX-x), 2)+math.Pow(float64(nebulaY-y), 2) <= math.Pow(float64(10+(seedx&11+1)*(seedy&11+1)), 2) {
					rgba.Set(x, y, internal.Brend(rgba.At(x, y), colorNebula(seeds[0]), 0.87))
					if (seedx+seedy)%23 == 3 {
						rgba.Set(x, y, colorStar(seedx))
					}
				}
			case y < cryptseaside.GroundY:
				rgba.Set(x, y, internal.Noise(colorNightSea(x, seedy), seedx))
				if math.Pow(float64(nebulaX-x), 2) <= float64((seedy*seedy+seedx*seedx)%255) {
					rgba.Set(x, y, internal.Brend(rgba.At(x, y), colorNebula(seeds[0]), 0.92))
				}
			default:
				rgba.Set(x, y, internal.Noise(ground(seedy), seedx))
			}

		}
	}
	return rgba
}

func colorNightSky(height float64, seeds []int) color.Color {
	cHigh := color.RGBA{
		R: uint8(4 + seeds[0]&15),
		G: uint8(4 + seeds[1]&15),
		B: uint8(16 + seeds[2]&15),
		A: 255,
	}
	cLow := color.RGBA{
		R: uint8(32 + seeds[0]&15),
		G: uint8(32 + seeds[1]&15),
		B: uint8(48 + seeds[2]&15),
		A: 255,
	}
	return internal.Brend(cHigh, cLow, height)
}

func colorNightSea(x, seed int) color.Color {
	r := seed & 15
	g := (seed >> 2) & 15
	b := (seed >> 4) & 15
	if x >= 128 {
		x = 255 - x
	}
	return color.RGBA{
		R: uint8(r + x/32),
		G: uint8(16 + g + x/16),
		B: uint8(32 + x/8 + b),
	}
}

func ground(seed int) color.Color {
	return color.RGBA{
		R: uint8(seed & 15),
		G: uint8((seed >> 2) & 15),
		B: uint8((seed >> 4) & 15),
		A: 255,
	}
}

func colorStar(seed int) color.Color {
	r := uint8(seed) & 63
	g := uint8(seed>>1) & 63
	b := uint8(seed>>2) & 63
	return color.RGBA{
		R: 255 - r,
		G: 255 - g,
		B: 255 - b,
		A: 255,
	}
}

func colorNebula(seed int) color.Color {
	switch seed & 3 {
	case 0:
		return color.RGBA{0, 192, 255, 255}
	case 1:
		return color.RGBA{0, 255, 255, 255}
	case 2:
		return color.RGBA{255, 0, 255, 255}
	default:
		return color.RGBA{255, 255, 0, 255}
	}
}
