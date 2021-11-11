package evening

import "image/color"

func colorBrightOrange(n int) color.Color {
	r := n & 7
	g := (n >> 2) & 7
	b := (n >> 5) & 7
	return color.RGBA{
		R: 220 + uint8(r),
		G: 118 + 2*uint8(g),
		B: 38 + 2*uint8(b),
		A: 255,
	}
}

func colorOrange(n int) color.Color {
	r := n & 7
	g := (n >> 2) & 7
	b := (n >> 5) & 7
	return color.RGBA{
		R: 178 + 2*uint8(r),
		G: 86 + 4*uint8(g),
		B: 38 + 2*uint8(b),
		A: 255,
	}
}

func colorWhite(n int) color.Color {
	r := n & 7
	g := (n >> 2) & 7
	b := (n >> 5) & 7
	return color.RGBA{
		R: 255 - uint8(r),
		G: 255 - uint8(g),
		B: 255 - uint8(b),
		A: 255,
	}
}

func colorYellow(n int) color.Color {
	r := n & 7
	g := (n >> 2) & 7
	b := (n >> 5) & 7
	return color.RGBA{
		R: 250 - uint8(r),
		G: 242 - 2*uint8(g),
		B: 232 - 2*uint8(b),
		A: 255,
	}
}

func colorBrightYellow(n int) color.Color {
	r := n & 7
	g := (n >> 2) & 7
	b := (n >> 4) & 7
	return color.RGBA{
		R: 229 + uint8(r),
		G: 140 + 2*uint8(g),
		B: 128 + 2*uint8(b),
		A: 255,
	}
}

func colorYellowgreen(n int) color.Color {
	r := n & 7
	g := (n >> 2) & 7
	b := (n >> 5) & 7
	return color.RGBA{
		R: 140 + 2*uint8(r),
		G: 86 + uint8(g),
		B: 45 + uint8(b),
		A: 255,
	}
}

func colorGreen(n int) color.Color {
	r := n & 7
	g := (n >> 2) & 7
	b := (n >> 4) & 7
	return color.RGBA{
		R: 25 - 2*uint8(r),
		G: 76 + 2*uint8(g),
		B: 40 + uint8(b),
		A: 255,
	}
}

func colorBlue(n int) color.Color {
	r := n & 4
	g := (n >> 2) & 4
	b := (n >> 4) & 4
	return color.RGBA{
		R: 16 + 2*uint8(r),
		G: 16 + 3*uint8(g),
		B: 46 + 3*uint8(b),
		A: 255,
	}
}

func colorGround(n int) color.Color {
	r := n & 4
	g := (n >> 2) & 4
	b := (n >> 4) & 4
	return color.RGBA{
		R: 32 + 2*uint8(r),
		G: 24 + 2*uint8(g),
		B: 16 + 2*uint8(b),
		A: 255,
	}
}
