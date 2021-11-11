package internal

import (
	"image/color"
	"reflect"
	"testing"
)

func TestNoise(t *testing.T) {
	c := color.RGBA{0, 0, 0, 255}
	n := 255
	want := color.RGBA{7, 7, 7, 255}
	got := Noise(c, n)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
