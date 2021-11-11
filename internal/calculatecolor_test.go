package internal

import (
	"image/color"
	"reflect"
	"testing"
)

func TestBrend(t *testing.T) {
	c0 := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	c1 := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	got := Brend(c0, c1, 127.0/255)
	want := color.RGBA{R: 128, G: 128, B: 128, A: 255}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v\nwant:%v", got, want)
	}
	c2 := color.RGBA{R: 100, G: 100, B: 100, A: 255}
	c3 := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	got = Brend(c2, c3, 0.5)
	want = color.RGBA{R: 150, G: 150, B: 150, A: 255}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v\nwant:%v", got, want)
	}
}
