package evening

import (
	"image"
	"math"

	"github.com/rhymof/cryptseaside"
	"github.com/rhymof/cryptseaside/internal"
)

func Sunset(rgba *image.RGBA, nano int64) *image.RGBA {
	seeds := internal.Sep(nano)
	seed := seeds[0] / 8
	for y := 0; y < 256; y++ {
		seedy := seeds[(11*y)%cryptseaside.CountSeeds]
		for x := 0; x < 256; x++ {
			seedx := seeds[(x+(3*y))%cryptseaside.CountSeeds]
			switch {
			case y < 60+seed:
				if (seedx+x)%(90+y) == 0 {
					rgba.Set(x, y, internal.Star(seedx))
				} else {
					rgba.Set(x, y, internal.Noise(colorBlue(seedy), seedx))
				}
			case y < 72+seed:
				if (seedx+x)%(127+y) == 0 {
					rgba.Set(x, y, internal.Star(seedx))
				} else {
					rgba.Set(x, y, internal.Noise(colorGreen(seedy), seedx))
				}
				if y < 72+seed && y >= 68+seed {
					rgba.Set(x, y, internal.Brend(rgba.At(x, y), colorYellowgreen(seedy), 0.5+0.1*float64(66+seed-y)))
				} else {
					rgba.Set(x, y, internal.Brend(rgba.At(x, y), colorBlue(seedy), 0.5+0.1*float64(y-60-seed)))
				}
			case y < 96+seed:
				rgba.Set(x, y, internal.Noise(colorYellowgreen(seedy), seedx))
			case y < 116+seed:
				rgba.Set(x, y, internal.Noise(colorOrange(seedy), seedx))
			case y < cryptseaside.HorizonY:
				rgba.Set(x, y, internal.Noise(colorBrightOrange(seedy), seedx))
			case y < cryptseaside.GroundY:
				if (seedx+x)%23 == 0 {
					rgba.Set(x, y, colorYellow(seedy))
				} else {
					rgba.Set(x, y, internal.Noise(colorBrightYellow(seedy), seedx))
				}
			default:
				rgba.Set(x, y, internal.Noise(colorGround(seedy), seedx))
			}
			if y < cryptseaside.GroundY {
				rgba.Set(x, y, internal.Brend(rgba.At(x, y), colorOrange(seedy), 0.5+0.5*float64(cryptseaside.GroundY-y)/cryptseaside.GroundY))
			}
		}
	}
	seed = seeds[seeds[0]%cryptseaside.CountSeeds]
	sx := (seed % 128) + 64
	sy := cryptseaside.HorizonY
	rgba = sun(rgba, sx, sy, seeds)
	return rgba
}

func sun(rgba *image.RGBA, sx, sy int, seeds []int) *image.RGBA {
	for y := -20; y < 26; y++ {
		seedy := seeds[(y+32)%cryptseaside.CountSeeds]
		if y <= 5 && y >= -5 {
			for x := -30 - int(math.Pow(5-math.Abs(float64(y)), 3)); x <= 30+int(math.Pow(5-math.Abs(float64(y)), 3)); x++ {
				if sx+x < 0 || sx+x >= 256 {
					continue
				}
				seedx := seeds[(2*x+y+1000)%cryptseaside.CountSeeds]
				rgba = setsun(rgba, sx, sy, x, y, seedx, seedy)
			}
		} else {
			for x := -28/(int(math.Abs(float64(y)))-4) + 1; x <= 28/(int(math.Abs(float64(y)))-4); x++ {
				if sx+x < 0 || sx+x >= 256 {
					continue
				}
				seedx := seeds[(2*x+y+1000)%cryptseaside.CountSeeds]
				rgba = setsun(rgba, sx, sy, x, y, seedx, seedy)
			}
		}
	}
	return rgba
}

func setsun(rgba *image.RGBA, sx, sy, x, y, seedx, seedy int) *image.RGBA {
	if math.Pow(float64(x), 2)+math.Pow(float64(y), 2) < 25 {
		rgba.Set(sx+x, sy+y, internal.Noise(colorWhite(seedx), seedy))
	} else {
		rgba.Set(sx+x, sy+y, internal.Noise(colorYellow(seedy), seedx))
	}
	return rgba
}
