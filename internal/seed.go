package internal

import "github.com/rhymof/cryptseaside"

func Sep(seed int64) []int {
	seeds := make([]int, 0, 64)
	for i := 0; i < cryptseaside.CountSeeds; i++ {
		seeds = append(seeds, int((seed>>i)&255))
	}
	return seeds
}
